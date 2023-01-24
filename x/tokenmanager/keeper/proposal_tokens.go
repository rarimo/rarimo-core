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

		if _, ok := k.GetItem(ctx, newItem.Index); !ok {
			return sdkerrors.Wrap(sdkerrors.ErrNotFound, "not found")
		}

		k.SetItem(ctx, *newItem)
	}

	return nil
}

func (k Keeper) HandleRemoveTokenItemProposal(ctx sdk.Context, proposal *types.RemoveTokenItemProposal) error {
	for _, index := range proposal.Index {
		k.RemoveItem(ctx, index)
	}

	return nil
}
