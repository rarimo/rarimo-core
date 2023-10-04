package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func (k Keeper) SetItem(ctx sdk.Context, item types.Item) {
	prefix.
		NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemKeyPrefix)).
		Set(types.ItemKey(item.Index), k.cdc.MustMarshal(&item))
}

func (k Keeper) GetItem(ctx sdk.Context, index string) (val types.Item, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemKeyPrefix))

	b := store.Get(types.ItemKey(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllItem returns all item
func (k Keeper) GetAllItem(ctx sdk.Context) (list []types.Item) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Item
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// RemoveItem removes an item from the store
func (k Keeper) RemoveItem(ctx sdk.Context, index string) {
	prefix.
		NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemKeyPrefix)).
		Delete(types.ItemKey(index))
}

func (k Keeper) SetOnChainItem(ctx sdk.Context, item types.OnChainItem) {
	prefix.
		NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OnChainItemKeyPrefix)).
		Set(types.OnChainItemKey(item.Index), k.cdc.MustMarshal(&item))
}

func (k Keeper) GetOnChainItem(ctx sdk.Context, index *types.OnChainItemIndex) (val types.OnChainItem, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OnChainItemKeyPrefix))

	b := store.Get(types.OnChainItemKey(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllOnChainItem returns all item
func (k Keeper) GetAllOnChainItem(ctx sdk.Context) (list []types.OnChainItem) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OnChainItemKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.OnChainItem
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// RemoveOnChainItem removes an item from the store
func (k Keeper) RemoveOnChainItem(ctx sdk.Context, index *types.OnChainItemIndex) {
	prefix.
		NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OnChainItemKeyPrefix)).
		Delete(types.OnChainItemKey(index))
}

func (k Keeper) SetSeed(ctx sdk.Context, seed types.Seed) {
	prefix.
		NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SeedKeyPrefix)).
		Set(types.SeedKey(seed.Seed), k.cdc.MustMarshal(&seed))
}

func (k Keeper) GetSeed(ctx sdk.Context, seed string) (val types.Seed, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SeedKeyPrefix))

	b := store.Get(types.SeedKey(seed))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetAllSeed(ctx sdk.Context) (list []types.Seed) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SeedKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Seed
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) RemoveSeed(ctx sdk.Context, seed string) {
	prefix.
		NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SeedKeyPrefix)).
		Delete(types.SeedKey(seed))
}
