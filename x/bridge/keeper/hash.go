package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/bridge/types"
)

// SetHash set a specific hash in the store from its index
func (k Keeper) SetHash(ctx sdk.Context, hash types.Hash) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HashKeyPrefix))
	b := k.cdc.MustMarshal(&hash)
	store.Set(types.HashKey(
		hash.Index,
	), b)
}

// GetHash returns a hash from its index
func (k Keeper) GetHash(
	ctx sdk.Context,
	index string,

) (val types.Hash, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HashKeyPrefix))

	b := store.Get(types.HashKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveHash removes a hash from the store
func (k Keeper) RemoveHash(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HashKeyPrefix))
	store.Delete(types.HashKey(
		index,
	))
}

// GetAllHash returns all hash
func (k Keeper) GetAllHash(ctx sdk.Context) (list []types.Hash) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HashKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Hash
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
