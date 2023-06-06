package main

import (
	"errors"
	"io"
	"os"

	"gitlab.com/rarimo/rarimo-core/cmd/rarimo-cored/cmd"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/debug"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	sdkserver "github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/ignite/cli/ignite/pkg/cosmoscmd"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	tmcfg "github.com/tendermint/tendermint/config"
	tmlog "github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
	"gitlab.com/rarimo/rarimo-core/app"
	"gitlab.com/rarimo/rarimo-core/app/params"
	"gitlab.com/rarimo/rarimo-core/ethereum/eip712"
	ethermintclient "gitlab.com/rarimo/rarimo-core/ethmintclient"
	"gitlab.com/rarimo/rarimo-core/ethmintcrypto/hd"
	encoding "gitlab.com/rarimo/rarimo-core/ethmintencoding"
	server "gitlab.com/rarimo/rarimo-core/ethmintserver"
	servercfg "gitlab.com/rarimo/rarimo-core/ethmintserver/config"
	srvflags "gitlab.com/rarimo/rarimo-core/ethmintserver/flags"
)

func newRootCmd() (*cobra.Command, params.EncodingConfig) {
	encodingConfig := encoding.MakeConfig(app.ModuleBasics)
	initClientCtx := client.Context{}.
		WithCodec(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(types.AccountRetriever{}).
		WithBroadcastMode(flags.BroadcastBlock).
		WithHomeDir(app.DefaultNodeHome).
		WithKeyringOptions(hd.EthSecp256k1Option()).
		WithViper(EnvPrefix)

	eip712.SetEncodingConfig(encodingConfig)

	rootCmd := &cobra.Command{
		Use:   "rarimocored",
		Short: "RarimoCore Daemon",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			// set the default command outputs
			cmd.SetOut(cmd.OutOrStdout())
			cmd.SetErr(cmd.ErrOrStderr())

			initClientCtx, err := client.ReadPersistentCommandFlags(initClientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			initClientCtx, err = config.ReadFromClientConfig(initClientCtx)
			if err != nil {
				return err
			}

			if err := client.SetCmdClientContextHandler(initClientCtx, cmd); err != nil {
				return err
			}

			// FIXME: replace AttoPhoton with bond denom
			customAppTemplate, customAppConfig := servercfg.AppConfig("stake") // TODO move to constant + do we need to overwrite it? we have default config

			return sdkserver.InterceptConfigsPreRunHandler(cmd, customAppTemplate, customAppConfig, tmcfg.DefaultConfig())
		},
	}

	rootCmd.AddCommand(
		ethermintclient.ValidateChainID(
			genutilcli.InitCmd(app.ModuleBasics, app.DefaultNodeHome),
		),
		genutilcli.CollectGenTxsCmd(banktypes.GenesisBalancesIterator{}, app.DefaultNodeHome),
		genutilcli.MigrateGenesisCmd(), // TODO: shouldn't this include the local app version instead of the SDK?
		genutilcli.GenTxCmd(app.ModuleBasics, encodingConfig.TxConfig, banktypes.GenesisBalancesIterator{}, app.DefaultNodeHome),
		genutilcli.ValidateGenesisCmd(app.ModuleBasics),
		cmd.AddGenesisAccountCmd(app.DefaultNodeHome),
		//tmcli.NewCompletionCmd(rootCmd, true),
		ethermintclient.NewTestnetCmd(app.ModuleBasics, banktypes.GenesisBalancesIterator{}),
		debug.Cmd(),
		config.Cmd(),
	)

	a := appCreator{encodingConfig}

	server.AddCommands(rootCmd, server.NewDefaultStartOptions(a.newApp, app.DefaultNodeHome), a.appExport, addModuleInitFlags)

	rootCmd.AddCommand(
		rpc.StatusCommand(),
		ethermintclient.KeyCommands(app.DefaultNodeHome),
	)

	rootCmd, err := srvflags.AddTxFlags(rootCmd)
	if err != nil {
		panic(err)
	}

	return rootCmd, encodingConfig
}

func addModuleInitFlags(startCmd *cobra.Command) {
	crisis.AddModuleInitFlags(startCmd)
}

// appCreator is an app creator.
type appCreator struct {
	encodingConfig params.EncodingConfig
}

func (c *appCreator) newApp(logger tmlog.Logger, db dbm.DB, traceStore io.Writer, appOpts servertypes.AppOptions) servertypes.Application {
	//	var cache sdk.MultiStorePersistentCache
	//
	//	if cast.ToBool(appOpts.Get(sdkserver.FlagInterBlockCache)) {
	//		cache = store.NewCommitKVStoreCacheManager()
	//	}

	skipUpgradeHeights := make(map[int64]bool)
	for _, h := range cast.ToIntSlice(appOpts.Get(sdkserver.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(h)] = true
	}

	//pruningOpts, err := sdkserver.GetPruningOptionsFromFlags(appOpts)
	//if err != nil {
	//	panic(err)
	//}
	//
	//snapshotDir := filepath.Join(cast.ToString(appOpts.Get(flags.FlagHome)), "data", "snapshots")
	//if err = os.MkdirAll(snapshotDir, os.ModePerm); err != nil {
	//	panic(err)
	//}
	//
	//snapshotDB, err := dbm.NewDB("metadata", sdkserver.GetAppDBBackend(appOpts), snapshotDir)
	//if err != nil {
	//	panic(err)
	//}
	//snapshotStore, err := snapshots.NewStore(snapshotDB, snapshotDir)
	//if err != nil {
	//	panic(err)
	//}
	//
	//snapshotOptions := snapshottypes.NewSnapshotOptions(
	//	cast.ToUint64(appOpts.Get(sdkserver.FlagStateSyncSnapshotInterval)),
	//	cast.ToUint32(appOpts.Get(sdkserver.FlagStateSyncSnapshotKeepRecent)),
	//)

	rarimoApp := app.New(
		logger, db, traceStore, true, skipUpgradeHeights,
		cast.ToString(appOpts.Get(flags.FlagHome)),
		cast.ToUint(appOpts.Get(sdkserver.FlagInvCheckPeriod)),
		c.encodingConfig,
		appOpts,
		//	baseapp.SetPruning(pruningOpts), // TODO figure out where do we need that
		//	baseapp.SetMinGasPrices(cast.ToString(appOpts.Get(sdkserver.FlagMinGasPrices))),
		//	baseapp.SetHaltHeight(cast.ToUint64(appOpts.Get(sdkserver.FlagHaltHeight))),
		//	baseapp.SetHaltTime(cast.ToUint64(appOpts.Get(sdkserver.FlagHaltTime))),
		//	baseapp.SetMinRetainBlocks(cast.ToUint64(appOpts.Get(sdkserver.FlagMinRetainBlocks))),
		//	baseapp.SetInterBlockCache(cache),
		//	baseapp.SetTrace(cast.ToBool(appOpts.Get(sdkserver.FlagTrace))),
		//	baseapp.SetIndexEvents(cast.ToStringSlice(appOpts.Get(sdkserver.FlagIndexEvents))),
		//	baseapp.SetSnapshot(snapshotStore, snapshotOptions),
		//	baseapp.SetIAVLCacheSize(cast.ToInt(appOpts.Get(sdkserver.FlagIAVLCacheSize))),
		//	baseapp.SetIAVLDisableFastNode(cast.ToBool(appOpts.Get(sdkserver.FlagDisableIAVLFastNode))),
	)

	return rarimoApp
}

func (c appCreator) appExport(
	logger tmlog.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailAllowedAddrs []string,
	appOpts servertypes.AppOptions,
) (servertypes.ExportedApp, error) {
	var ethermintApp cosmoscmd.App
	homePath, ok := appOpts.Get(flags.FlagHome).(string)
	if !ok || homePath == "" {
		return servertypes.ExportedApp{}, errors.New("application home not set")
	}

	if height != -1 {
		ethermintApp = app.New(logger, db, traceStore, false, map[int64]bool{}, "", uint(1), c.encodingConfig, appOpts)

		if err := ethermintApp.LoadHeight(height); err != nil {
			return servertypes.ExportedApp{}, err
		}
	} else {
		ethermintApp = app.New(logger, db, traceStore, true, map[int64]bool{}, "", uint(1), c.encodingConfig, appOpts)
	}

	return ethermintApp.ExportAppStateAndValidators(forZeroHeight, jailAllowedAddrs)
}
