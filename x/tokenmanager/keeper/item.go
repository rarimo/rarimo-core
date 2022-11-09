package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarify-protocol/rarimo-core/x/tokenmanager/types"
)

// SetItem set a specific item in the store from its index
func (k Keeper) SetItem(ctx sdk.Context, item types.Item) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemKeyPrefix))
	b := k.cdc.MustMarshal(&item)
	store.Set(types.ItemKey(
		item.TokenAddress,
		item.TokenId,
		item.Chain,
	), b)
}

// GetItem returns an item from its index
func (k Keeper) GetItem(
	ctx sdk.Context,
	tokenAddress string,
	tokenId string,
	chain string,
) (val types.Item, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemKeyPrefix))

	b := store.Get(types.ItemKey(
		tokenAddress,
		tokenId,
		chain,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetItem returns an item from its index
func (k Keeper) GetItemByChain(
	ctx sdk.Context,
	infoIndex string,
	chain string,
) (val types.Item, found bool) {
	info, ok := k.GetInfo(ctx, infoIndex)
	if !ok {
		return val, false
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemKeyPrefix))
	b := store.Get(types.ItemKey(
		info.Chains[chain].TokenAddress,
		info.Chains[chain].TokenId,
		chain,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveItem removes an item from the store
func (k Keeper) RemoveItem(
	ctx sdk.Context,
	tokenAddress string,
	tokenId string,
	chain string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemKeyPrefix))
	store.Delete(types.ItemKey(
		tokenAddress,
		tokenId,
		chain,
	))
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
