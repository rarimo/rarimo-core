package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/identity/types"
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
	if err := params.Validate(); err != nil {
		panic("failed to set params: " + err.Error())
	}

	b := k.cdc.MustMarshal(&params)
	ctx.KVStore(k.storeKey).Set(types.KeyPrefix(types.ParamsKey), b)
}

func (k Keeper) GetRootKey(ctx sdk.Context) string {
	params := k.GetParams(ctx)
	return params.TreapRootKey
}

func (k Keeper) SetRootKey(ctx sdk.Context, root string) {
	params := k.GetParams(ctx)
	params.TreapRootKey = root
	k.SetParams(ctx, params)
}
