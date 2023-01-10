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
	prefix.
		NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollectionKeyPrefix)).
		Set(types.CollectionKey(info.Index), k.cdc.MustMarshal(&info))
}

func (k Keeper) RemoveCollectionInfo(ctx sdk.Context, index string) {
	prefix.
		NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollectionKeyPrefix)).
		Delete(types.CollectionKey(index))
}
