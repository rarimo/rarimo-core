package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/oraclemanager/types"
)

// SetOracle set a specific oracle in the store from its index
func (k Keeper) SetOracle(ctx sdk.Context, op types.Oracle) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OracleKeyPrefix))
	b := k.cdc.MustMarshal(&op)
	store.Set(types.OracleKey(
		op.Index,
	), b)
}

// GetOracle returns an oracle from its index
func (k Keeper) GetOracle(
	ctx sdk.Context,
	index *types.OracleIndex,
) (val types.Oracle, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OracleKeyPrefix))

	b := store.Get(types.OracleKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOracle removes an oracle from the store
func (k Keeper) RemoveOracle(
	ctx sdk.Context,
	index *types.OracleIndex,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OracleKeyPrefix))
	store.Delete(types.OracleKey(
		index,
	))
}

// GetAllOracle returns all oracles
func (k Keeper) GetAllOracle(ctx sdk.Context) (list []types.Oracle) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OracleKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Oracle
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) GetOracleForChain(ctx sdk.Context, chain string) (list []types.Oracle) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OracleKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte(chain))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Oracle
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
