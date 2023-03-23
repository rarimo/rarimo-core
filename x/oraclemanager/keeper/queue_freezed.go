package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/oraclemanager/types"
)

func (k Keeper) AddToFreezedQueue(ctx sdk.Context, block uint64, index *types.OracleIndex) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.FreezedOracleQueueKey(index, block), []byte{0x1})
}

func (k Keeper) RemoveFromFreezedQueue(ctx sdk.Context, block uint64, index *types.OracleIndex) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.FreezedOracleQueueKey(index, block))
}

func (k Keeper) IterateOverFreezedQueue(ctx sdk.Context, block uint64, cb func(oracle types.Oracle) (stop bool)) {
	iterator := k.FreezedQueueIterator(ctx, block)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		index, _ := types.SplitFreezedOracleQueueKey(iterator.Key())
		oracle, ok := k.GetOracle(ctx, index)
		if !ok {
			panic(fmt.Sprintf("oracle %s does not exist", index.String()))
		}

		if cb(oracle) {
			break
		}
	}
}

func (k Keeper) FreezedQueueIterator(ctx sdk.Context, endBlock uint64) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.FreezedOracleQueueByBlockKey(endBlock))
}
