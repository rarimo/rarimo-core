package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

// SetOperation set a specific deposit in the store from its index
func (k Keeper) SetOperation(ctx sdk.Context, op types.Operation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OperationKeyPrefix))
	b := k.cdc.MustMarshal(&op)
	store.Set(types.OperationKey(
		op.Index,
	), b)
}

// GetOperation returns an operation from its index
func (k Keeper) GetOperation(
	ctx sdk.Context,
	index string,
) (val types.Operation, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OperationKeyPrefix))

	b := store.Get(types.OperationKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOperation removes an operation from the store
func (k Keeper) RemoveOperation(
	ctx sdk.Context,
	index string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OperationKeyPrefix))
	store.Delete(types.OperationKey(
		index,
	))
}

// GetAllOperation returns all operation
func (k Keeper) GetAllOperation(ctx sdk.Context) (list []types.Operation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OperationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Operation
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) SetOperationConfirmationId(ctx sdk.Context, op, conf string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OperationConfirmationIdKeyPrefix))
	store.Set(types.OperationKey(op), []byte(conf))
}

func (k Keeper) RemoveOperationConfirmationId(ctx sdk.Context, op string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OperationConfirmationIdKeyPrefix))
	store.Delete(types.OperationKey(op))
}

func (k Keeper) GetOperationConfirmationId(ctx sdk.Context, op string) (string, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OperationConfirmationIdKeyPrefix))

	b := store.Get(types.OperationKey(
		op,
	))
	if b == nil {
		return "", false
	}

	return string(b), true
}
