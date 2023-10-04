package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/oraclemanager/types"
	rarimotypes "github.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k Keeper) AddToMonitorQueue(ctx sdk.Context, block uint64, operationId string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.MonitorOperationQueueKey(operationId, block), []byte{0x1})
}

func (k Keeper) RemoveFromMonitorQueue(ctx sdk.Context, block uint64, operationId string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.MonitorOperationQueueKey(operationId, block))
}

func (k Keeper) IterateOverMonitorQueue(ctx sdk.Context, block uint64, cb func(operation rarimotypes.Operation) (stop bool)) {
	iterator := k.MonitorQueueIterator(ctx, block)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		operationId, _ := types.SplitMonitorOperationQueueKey(iterator.Key())
		proposal, found := k.rarimo.GetOperation(ctx, operationId)
		if !found {
			panic(fmt.Sprintf("operation %s does not exist", operationId))
		}

		if cb(proposal) {
			break
		}
	}
}

func (k Keeper) MonitorQueueIterator(ctx sdk.Context, endBlock uint64) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.MonitorOperationQueueByBlockKey(endBlock))
}
