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

	if _, ok := k.GetVote(ctx, vote.Index); ok {
		return false, sdkerrors.Wrap(sdkerrors.ErrNotFound, "vote already exists")
	}

	if operation.Status != types.OpStatus_INITIALIZED {
		return false, nil
	}

	k.SetVote(ctx, vote)

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
	default:
		// Nothing to do
	}

	k.ApproveOperationDefault(ctx, op)

	return nil
}

func (k Keeper) ApproveOperationDefault(ctx sdk.Context, op types.Operation) {
	op.Status = types.OpStatus_APPROVED
	k.SetOperation(ctx, op)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeOperationApproved,
		sdk.NewAttribute(types.AttributeKeyOperationId, op.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, op.OperationType.String()),
	))
}

func (k Keeper) ApproveTransferOperation(ctx sdk.Context, transfer *types.Transfer) error {
	data, ok := k.tm.GetCollectionData(ctx, &tokentypes.CollectionDataIndex{Chain: transfer.From.Chain, Address: transfer.From.Address})
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "collection data does not exists")
	}

	from, ok := k.tm.GetOnChainItem(ctx, &transfer.From)
	if !ok {
		// Item also does not exist
		item := tokentypes.Item{
			Index:      transfer.Origin,
			Collection: data.Collection,
			Meta:       *transfer.Meta,
			OnChain:    []*tokentypes.OnChainItemIndex{&transfer.From},
		}

		k.tm.SetItem(ctx, item)
		ctx.EventManager().EmitEvent(sdk.NewEvent(tokentypes.EventTypeItemCreated,
			sdk.NewAttribute(tokentypes.AttributeKeyItemIndex, item.Index),
		))

		from = tokentypes.OnChainItem{
			Index: &transfer.From,
			Item:  item.Index,
		}

		k.tm.SetOnChainItem(ctx, from)
		ctx.EventManager().EmitEvent(sdk.NewEvent(tokentypes.EventTypeOnChainItemCreated,
			sdk.NewAttribute(tokentypes.AttributeKeyItemIndex, from.Item),
			sdk.NewAttribute(tokentypes.AttributeKeyOnChainItemChain, from.Index.Chain),
		))

		// Indexing seed and check if already exists. Meta not-nil validated during op creation
		if item.Meta.Seed != "" {
			if _, ok := k.tm.GetSeed(ctx, item.Meta.Seed); ok {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "seed already exists")
			}

			k.tm.SetSeed(ctx, tokentypes.Seed{
				Seed: item.Meta.Seed,
				Item: item.Index,
			})
			ctx.EventManager().EmitEvent(sdk.NewEvent(tokentypes.EventTypeSeedCreated,
				sdk.NewAttribute(tokentypes.AttributeKeyItemIndex, item.Index),
				sdk.NewAttribute(tokentypes.AttributeKeySeed, item.Meta.Seed),
			))
		}
	}

	to, ok := k.tm.GetOnChainItem(ctx, &transfer.To)
	if !ok {
		item, ok := k.tm.GetItem(ctx, from.Item)
		if !ok {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "item does not exists but 'from' on chain item found")
		}

		item.OnChain = append(item.OnChain, &transfer.To)
		k.tm.SetItem(ctx, item)

		to = tokentypes.OnChainItem{
			Index: &transfer.To,
			Item:  item.Index,
		}

		k.tm.SetOnChainItem(ctx, to)
		ctx.EventManager().EmitEvent(sdk.NewEvent(tokentypes.EventTypeOnChainItemCreated,
			sdk.NewAttribute(tokentypes.AttributeKeyItemIndex, to.Item),
			sdk.NewAttribute(tokentypes.AttributeKeyOnChainItemChain, to.Index.Chain),
		))
	}

	return nil
}
