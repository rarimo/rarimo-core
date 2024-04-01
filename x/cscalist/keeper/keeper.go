package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/cscalist/types"
)

type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey storetypes.StoreKey
	memKey   storetypes.StoreKey
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
	}
}

func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	store := ctx.KVStore(k.storeKey)

	b := store.Get(types.KeyPrefix(types.ParamsKey))
	if b == nil {
		return types.DefaultParams()
	}

	k.cdc.MustUnmarshal(b, &params)
	return params
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	if err := params.Validate(); err != nil {
		panic("failed to set params: " + err.Error())
	}

	b := k.cdc.MustMarshal(&params)
	ctx.KVStore(k.storeKey).Set(types.KeyPrefix(types.ParamsKey), b)
}

func (k Keeper) GetRootKey(ctx sdk.Context) string {
	params := k.GetParams(ctx)
	return params.RootKey
}

func (k Keeper) SetRootKey(ctx sdk.Context, root string) {
	params := k.GetParams(ctx)
	params.RootKey = root
	k.SetParams(ctx, params)
}
