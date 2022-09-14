package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarify-protocol/rarimo-core/x/tokenmanager/saver"
	"gitlab.com/rarify-protocol/rarimo-core/x/tokenmanager/types"
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

func (k Keeper) GetNetwork(ctx sdk.Context, network string) (param *types.ChainParams, ok bool) {
	var params types.Params
	k.paramstore.GetParamSet(ctx, &params)
	param, ok = params.Networks[network]
	return param, ok
}
