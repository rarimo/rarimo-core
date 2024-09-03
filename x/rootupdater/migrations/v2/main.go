package v2

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/rootupdater/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	ctx.Logger().Info(fmt.Sprintf("Performing v1.1.4-rc6 %s module migrations", types.ModuleName))

	store := ctx.KVStore(storeKey)

	oldParamsBytes := store.Get(types.KeyPrefix(types.ParamsKey))
	store.Delete(types.KeyPrefix(types.ParamsKey))

	oldParams := OldParamas{}
	cdc.MustUnmarshal(oldParamsBytes, &oldParams)
	params := types.Params{
		DestinationContractAddress: "0x2af05993a27df83094a963af64b5d25296230544",
		SourceContractAddress:      oldParams.ContractAddress,
		Root:                       oldParams.Root,
		RootTimestamp:              oldParams.RootTimestamp,
		LastSignedRootIndex:        oldParams.LastSignedRootIndex,
		LastSignedRoot:             oldParams.LastSignedRoot,
		BlockHeight:                oldParams.BlockHeight,
		EventName:                  oldParams.EventName,
	}

	store.Set(types.KeyPrefix(types.ParamsKey), cdc.MustMarshal(&params))

	return nil
}
