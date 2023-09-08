package v2

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/identity/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	ctx.Logger().Info(fmt.Sprintf("Performing v0.4.0 %s module migrations", types.ModuleName))

	var states []StateInfo

	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.StateInfoKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val StateInfo
		cdc.MustUnmarshal(iterator.Value(), &val)
		states = append(states, val)
	}

	// Remove all states
	ctx.KVStore(storeKey).Delete(types.KeyPrefix(types.StateInfoKeyPrefix))

	for _, state := range states {
		newState := types.StateInfo{
			Index:                    state.Index,
			Hash:                     state.Hash,
			CreatedAtTimestamp:       state.CreatedAtTimestamp,
			CreatedAtBlock:           uint64(ctx.BlockHeight()),
			LastUpdateOperationIndex: state.LastUpdateOperationIndex,
		}

		store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.StateInfoKeyPrefix))
		b := cdc.MustMarshal(&newState)
		store.Set(types.StateInfoKey(newState.Index), b)
	}

	return nil
}
