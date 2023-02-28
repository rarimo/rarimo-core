package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func (k Keeper) HandleSetNetworkProposal(ctx sdk.Context, proposal *types.SetNetworkProposal) error {
	params := k.GetParams(ctx)
	for i, network := range params.Networks {
		if network.Name == proposal.NetworkParams.Name {
			params.Networks[i] = proposal.NetworkParams
			k.SetParams(ctx, params)
			return nil
		}
	}

	params.Networks = append(params.Networks, proposal.NetworkParams)
	k.SetParams(ctx, params)
	return nil
}

func (k Keeper) HandleUpdateTokenItemProposal(ctx sdk.Context, proposal *types.UpdateTokenItemProposal) error {
	for _, newItem := range proposal.Item {

		oldItem, ok := k.GetItem(ctx, newItem.Index)
		if !ok {
			return sdkerrors.Wrap(sdkerrors.ErrNotFound, "not found")
		}

		if oldItem.Meta.Seed != newItem.Meta.Seed {
			k.RemoveSeed(ctx, oldItem.Meta.Seed)
			if newItem.Meta.Seed != "" {
				k.SetSeed(ctx, types.Seed{
					Seed: newItem.Meta.Seed,
					Item: newItem.Index,
				})
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
		}

		k.RemoveSeed(ctx, item.Meta.Seed)
		k.RemoveItem(ctx, index)
	}

	return nil
}
