package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/multisig/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	store := ctx.KVStore(k.storeKey)

	b := store.Get(types.KeyPrefix(types.ParamsKey))
	if b == nil {
		return types.DefaultParams()
	}

	k.cdc.MustUnmarshal(b, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&params)
	store.Set(types.KeyPrefix(types.ParamsKey), b)
}
