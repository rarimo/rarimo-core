package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func (k Keeper) GetCollectionInfo(ctx sdk.Context, index string) *types.CollectionInfo {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollectionKeyPrefix))

	b := store.Get(types.CollectionKey(index))
	if b == nil {
		return nil
	}

	var val types.CollectionInfo
	k.cdc.MustUnmarshal(b, &val)
	return &val
}

func (k Keeper) GetCollectionInfoByNetworkParams(ctx sdk.Context, chain string, address string) *types.CollectionInfo {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollectionAddrPrefix))

	b := store.Get(types.CollectionNetworkKey(chain, address))
	if b == nil {
		return nil
	}

	var val types.CollectionInfo
	k.cdc.MustUnmarshal(b, &val)
	return &val
}

// GetAllItem returns all item
func (k Keeper) GetAllCollections(ctx sdk.Context) (list []types.CollectionInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollectionKeyPrefix))

	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CollectionInfo
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) PutCollectionInfo(ctx sdk.Context, info types.CollectionInfo) {
	infoProto := k.cdc.MustMarshal(&info)

	prefix.
		NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollectionKeyPrefix)).
		Set(types.CollectionKey(info.Index), infoProto)

	for chain, params := range info.ChainParams { // making collection searchable by every address
		prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollectionAddrPrefix)).
			Set(types.CollectionNetworkKey(chain, params.Address), infoProto)
	}
}

// RemoveCollectionInfoByAddress removes only the collection info associated with the particular chain:address pair
func (k Keeper) RemoveCollectionInfoByAddress(ctx sdk.Context, chain string, address string) {
	prefix.
		NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollectionKeyPrefix)).
		Delete(types.CollectionNetworkKey(chain, address))
}

// RemoveCollectionInfo completely removes the collection from the store
// (all the records associated with both index and all chain:address pairs).
// Note: cannot remove collection with items
func (k Keeper) RemoveCollectionInfo(ctx sdk.Context, index string) {
	collection := k.GetCollectionInfo(ctx, index)

	if len(collection.Items) != 0 {
		panic("cannot remove collection with items")
	}

	for chain, params := range collection.ChainParams { // removing collection from address index
		k.RemoveCollectionInfoByAddress(ctx, chain, params.Address)
	}

	prefix.
		NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollectionKeyPrefix)).
		Delete(types.CollectionKey(index))
}
