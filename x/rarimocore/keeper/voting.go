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

	return true, nil
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
