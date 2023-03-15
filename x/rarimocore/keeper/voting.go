package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/pkg"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
	tokentypes "gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func (k Keeper) CreateVote(ctx sdk.Context, vote types.Vote) (bool, error) {
	operation, ok := k.GetOperation(ctx, vote.Index.Operation)
	if !ok {
		return false, sdkerrors.Wrap(sdkerrors.ErrNotFound, "operation not found")
	}

	k.SetVote(ctx, vote)

	if operation.Status != types.OpStatus_INITIALIZED && operation.Status != types.OpStatus_NOT_APPROVED {
		return false, nil
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeVoted,
		sdk.NewAttribute(types.AttributeKeyOperationId, operation.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, operation.OperationType.String()),
		sdk.NewAttribute(types.AttributeKeyVotingChoice, vote.Vote.String()),
	))

	yesResult := sdk.ZeroDec()
	noResult := sdk.ZeroDec()

	// Setting votes in validators map
	k.IterateVotes(ctx, vote.Index.Operation, func(vote types.Vote) (stop bool) {
		voter := sdk.MustAccAddressFromBech32(vote.Index.Validator)

		if validator := k.staking.Validator(ctx, sdk.ValAddress(voter)); validator != nil {
			switch vote.Vote {
			case types.VoteType_YES:
				yesResult = yesResult.Add(validator.GetBondedTokens().ToDec())
			case types.VoteType_NO:
				noResult = noResult.Add(validator.GetBondedTokens().ToDec())
			}
		}

		return false
	})

	params := k.GetParams(ctx)
	quorum, _ := sdk.NewDecFromStr(params.VoteQuorum)
	threshold, _ := sdk.NewDecFromStr(params.VoteThreshold)

	totalVotingPower := yesResult.Add(noResult)

	// If there is not enough quorum of votes, finish the flow
	percentVoting := totalVotingPower.Quo(k.staking.TotalBondedTokens(ctx).ToDec())
	if percentVoting.LT(quorum) {
		return false, nil
	}

	var firstUpdate = false

	if operation.Status == types.OpStatus_INITIALIZED {
		// Firstly moving from initialized to approved/unapproved status
		firstUpdate = true
	}

	if yesResult.Quo(totalVotingPower).GT(threshold) {
		if err := k.ApproveOperation(ctx, operation); err != nil {
			return false, err
		}

		return firstUpdate, nil
	}

	if err := k.UnapproveOperation(ctx, operation); err != nil {
		return false, err
	}

	return firstUpdate, nil
}

func (k Keeper) UnapproveOperation(ctx sdk.Context, op types.Operation) error {
	op.Status = types.OpStatus_NOT_APPROVED
	k.SetOperation(ctx, op)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeOperationRejected,
		sdk.NewAttribute(types.AttributeKeyOperationId, op.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, op.OperationType.String()),
	))

	return nil
}

func (k Keeper) ApproveOperation(ctx sdk.Context, op types.Operation) error {
	switch op.OperationType {
	case types.OpType_TRANSFER:
		transfer, _ := pkg.GetTransfer(op)
		if err := k.ApproveTransferOperation(ctx, transfer); err != nil {
			return err
		}

		op.Status = types.OpStatus_APPROVED
		k.SetOperation(ctx, op)
	default:
		// Nothing to do
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeOperationApproved,
		sdk.NewAttribute(types.AttributeKeyOperationId, op.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, op.OperationType.String()),
	))

	return nil
}

func (k Keeper) ApproveTransferOperation(ctx sdk.Context, transfer *types.Transfer) error {
	data, ok := k.tm.GetCollectionData(ctx, &tokentypes.CollectionDataIndex{Chain: transfer.From.Chain, Address: transfer.From.Address})
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "collection data does not exists")
	}

	from, ok := k.tm.GetOnChainItem(ctx, transfer.From)
	if !ok {
		// Item also does not exist
		item := tokentypes.Item{
			Index:      transfer.Origin,
			Collection: data.Collection,
			Meta:       transfer.Meta,
			OnChain:    []*tokentypes.OnChainItemIndex{transfer.From},
		}

		// Indexing seed and check if already exists. Meta not-nil validated during op creation
		if item.Meta.Seed != "" {
			if _, ok := k.tm.GetSeed(ctx, item.Meta.Seed); ok {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "seed already exists")
			}

			k.tm.SetSeed(ctx, tokentypes.Seed{
				Seed: item.Meta.Seed,
				Item: item.Index,
			})
		}

		k.tm.SetItem(ctx, item)
		from = tokentypes.OnChainItem{
			Index: transfer.From,
			Item:  item.Index,
		}

		k.tm.SetOnChainItem(ctx, from)
	}

	to, ok := k.tm.GetOnChainItem(ctx, transfer.From)
	if !ok {
		item, ok := k.tm.GetItem(ctx, from.Item)
		if !ok {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "item does not exists but 'from' on chain item found")
		}

		item.OnChain = append(item.OnChain, transfer.To)
		k.tm.SetItem(ctx, item)

		to = tokentypes.OnChainItem{
			Index: transfer.To,
			Item:  item.Index,
		}

		k.tm.SetOnChainItem(ctx, to)
	}

	return nil
}
