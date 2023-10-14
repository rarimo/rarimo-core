package v4

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/rarimocore/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	ctx.Logger().Info(fmt.Sprintf("Performing v0.0.7 %s module migrations", types.ModuleName))

	params := types.Params{}
	b := ctx.KVStore(storeKey).Get(types.KeyPrefix(types.ParamsKey))
	cdc.MustUnmarshal(b, &params)

	for _, party := range params.Parties {
		party.ReportedSessions = []string{}
	}

	b = cdc.MustMarshal(&params)
	ctx.KVStore(storeKey).Set(types.KeyPrefix(types.ParamsKey), b)
	return nil
}
