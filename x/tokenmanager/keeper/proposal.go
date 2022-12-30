package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func (k Keeper) SetTokenInfoProposal(ctx sdk.Context, proposal *types.SetTokenInfoProposal) error {
	k.SetInfo(ctx, *proposal.Info)
	return nil
}

func (k Keeper) SetTokenItemProposal(ctx sdk.Context, proposal *types.SetTokenItemProposal) error {
	if _, ok := k.GetInfo(ctx, proposal.Item.Index); ok {
		k.SetItem(ctx, *proposal.Item)
		return nil
	}
	return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("info with index %s not found", proposal.Item.Index))
}

func (k Keeper) RemoveTokenItemProposal(ctx sdk.Context, proposal *types.RemoveTokenItemProposal) error {
	k.RemoveItem(ctx, proposal.TokenAddress, proposal.TokenId, proposal.Network)
	return nil
}

func (k Keeper) RemoveTokenInfoProposal(ctx sdk.Context, proposal *types.RemoveTokenInfoProposal) error {
	k.RemoveInfo(ctx, proposal.Index)
	return nil
}

func (k Keeper) SetNetworkProposal(ctx sdk.Context, proposal *types.SetNetworkProposal) error {
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
