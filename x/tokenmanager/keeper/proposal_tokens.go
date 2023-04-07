package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func (k Keeper) HandleUpdateTokenItemProposal(ctx sdk.Context, proposal *types.UpdateTokenItemProposal) error {
	for _, newItem := range proposal.Item {

		oldItem, ok := k.GetItem(ctx, newItem.Index)
		if !ok {
			return sdkerrors.Wrap(sdkerrors.ErrNotFound, "not found")
		}

		if oldItem.Meta.Seed != newItem.Meta.Seed {
			k.RemoveSeed(ctx, oldItem.Meta.Seed)

			ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeSeedRemoved,
				sdk.NewAttribute(types.AttributeKeyItemIndex, oldItem.Index),
				sdk.NewAttribute(types.AttributeKeySeed, oldItem.Meta.Seed),
			))

			if newItem.Meta.Seed != "" {
				k.SetSeed(ctx, types.Seed{
					Seed: newItem.Meta.Seed,
					Item: newItem.Index,
				})

				ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeSeedCreated,
					sdk.NewAttribute(types.AttributeKeyItemIndex, newItem.Index),
					sdk.NewAttribute(types.AttributeKeySeed, newItem.Meta.Seed),
				))
			}
		}

		k.SetItem(ctx, *newItem)
	}

	return nil
}

func (k Keeper) HandleRemoveTokenItemProposal(ctx sdk.Context, proposal *types.RemoveTokenItemProposal) error {
	for _, index := range proposal.Index {
		item, ok := k.GetItem(ctx, index)
		if !ok {
			return sdkerrors.Wrap(sdkerrors.ErrNotFound, "not found")
		}

		for _, onChainIndex := range item.OnChain {
			k.RemoveOnChainItem(ctx, onChainIndex)

			ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeOnChainItemRemoved,
				sdk.NewAttribute(types.AttributeKeyItemIndex, index),
				sdk.NewAttribute(types.AttributeKeyOnChainItemChain, onChainIndex.Chain),
			))
		}

		k.RemoveSeed(ctx, item.Meta.Seed)
		ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeSeedRemoved,
			sdk.NewAttribute(types.AttributeKeyItemIndex, item.Index),
			sdk.NewAttribute(types.AttributeKeySeed, item.Meta.Seed),
		))

		k.RemoveItem(ctx, index)
		ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeItemRemoved,
			sdk.NewAttribute(types.AttributeKeyItemIndex, index),
		))
	}

	return nil
}
