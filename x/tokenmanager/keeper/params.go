package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) GetNetwork(ctx sdk.Context, name string) (param types.NetworkParams, ok bool) {
	params := k.GetParams(ctx)
	for _, network := range params.Networks {
		if network.Name == name {
			return *network, true
		}
	}

	return
}
