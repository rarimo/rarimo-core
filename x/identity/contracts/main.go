package contracts

import (
	"embed"

	"github.com/cosmos/cosmos-sdk/codec"
	evmtypes "gitlab.com/rarimo/rarimo-core/x/evm/types"
)

//go:embed accounts/*.json
var Accounts embed.FS

func MustGetEVMAccounts(cdc codec.JSONCodec) []evmtypes.GenesisAccount {
	files, err := Accounts.ReadDir("accounts")
	if err != nil {
		panic(err)
	}

	var result []evmtypes.GenesisAccount

	for _, file := range files {
		val, err := Accounts.ReadFile("accounts/" + file.Name())
		if err != nil {
			panic(err)
		}

		var acc evmtypes.GenesisAccount
		cdc.MustUnmarshalJSON(val, &acc)
		result = append(result, acc)
	}

	return result
}
