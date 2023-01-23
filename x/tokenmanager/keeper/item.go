package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func (k Keeper) PutItem(ctx sdk.Context, item types.Item) {
	if item.Index == nil {
		return
	}

	prefix.
		NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemKeyPrefix)).
		Set(types.ItemKey(item.Index), k.cdc.MustMarshal(&item))
}

func (k Keeper) GetItem(ctx sdk.Context, index *types.ItemIndex) (val types.Item, found bool) {
	if index == nil {
		return val, false
	}

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
func (k Keeper) RemoveItem(ctx sdk.Context, index *types.ItemIndex) {
	if index == nil {
		return
	}

	prefix.
		NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemKeyPrefix)).
		Delete(types.ItemKey(index))
}
