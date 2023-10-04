package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func (k Keeper) GetCollection(ctx sdk.Context, index string) (val types.Collection, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollectionKeyPrefix))

	b := store.Get(types.CollectionKey(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetAllCollections(ctx sdk.Context) (list []types.Collection) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollectionKeyPrefix))

	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Collection
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) SetCollection(ctx sdk.Context, collection types.Collection) {
	prefix.
		NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollectionKeyPrefix)).
		Set(types.CollectionKey(collection.Index), k.cdc.MustMarshal(&collection))
}

func (k Keeper) RemoveCollection(ctx sdk.Context, index string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollectionDataKeyPrefix))

	iterator := sdk.KVStorePrefixIterator(store, []byte(index))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		store.Delete(iterator.Key())
	}

	store = prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemKeyPrefix))

	iterator = sdk.KVStorePrefixIterator(store, []byte(index))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		store.Delete(iterator.Key())
	}

	prefix.
		NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollectionKeyPrefix)).
		Delete(types.CollectionKey(index))
}

func (k Keeper) GetCollectionData(ctx sdk.Context, index *types.CollectionDataIndex) (val types.CollectionData, found bool) {
	if index == nil {
		return val, false
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollectionDataKeyPrefix))

	b := store.Get(types.CollectionDataKey(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetAllCollectionData(ctx sdk.Context) (list []types.CollectionData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollectionDataKeyPrefix))

	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CollectionData
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) SetCollectionData(ctx sdk.Context, data types.CollectionData) {
	prefix.
		NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollectionDataKeyPrefix)).
		Set(types.CollectionDataKey(data.Index), k.cdc.MustMarshal(&data))
}

func (k Keeper) RemoveCollectionData(ctx sdk.Context, index *types.CollectionDataIndex) {
	if index == nil {
		return
	}

	prefix.
		NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollectionDataKeyPrefix)).
		Delete(types.CollectionDataKey(index))
}
