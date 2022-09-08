package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

// SetDeposit set a specific deposit in the store from its index
func (k Keeper) SetDeposit(ctx sdk.Context, deposit types.Deposit) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositKeyPrefix))
	b := k.cdc.MustMarshal(&deposit)
	store.Set(types.DepositKey(
		deposit.Index,
	), b)
}

// GetDeposit returns a deposit from its index
func (k Keeper) GetDeposit(
	ctx sdk.Context,
	index string,
) (val types.Deposit, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositKeyPrefix))

	b := store.Get(types.DepositKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDeposit removes a deposit from the store
func (k Keeper) RemoveDeposit(
	ctx sdk.Context,
	index string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositKeyPrefix))
	store.Delete(types.DepositKey(
		index,
	))
}

// GetAllDeposit returns all deposit
func (k Keeper) GetAllDeposit(ctx sdk.Context) (list []types.Deposit) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Deposit
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
