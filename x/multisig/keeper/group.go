package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/multisig/types"
)

// SetGroup set a specific group in the store from its groupAccount
func (k Keeper) SetGroup(ctx sdk.Context, group types.Group) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupKeyPrefix))
	b := k.cdc.MustMarshal(&group)
	store.Set(types.GroupKey(group.Account), b)
}

// GetGroup returns a group from its groupAccount
func (k Keeper) GetGroup(ctx sdk.Context, groupAccount string) (val types.Group, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupKeyPrefix))

	b := store.Get(types.GroupKey(groupAccount))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveGroup removes a group from the store
func (k Keeper) RemoveGroup(ctx sdk.Context, groupAccount string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupKeyPrefix))
	store.Delete(types.GroupKey(groupAccount))
}

// GetAllGroup returns all groups
func (k Keeper) GetAllGroup(ctx sdk.Context) (list []types.Group) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Group
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
