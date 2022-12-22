package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/saver"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	saver.Set(params.Networks)
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) GetNetwork(ctx sdk.Context, name string) (param *types.NetworkParams, ok bool) {
	var params types.Params
	k.paramstore.GetParamSet(ctx, &params)

	for _, network := range params.Networks {
		if network.Name == name {
			param, ok = network, true
			break
		}
	}

	return param, ok
}

func (k Keeper) GetSaverType(ctx sdk.Context, name string, coreType types.Type) uint32 {
	var params types.Params
	k.paramstore.GetParamSet(ctx, &params)

	for _, network := range params.Networks {
		if network.Name == name {
			for _, binding := range network.Types {
				if binding.CoreType == uint32(coreType) {
					return binding.SaverType
				}
			}
		}
	}

	return 0
}
