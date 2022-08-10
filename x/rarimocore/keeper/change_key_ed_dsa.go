package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

// SetChangeKeyEdDSA set a specific changeKeyEdDSA in the store from its index
func (k Keeper) SetChangeKeyEdDSA(ctx sdk.Context, changeKeyEdDSA types.ChangeKeyEdDSA) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChangeKeyEdDSAKeyPrefix))
	b := k.cdc.MustMarshal(&changeKeyEdDSA)
	store.Set(types.ChangeKeyEdDSAKey(
		changeKeyEdDSA.NewKey,
	), b)
}

// GetChangeKeyEdDSA returns a changeKeyEdDSA from its index
func (k Keeper) GetChangeKeyEdDSA(
	ctx sdk.Context,
	newKey string,

) (val types.ChangeKeyEdDSA, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChangeKeyEdDSAKeyPrefix))

	b := store.Get(types.ChangeKeyEdDSAKey(
		newKey,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveChangeKeyEdDSA removes a changeKeyEdDSA from the store
func (k Keeper) RemoveChangeKeyEdDSA(
	ctx sdk.Context,
	newKey string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChangeKeyEdDSAKeyPrefix))
	store.Delete(types.ChangeKeyEdDSAKey(
		newKey,
	))
}

// GetAllChangeKeyEdDSA returns all changeKeyEdDSA
func (k Keeper) GetAllChangeKeyEdDSA(ctx sdk.Context) (list []types.ChangeKeyEdDSA) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChangeKeyEdDSAKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ChangeKeyEdDSA
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
