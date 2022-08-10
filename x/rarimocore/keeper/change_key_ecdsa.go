package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

// SetChangeKeyECDSA set a specific changeKeyECDSA in the store from its index
func (k Keeper) SetChangeKeyECDSA(ctx sdk.Context, changeKeyECDSA types.ChangeKeyECDSA) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChangeKeyECDSAKeyPrefix))
	b := k.cdc.MustMarshal(&changeKeyECDSA)
	store.Set(types.ChangeKeyECDSAKey(
		changeKeyECDSA.NewKey,
	), b)
}

// GetChangeKeyECDSA returns a changeKeyECDSA from its index
func (k Keeper) GetChangeKeyECDSA(
	ctx sdk.Context,
	newKey string,

) (val types.ChangeKeyECDSA, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChangeKeyECDSAKeyPrefix))

	b := store.Get(types.ChangeKeyECDSAKey(
		newKey,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveChangeKeyECDSA removes a changeKeyECDSA from the store
func (k Keeper) RemoveChangeKeyECDSA(
	ctx sdk.Context,
	newKey string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChangeKeyECDSAKeyPrefix))
	store.Delete(types.ChangeKeyECDSAKey(
		newKey,
	))
}

// GetAllChangeKeyECDSA returns all changeKeyECDSA
func (k Keeper) GetAllChangeKeyECDSA(ctx sdk.Context) (list []types.ChangeKeyECDSA) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChangeKeyECDSAKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ChangeKeyECDSA
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
