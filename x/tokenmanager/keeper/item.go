package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func (k Keeper) PutItem(ctx sdk.Context, item types.Item, keys ...string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemKeyPrefix))
	b := k.cdc.MustMarshal(&item)

	for _, key := range keys {
		store.Set([]byte(key), b) // making item searchable by every passed key
	}
}

func (k Keeper) GetItemByIndex(ctx sdk.Context, index string) (val types.Item, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemKeyPrefix))

	b := store.Get([]byte(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetItem returns an item from its index
func (k Keeper) GetItem(
	ctx sdk.Context,
	tokenAddress string,
	tokenId string,
	chain string,
) (val types.Item, found bool) {
	return k.GetItemByIndex(ctx, types.ItemIndex(
		tokenAddress,
		tokenId,
		chain,
	))
}

// RemoveItem removes an item from the store
func (k Keeper) RemoveItem(
	ctx sdk.Context,
	tokenAddress string,
	tokenId string,
	chain string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemKeyPrefix))
	store.Delete([]byte(types.ItemIndex(
		tokenAddress,
		tokenId,
		chain,
	)))
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
