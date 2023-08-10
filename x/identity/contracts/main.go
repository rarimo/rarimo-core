package contracts

import (
	"embed"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	evmtypes "gitlab.com/rarimo/rarimo-core/x/evm/types"
)

//go:embed evm/*.json
var EVMAccounts embed.FS

//go:embed auth/*.json
var AuthAccounts embed.FS

func MustGetEVMAccounts(cdc codec.JSONCodec) (res1 []evmtypes.GenesisAccount, res2 []*types.Any) {
	files, err := EVMAccounts.ReadDir("evm")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		val, err := EVMAccounts.ReadFile("evm/" + file.Name())
		if err != nil {
			panic(err)
		}

		var acc evmtypes.GenesisAccount
		cdc.MustUnmarshalJSON(val, &acc)
		res1 = append(res1, acc)
	}

	files, err = AuthAccounts.ReadDir("auth")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		val, err := AuthAccounts.ReadFile("auth/" + file.Name())
		if err != nil {
			panic(err)
		}

		var acc types.Any
		cdc.MustUnmarshalJSON(val, &acc)
		res2 = append(res2, &acc)
	}

	return
}
