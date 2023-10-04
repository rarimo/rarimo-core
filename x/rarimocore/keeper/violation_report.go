package keeper

import (
	"bytes"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/rarimocore/types"
)

// SetViolationReport set a violation report in the store from its index
func (k Keeper) SetViolationReport(ctx sdk.Context, v types.ViolationReport) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ViolationReportKeyPrefix))
	b := k.cdc.MustMarshal(&v)
	store.Set(types.ViolationReportKey(v.Index), b)
}

// GetViolationReport returns a ViolationReport from its index
func (k Keeper) GetViolationReport(
	ctx sdk.Context,
	index *types.ViolationReportIndex,
) (val types.ViolationReport, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ViolationReportKeyPrefix))

	b := store.Get(types.ViolationReportKey(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveViolationReport removes a ViolationReport from the store
func (k Keeper) RemoveViolationReport(
	ctx sdk.Context,
	index *types.ViolationReportIndex,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ViolationReportKeyPrefix))
	store.Delete(types.ViolationReportKey(index))
}

// GetAllViolationReport returns all ViolationReport
func (k Keeper) GetAllViolationReport(ctx sdk.Context) (list []types.ViolationReport) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ViolationReportKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ViolationReport
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) IterateViolationReports(ctx sdk.Context, sessionId, offender string, f func(report types.ViolationReport) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ViolationReportKeyPrefix))
	storePrefix := bytes.Join([][]byte{[]byte(sessionId), []byte(offender)}, []byte("/"))

	iterator := sdk.KVStorePrefixIterator(store, storePrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var report types.ViolationReport
		k.cdc.MustUnmarshal(iterator.Value(), &report)
		if f(report) {
			break
		}
	}
}
