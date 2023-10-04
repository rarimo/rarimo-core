package v3

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	ctx.Logger().Info(fmt.Sprintf("Performing v0.0.6 %s module migrations", types.ModuleName))

	params := types.Params{}
	b := ctx.KVStore(storeKey).Get(types.KeyPrefix(types.ParamsKey))
	cdc.MustUnmarshal(b, &params)

	params.KeyECDSA = removeKeyConstantPrefix(params.KeyECDSA)
	for _, party := range params.Parties {
		party.PubKey = removeKeyConstantPrefix(party.PubKey)
	}

	b = cdc.MustMarshal(&params)
	ctx.KVStore(storeKey).Set(types.KeyPrefix(types.ParamsKey), b)
	return nil
}

func removeKeyConstantPrefix(key string) string {
	// Values are valid in params
	if b := hexutil.MustDecode(key); len(b) > 0 && b[0] == 0x04 {
		return hexutil.Encode(b[1:])
	}

	return key
}
