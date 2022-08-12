package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
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

func (k Keeper) UpdateKeyECDSA(ctx sdk.Context, newKey string) {
	k.paramstore.Set(ctx, types.KeyECDSA, newKey)
}

func (k Keeper) UpdateKeyEdDSA(ctx sdk.Context, newKey string) {
	k.paramstore.Set(ctx, types.KeyEdDSA, newKey)
}

func (k Keeper) GetKeyECDSA(ctx sdk.Context) string {
	params := k.GetParams(ctx)
	return params.GetKeyECDSA()
}

func (k Keeper) GetKeyEdDSA(ctx sdk.Context) string {
	params := k.GetParams(ctx)
	return params.GetKeyEdDSA()
}
