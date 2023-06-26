package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/vestingmint/types"
)

func (k Keeper) SetVesting(ctx sdk.Context, v types.Vesting) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestingKeyPrefix))
	b := k.cdc.MustMarshal(&v)
	store.Set(types.VestingKey(
		v.Index,
	), b)
}

func (k Keeper) GetVesting(
	ctx sdk.Context,
	index uint64,
) (val types.Vesting, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestingKeyPrefix))

	b := store.Get(types.VestingKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) RemoveVesting(
	ctx sdk.Context,
	index uint64,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestingKeyPrefix))
	store.Delete(types.VestingKey(
		index,
	))
}

// GetAllVesting returns all Vestings
func (k Keeper) GetAllVesting(ctx sdk.Context) (list []types.Vesting) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestingKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Vesting
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) AddToActiveVestingsQueue(ctx sdk.Context, block, index uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ActiveVestingsQueueKey(block, index), []byte{0x1})
}

func (k Keeper) RemoveFromActiveVestingsQueue(ctx sdk.Context, block, index uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.ActiveVestingsQueueKey(block, index))
}

func (k Keeper) ActiveVestingsIterator(ctx sdk.Context, block uint64) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.ActiveVestingsQueueBlockKey(block))
}

func (k Keeper) IterateOverActiveVestingsQueue(ctx sdk.Context, block uint64, cb func(vesting types.Vesting) (stop bool)) {
	iterator := k.ActiveVestingsIterator(ctx, block)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		index := types.GetVestingIndexFromQueueKey(iterator.Key())
		vesting, found := k.GetVesting(ctx, index)
		if !found {
			panic(fmt.Sprintf("operation %d does not exist", index))
		}

		if cb(vesting) {
			break
		}
	}
}
