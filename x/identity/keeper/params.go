package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/identity/types"
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

func (k Keeper) SetGIST(ctx sdk.Context, gist string) {
	params := k.GetParams(ctx)
	params.GISTHash = gist
	params.GISTUpdatedTimestamp = uint64(ctx.BlockTime().UTC().Unix())
	k.SetParams(ctx, params)
}

func (k Keeper) AddToWaitingList(ctx sdk.Context, id string) {
	params := k.GetParams(ctx)
	params.StatesWaitingForSign = append(params.StatesWaitingForSign, id)
	k.SetParams(ctx, params)
}
