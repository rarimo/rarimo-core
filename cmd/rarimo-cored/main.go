package main

import (
	"os"

	"github.com/ignite/cli/ignite/pkg/cosmoscmd"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/app"
	cmdcfg "gitlab.com/rarimo/rarimo-core/cmd/config"
)

const EnvPrefix = "RARIMOCORE"

func main() {
	setupConfig()
	cmdcfg.RegisterDenoms()
	rootCmd, _ := newRootCmd()

	if err := svrcmd.Execute(rootCmd, EnvPrefix, app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)
		default:
			os.Exit(1)
		}
	}
}

func setupConfig() {
	// set the address prefixes
	config := sdk.GetConfig()
	cmdcfg.SetBip44CoinType(config)
	cosmoscmd.SetPrefixes(app.AccountAddressPrefix) // this seals the config
}
