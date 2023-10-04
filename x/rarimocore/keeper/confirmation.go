package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/rarimocore/types"
)

// SetConfirmation set a specific confirmation in the store from its index
func (k Keeper) SetConfirmation(ctx sdk.Context, confirmation types.Confirmation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ConfirmationKeyPrefix))
	b := k.cdc.MustMarshal(&confirmation)
	store.Set(types.ConfirmationKey(
		confirmation.Root,
	), b)
}

// GetConfirmation returns a confirmation from its index
func (k Keeper) GetConfirmation(
	ctx sdk.Context,
	root string,

) (val types.Confirmation, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ConfirmationKeyPrefix))

	b := store.Get(types.ConfirmationKey(
		root,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveConfirmation removes a confirmation from the store
func (k Keeper) RemoveConfirmation(
	ctx sdk.Context,
	height string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ConfirmationKeyPrefix))
	store.Delete(types.ConfirmationKey(
		height,
	))
}

// GetAllConfirmation returns all confirmation
func (k Keeper) GetAllConfirmation(ctx sdk.Context) (list []types.Confirmation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ConfirmationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Confirmation
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
