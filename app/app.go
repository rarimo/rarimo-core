package app

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/x/auth/posthandler"
	ante2 "github.com/rarimo/rarimo-core/ethermint/ante"

	ethermint "github.com/rarimo/rarimo-core/ethermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	nodeservice "github.com/cosmos/cosmos-sdk/client/grpc/node"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/cosmos/cosmos-sdk/x/group"
	groupkeeper "github.com/cosmos/cosmos-sdk/x/group/keeper"
	groupmodule "github.com/cosmos/cosmos-sdk/x/group/module"
	"github.com/cosmos/cosmos-sdk/x/mint"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ica "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/controller/keeper"
	icacontrollertypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/controller/types"
	icahost "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/host"
	icahostkeeper "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/host/keeper"
	icahosttypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/types"
	"github.com/cosmos/ibc-go/v6/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v6/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v6/modules/core"
	ibcclient "github.com/cosmos/ibc-go/v6/modules/core/02-client"
	ibcclientclient "github.com/cosmos/ibc-go/v6/modules/core/02-client/client"
	ibcclienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	ibcporttypes "github.com/cosmos/ibc-go/v6/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/v6/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/v6/modules/core/keeper"
	appparams "github.com/rarimo/rarimo-core/app/params"
	"github.com/rarimo/rarimo-core/docs"
	srvflags "github.com/rarimo/rarimo-core/ethermint/server/flags"
	bridgemodule "github.com/rarimo/rarimo-core/x/bridge"
	bridgemodulekeeper "github.com/rarimo/rarimo-core/x/bridge/keeper"
	bridgemoduletypes "github.com/rarimo/rarimo-core/x/bridge/types"
	cscalistmodule "github.com/rarimo/rarimo-core/x/cscalist"
	cscalistkeeper "github.com/rarimo/rarimo-core/x/cscalist/keeper"
	cscalisttypes "github.com/rarimo/rarimo-core/x/cscalist/types"
	"github.com/rarimo/rarimo-core/x/evm"
	evmkeeper "github.com/rarimo/rarimo-core/x/evm/keeper"
	evmtypes "github.com/rarimo/rarimo-core/x/evm/types"
	"github.com/rarimo/rarimo-core/x/evm/vm/geth"
	"github.com/rarimo/rarimo-core/x/feemarket"
	feemarketkeeper "github.com/rarimo/rarimo-core/x/feemarket/keeper"
	feemarkettypes "github.com/rarimo/rarimo-core/x/feemarket/types"
	identitymodule "github.com/rarimo/rarimo-core/x/identity"
	identitymodulekeeper "github.com/rarimo/rarimo-core/x/identity/keeper"
	identitymoduletypes "github.com/rarimo/rarimo-core/x/identity/types"
	multisigmodule "github.com/rarimo/rarimo-core/x/multisig"
	multisigmodulekeeper "github.com/rarimo/rarimo-core/x/multisig/keeper"
	multisigmoduletypes "github.com/rarimo/rarimo-core/x/multisig/types"
	oraclemanagermodule "github.com/rarimo/rarimo-core/x/oraclemanager"
	oraclemanagermodulekeeper "github.com/rarimo/rarimo-core/x/oraclemanager/keeper"
	oraclemanagermoduletypes "github.com/rarimo/rarimo-core/x/oraclemanager/types"
	rarimocoremodule "github.com/rarimo/rarimo-core/x/rarimocore"
	rarimocoremodulekeeper "github.com/rarimo/rarimo-core/x/rarimocore/keeper"
	rarimocoremoduletypes "github.com/rarimo/rarimo-core/x/rarimocore/types"
	rootupdatermodule "github.com/rarimo/rarimo-core/x/rootupdater"
	rootupdatermodulekeeper "github.com/rarimo/rarimo-core/x/rootupdater/keeper"
	rootupdatermoduletypes "github.com/rarimo/rarimo-core/x/rootupdater/types"
	tokenmanagermodule "github.com/rarimo/rarimo-core/x/tokenmanager"
	tokenmanagermodulekeeper "github.com/rarimo/rarimo-core/x/tokenmanager/keeper"
	tokenmanagermoduletypes "github.com/rarimo/rarimo-core/x/tokenmanager/types"
	vestingmintmodule "github.com/rarimo/rarimo-core/x/vestingmint"
	vestingmintmodulekeeper "github.com/rarimo/rarimo-core/x/vestingmint/keeper"
	vestingmintmoduletypes "github.com/rarimo/rarimo-core/x/vestingmint/types"
	"github.com/spf13/cast"
	abci "github.com/tendermint/tendermint/abci/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"
	// this line is used by starport scaffolding # stargate/app/moduleImport
)

const (
	AccountAddressPrefix = "rarimo"
	Name                 = "rarimo-core"
)

// this line is used by starport scaffolding # stargate/wasm/app/enabledProposals

func getGovProposalHandlers() []govclient.ProposalHandler {
	var govProposalHandlers []govclient.ProposalHandler
	// this line is used by starport scaffolding # stargate/app/govProposalHandlers

	govProposalHandlers = append(govProposalHandlers,
		paramsclient.ProposalHandler,
		distrclient.ProposalHandler,
		upgradeclient.LegacyProposalHandler,
		upgradeclient.LegacyCancelProposalHandler,
		ibcclientclient.UpdateClientProposalHandler,
		ibcclientclient.UpgradeProposalHandler,
		// this line is used by starport scaffolding # stargate/app/govProposalHandler
	)

	return govProposalHandlers
}

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(getGovProposalHandlers()),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		groupmodule.AppModuleBasic{},
		ibc.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		transfer.AppModuleBasic{},
		ica.AppModuleBasic{},
		vesting.AppModuleBasic{},
		rarimocoremodule.AppModuleBasic{},
		tokenmanagermodule.AppModuleBasic{},
		bridgemodule.AppModuleBasic{},
		oraclemanagermodule.AppModuleBasic{},
		multisigmodule.AppModuleBasic{},
		// Ethermint modules
		evm.AppModuleBasic{},
		feemarket.AppModuleBasic{},

		vestingmintmodule.AppModuleBasic{},
		identitymodule.AppModuleBasic{},
		cscalistmodule.AppModuleBasic{},
		rootupdatermodule.AppModuleBasic{},

		// this line is used by starport scaffolding # stargate/app/moduleBasic
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:          nil,
		distrtypes.ModuleName:               nil,
		icatypes.ModuleName:                 nil,
		minttypes.ModuleName:                {authtypes.Minter},
		stakingtypes.BondedPoolName:         {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName:      {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:                 {authtypes.Burner},
		ibctransfertypes.ModuleName:         {authtypes.Minter, authtypes.Burner},
		rarimocoremoduletypes.ModuleName:    nil,
		oraclemanagermoduletypes.ModuleName: nil,
		bridgemoduletypes.ModuleName:        {authtypes.Minter, authtypes.Burner},
		multisigmoduletypes.ModuleName:      nil,
		evmtypes.ModuleName:                 {authtypes.Minter, authtypes.Burner}, // used for secure addition and subtraction of balance using module account
		vestingmintmoduletypes.ModuleName:   {authtypes.Minter},
		// this line is used by starport scaffolding # stargate/app/maccPerms
	}
)

var (
	_ servertypes.Application = (*App)(nil)
	_ simapp.App              = (*App)(nil)
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+Name)
}

// App extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type App struct {
	*baseapp.BaseApp

	cdc               *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// keepers
	AccountKeeper    authkeeper.AccountKeeper
	AuthzKeeper      authzkeeper.Keeper
	BankKeeper       bankkeeper.Keeper
	CapabilityKeeper *capabilitykeeper.Keeper
	StakingKeeper    stakingkeeper.Keeper
	SlashingKeeper   slashingkeeper.Keeper
	MintKeeper       mintkeeper.Keeper
	DistrKeeper      distrkeeper.Keeper
	GovKeeper        govkeeper.Keeper
	CrisisKeeper     crisiskeeper.Keeper
	UpgradeKeeper    upgradekeeper.Keeper
	ParamsKeeper     paramskeeper.Keeper
	IBCKeeper        *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	EvidenceKeeper   evidencekeeper.Keeper
	TransferKeeper   ibctransferkeeper.Keeper
	ICAHostKeeper    icahostkeeper.Keeper
	FeeGrantKeeper   feegrantkeeper.Keeper
	GroupKeeper      groupkeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper      capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper capabilitykeeper.ScopedKeeper
	ScopedICAHostKeeper  capabilitykeeper.ScopedKeeper

	VestingmintKeeper vestingmintmodulekeeper.Keeper
	IdentityKeeper    identitymodulekeeper.Keeper
	CSCAListKeeper    cscalistkeeper.Keeper

	RootupdaterKeeper rootupdatermodulekeeper.Keeper

	// this line is used by starport scaffolding # stargate/app/keeperDeclaration
	RarimocoreKeeper rarimocoremodulekeeper.Keeper

	TokenmanagerKeeper tokenmanagermodulekeeper.Keeper

	BridgeKeeper bridgemodulekeeper.Keeper

	OraclemanagerKeeper oraclemanagermodulekeeper.Keeper

	MultisigKeeper multisigmodulekeeper.Keeper

	// Ethermint keepers
	EvmKeeper       *evmkeeper.Keeper
	FeeMarketKeeper feemarketkeeper.Keeper

	// this line is used by starport scaffolding # stargate/app/keeperDeclaration

	// mm is the module manager
	mm *module.Manager

	// sm is the simulation manager
	sm           *module.SimulationManager
	configurator module.Configurator
}

// New returns a reference to an initialized blockchain app
func New(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	encodingConfig appparams.EncodingConfig,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *App {
	appCodec := encodingConfig.Marshaler
	cdc := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	bApp := baseapp.NewBaseApp(
		Name,
		logger,
		db,
		encodingConfig.TxConfig.TxDecoder(),
		baseAppOptions...,
	)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := sdk.NewKVStoreKeys(
		authtypes.StoreKey, authz.ModuleName, banktypes.StoreKey, stakingtypes.StoreKey,
		minttypes.StoreKey, distrtypes.StoreKey, slashingtypes.StoreKey, govtypes.StoreKey,
		paramstypes.StoreKey, ibchost.StoreKey, upgradetypes.StoreKey, feegrant.StoreKey, evidencetypes.StoreKey,
		ibctransfertypes.StoreKey, icahosttypes.StoreKey, capabilitytypes.StoreKey, group.StoreKey,
		icacontrollertypes.StoreKey, rarimocoremoduletypes.StoreKey,
		tokenmanagermoduletypes.StoreKey, bridgemoduletypes.StoreKey, oraclemanagermoduletypes.StoreKey,
		multisigmoduletypes.StoreKey, evmtypes.StoreKey, feemarkettypes.StoreKey,
		vestingmintmoduletypes.StoreKey,
		identitymoduletypes.StoreKey,
		cscalisttypes.StoreKey,
		rootupdatermoduletypes.StoreKey,

		// this line is used by starport scaffolding # stargate/app/storeKey
	)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey, evmtypes.TransientKey, feemarkettypes.TransientKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	app := &App{
		BaseApp:           bApp,
		cdc:               cdc,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	app.ParamsKeeper = initParamsKeeper(
		appCodec,
		cdc,
		keys[paramstypes.StoreKey],
		tkeys[paramstypes.TStoreKey],
	)

	// set the BaseApp's parameter store
	bApp.SetParamStore(app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramstypes.ConsensusParamsKeyTable()))

	// add capability keeper and ScopeToModule for ibc module
	app.CapabilityKeeper = capabilitykeeper.NewKeeper(
		appCodec,
		keys[capabilitytypes.StoreKey],
		memKeys[capabilitytypes.MemStoreKey],
	)

	// grant capabilities for the ibc and ibc-transfer modules
	scopedIBCKeeper := app.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
	scopedICAControllerKeeper := app.CapabilityKeeper.ScopeToModule(icacontrollertypes.SubModuleName)
	scopedTransferKeeper := app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	scopedICAHostKeeper := app.CapabilityKeeper.ScopeToModule(icahosttypes.SubModuleName)
	// this line is used by starport scaffolding # stargate/app/scopedKeeper

	// add keepers
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		keys[authtypes.StoreKey],
		app.GetSubspace(authtypes.ModuleName),
		ethermint.ProtoAccount,
		maccPerms,
		sdk.Bech32PrefixAccAddr,
	)

	app.AuthzKeeper = authzkeeper.NewKeeper(
		keys[authz.ModuleName],
		appCodec,
		app.MsgServiceRouter(),
		app.AccountKeeper,
	)

	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		keys[banktypes.StoreKey],
		app.AccountKeeper,
		app.GetSubspace(banktypes.ModuleName),
		app.BlockedModuleAccountAddrs(),
	)

	app.StakingKeeper = stakingkeeper.NewKeeper(
		appCodec,
		keys[stakingtypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		app.GetSubspace(stakingtypes.ModuleName),
	)

	app.MintKeeper = mintkeeper.NewKeeper(
		appCodec,
		keys[minttypes.StoreKey],
		app.GetSubspace(minttypes.ModuleName),
		&app.StakingKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		authtypes.FeeCollectorName,
	)

	app.DistrKeeper = distrkeeper.NewKeeper(
		appCodec,
		keys[distrtypes.StoreKey],
		app.GetSubspace(distrtypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		&app.StakingKeeper,
		authtypes.FeeCollectorName,
	)

	app.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec,
		keys[slashingtypes.StoreKey],
		&app.StakingKeeper,
		app.GetSubspace(slashingtypes.ModuleName),
	)

	app.CrisisKeeper = crisiskeeper.NewKeeper(
		app.GetSubspace(crisistypes.ModuleName),
		invCheckPeriod,
		app.BankKeeper,
		authtypes.FeeCollectorName,
	)

	groupConfig := group.DefaultConfig()
	/*
		Example of setting group params:
		groupConfig.MaxMetadataLen = 1000
	*/
	app.GroupKeeper = groupkeeper.NewKeeper(
		keys[group.StoreKey],
		appCodec,
		app.MsgServiceRouter(),
		app.AccountKeeper,
		groupConfig,
	)

	app.FeeGrantKeeper = feegrantkeeper.NewKeeper(
		appCodec,
		keys[feegrant.StoreKey],
		app.AccountKeeper,
	)

	app.UpgradeKeeper = upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		keys[upgradetypes.StoreKey],
		appCodec,
		homePath,
		app.BaseApp,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	// ... other modules keepers

	// Create IBC Keeper
	app.IBCKeeper = ibckeeper.NewKeeper(
		appCodec, keys[ibchost.StoreKey],
		app.GetSubspace(ibchost.ModuleName),
		app.StakingKeeper,
		app.UpgradeKeeper,
		scopedIBCKeeper,
	)

	app.TokenmanagerKeeper = *tokenmanagermodulekeeper.NewKeeper(
		appCodec,
		keys[tokenmanagermoduletypes.StoreKey],
		keys[tokenmanagermoduletypes.MemStoreKey],
		nil, // will be set up later (after rarimocore keeper initialization)
	)
	tokenmanagerModule := tokenmanagermodule.NewAppModule(appCodec, app.TokenmanagerKeeper, app.AccountKeeper, app.BankKeeper) // <- ACTUAL module creation in app.go that you need

	// should be initialized before adding to the gov route
	app.RarimocoreKeeper = *rarimocoremodulekeeper.NewKeeper(
		appCodec,
		keys[rarimocoremoduletypes.StoreKey],
		keys[rarimocoremoduletypes.MemStoreKey],
		&app.TokenmanagerKeeper,
		&app.StakingKeeper,
		app.BankKeeper,
	)
	rarimocoreModule := rarimocoremodule.NewAppModule(appCodec, app.RarimocoreKeeper, app.AccountKeeper, app.BankKeeper) // <- ACTUAL module creation in app.go that you need

	app.TokenmanagerKeeper.SetRarimoCore(app.RarimocoreKeeper)

	app.IdentityKeeper = *identitymodulekeeper.NewKeeper(
		appCodec,
		keys[identitymoduletypes.StoreKey],
		keys[identitymoduletypes.MemStoreKey],
		app.RarimocoreKeeper,
	)
	identityModule := identitymodule.NewAppModule(appCodec, app.IdentityKeeper, app.AccountKeeper, app.BankKeeper)

	app.CSCAListKeeper = *cscalistkeeper.NewKeeper(
		appCodec,
		keys[cscalisttypes.StoreKey],
		keys[cscalisttypes.MemStoreKey],
		app.RarimocoreKeeper,
	)
	cscaListModule := cscalistmodule.NewAppModule(appCodec, app.CSCAListKeeper)

	// Create Transfer Keepers
	app.TransferKeeper = ibctransferkeeper.NewKeeper(
		appCodec,
		keys[ibctransfertypes.StoreKey],
		app.GetSubspace(ibctransfertypes.ModuleName),
		app.IBCKeeper.ChannelKeeper,
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		scopedTransferKeeper,
	)
	transferModule := transfer.NewAppModule(app.TransferKeeper) // <- ACTUAL module creation in app.go that you need
	transferIBCModule := transfer.NewIBCModule(app.TransferKeeper)

	app.ICAHostKeeper = icahostkeeper.NewKeeper(
		appCodec, keys[icahosttypes.StoreKey],
		app.GetSubspace(icahosttypes.SubModuleName),
		app.IBCKeeper.ChannelKeeper,
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		scopedICAHostKeeper,
		app.MsgServiceRouter(),
	)
	icaControllerKeeper := icacontrollerkeeper.NewKeeper(
		appCodec, keys[icacontrollertypes.StoreKey],
		app.GetSubspace(icacontrollertypes.SubModuleName),
		app.IBCKeeper.ChannelKeeper, // may be replaced with middleware such as ics29 fee
		app.IBCKeeper.ChannelKeeper, &app.IBCKeeper.PortKeeper,
		scopedICAControllerKeeper, app.MsgServiceRouter(),
	)
	icaModule := ica.NewAppModule(&icaControllerKeeper, &app.ICAHostKeeper) // <- ACTUAL module creation in app.go that you need
	icaHostIBCModule := icahost.NewIBCModule(app.ICAHostKeeper)

	// Create evidence Keeper for to register the IBC light client misbehaviour evidence route
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec,
		keys[evidencetypes.StoreKey],
		&app.StakingKeeper,
		app.SlashingKeeper,
	)
	// If evidence needs to be handled for the app, set routes in router here and seal
	app.EvidenceKeeper = *evidenceKeeper

	govRouter := govv1beta1.NewRouter()
	govRouter.
		AddRoute(govtypes.RouterKey, govv1beta1.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
		AddRoute(distrtypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(app.DistrKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(app.IBCKeeper.ClientKeeper)).
		AddRoute(rarimocoremoduletypes.RouterKey, rarimocoremodule.NewProposalHandler(app.RarimocoreKeeper)).
		AddRoute(tokenmanagermoduletypes.RouterKey, tokenmanagermodule.NewProposalHandler(app.TokenmanagerKeeper)).
		//AddRoute(banktypes.RouterKey, bank.NewProposalHandler(app.BankKeeper)).
		AddRoute(oraclemanagermoduletypes.RouterKey, oraclemanagermodule.NewProposalHandler(app.OraclemanagerKeeper)).
		AddRoute(bridgemoduletypes.RouterKey, bridgemodule.NewProposalHandler(app.BridgeKeeper)).
		AddRoute(cscalisttypes.RouterKey, cscalistmodule.NewProposalHandler(app.CSCAListKeeper))

	govConfig := govtypes.DefaultConfig()
	app.GovKeeper = govkeeper.NewKeeper(
		appCodec,
		keys[govtypes.StoreKey],
		app.GetSubspace(govtypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		&app.StakingKeeper,
		govRouter,
		app.MsgServiceRouter(),
		govConfig,
	)

	app.OraclemanagerKeeper = *oraclemanagermodulekeeper.NewKeeper(
		appCodec,
		keys[oraclemanagermoduletypes.StoreKey],
		keys[oraclemanagermoduletypes.MemStoreKey],
		&app.RarimocoreKeeper,
		app.BankKeeper,
		app.AccountKeeper,
	)
	oraclemanagerModule := oraclemanagermodule.NewAppModule(appCodec, app.OraclemanagerKeeper, app.AccountKeeper, app.BankKeeper)

	app.BridgeKeeper = *bridgemodulekeeper.NewKeeper(
		appCodec,
		keys[bridgemoduletypes.StoreKey],
		keys[bridgemoduletypes.MemStoreKey],
		app.BankKeeper,
		app.RarimocoreKeeper,
		app.TokenmanagerKeeper,
	)
	bridgeModule := bridgemodule.NewAppModule(appCodec, app.BridgeKeeper, app.AccountKeeper, app.BankKeeper) // <- ACTUAL module creation in app.go that you need

	app.MultisigKeeper = *multisigmodulekeeper.NewKeeper(
		appCodec,
		keys[multisigmoduletypes.StoreKey],
		keys[multisigmoduletypes.MemStoreKey],
		app.MsgServiceRouter(),
		app.AccountKeeper,
	)
	multisigModule := multisigmodule.NewAppModule(appCodec, app.MultisigKeeper, app.AccountKeeper) // <- ACTUAL module creation in app.go that you need

	// Create Ethermint keepers
	feeMarketSs := app.GetSubspace(feemarkettypes.ModuleName)
	app.FeeMarketKeeper = feemarketkeeper.NewKeeper(
		appCodec, authtypes.NewModuleAddress(govtypes.ModuleName),
		keys[feemarkettypes.StoreKey], tkeys[feemarkettypes.TransientKey], feeMarketSs,
	)
	feeMarketModule := feemarket.NewAppModule(app.FeeMarketKeeper, feeMarketSs) // <- ACTUAL module creation in app.go that you need

	app.RootupdaterKeeper = *rootupdatermodulekeeper.NewKeeper(
		appCodec,
		keys[rootupdatermoduletypes.StoreKey],
		keys[rootupdatermoduletypes.MemStoreKey],
		app.GetSubspace(rootupdatermoduletypes.ModuleName),
		app.RarimocoreKeeper,
	)
	rootupdaterModule := rootupdatermodule.NewAppModule(appCodec, app.RootupdaterKeeper)

	// Set authority to x/gov module account to only expect the module account to update params
	evmSs := app.GetSubspace(evmtypes.ModuleName)
	app.EvmKeeper = evmkeeper.NewKeeper(
		appCodec, keys[evmtypes.StoreKey], tkeys[evmtypes.TransientKey], authtypes.NewModuleAddress(govtypes.ModuleName),
		app.AccountKeeper, app.BankKeeper, app.StakingKeeper, app.FeeMarketKeeper,
		nil, geth.NewEVM, cast.ToString(appOpts.Get(srvflags.EVMTracer)), evmSs,
	)
	app.EvmKeeper = app.EvmKeeper.SetHooks(evmkeeper.NewMultiEvmHooks(
		app.IdentityKeeper,
		app.RootupdaterKeeper,
	))

	evmModule := evm.NewAppModule(app.EvmKeeper, app.AccountKeeper, evmSs) // <- ACTUAL module creation in app.go that you need

	app.VestingmintKeeper = *vestingmintmodulekeeper.NewKeeper(
		appCodec,
		keys[vestingmintmoduletypes.StoreKey],
		keys[vestingmintmoduletypes.MemStoreKey],
		app.BankKeeper,
	)
	vestingmintModule := vestingmintmodule.NewAppModule(appCodec, app.VestingmintKeeper, app.AccountKeeper, app.BankKeeper)

	// this line is used by starport scaffolding # stargate/app/keeperDefinition

	/**** IBC Routing ****/

	// Sealing prevents other modules from creating scoped sub-keepers
	app.CapabilityKeeper.Seal()

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := ibcporttypes.NewRouter()
	ibcRouter.AddRoute(icahosttypes.SubModuleName, icaHostIBCModule).
		AddRoute(ibctransfertypes.ModuleName, transferIBCModule)
	// this line is used by starport scaffolding # ibc/app/router
	app.IBCKeeper.SetRouter(ibcRouter)

	/**** Module Hooks ****/

	// register hooks after all modules have been initialized

	app.StakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(
			// insert staking hooks receivers here
			app.DistrKeeper.Hooks(),
			app.SlashingKeeper.Hooks(),
		),
	)

	app.GovKeeper.SetHooks(
		govtypes.NewMultiGovHooks(
		// insert governance hooks receivers here
		),
	)

	/**** Module Options ****/

	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	skipGenesisInvariants := cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.

	app.mm = module.NewManager(
		genutil.NewAppModule( // <- ACTUAL module creation in app.go that you need
			app.AccountKeeper, app.StakingKeeper, app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		auth.NewAppModule(appCodec, app.AccountKeeper, nil),                                                                 // <- ACTUAL module creation in app.go that you need
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),       // <- ACTUAL module creation in app.go that you need
		vesting.NewAppModule(app.AccountKeeper, app.BankKeeper),                                                             // <- ACTUAL module creation in app.go that you need
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),                                                      // <- ACTUAL module creation in app.go that you need
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),                                                            // <- ACTUAL module creation in app.go that you need
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry), // <- ACTUAL module creation in app.go that you need
		groupmodule.NewAppModule(appCodec, app.GroupKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),       // <- ACTUAL module creation in app.go that you need
		crisis.NewAppModule(&app.CrisisKeeper, skipGenesisInvariants),                                                       // <- ACTUAL module creation in app.go that you need
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),                                        // <- ACTUAL module creation in app.go that you need
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper),                                                      // <- ACTUAL module creation in app.go that you need
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),           // <- ACTUAL module creation in app.go that you need
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),                 // <- ACTUAL module creation in app.go that you need
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),                                // <- ACTUAL module creation in app.go that you need
		upgrade.NewAppModule(app.UpgradeKeeper),                                                                             // <- ACTUAL module creation in app.go that you need
		evidence.NewAppModule(app.EvidenceKeeper),                                                                           // <- ACTUAL module creation in app.go that you need
		ibc.NewAppModule(app.IBCKeeper),                                                                                     // <- ACTUAL module creation in app.go that you need
		params.NewAppModule(app.ParamsKeeper),                                                                               // <- ACTUAL module creation in app.go that you need
		transferModule,
		icaModule,
		tokenmanagerModule,
		rarimocoreModule,
		bridgeModule,
		oraclemanagerModule,
		multisigModule,
		feeMarketModule,
		evmModule,
		vestingmintModule,
		identityModule,
		cscaListModule,
		rootupdaterModule,

		// this line is used by starport scaffolding # stargate/app/appModule
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	app.mm.SetOrderBeginBlockers(
		// upgrades should be run first
		upgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		feemarkettypes.ModuleName,
		evmtypes.ModuleName,
		minttypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		govtypes.ModuleName,
		crisistypes.ModuleName,
		ibctransfertypes.ModuleName,
		ibchost.ModuleName,
		icatypes.ModuleName,
		genutiltypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		group.ModuleName,
		paramstypes.ModuleName,
		vestingtypes.ModuleName,
		tokenmanagermoduletypes.ModuleName,
		rarimocoremoduletypes.ModuleName,
		bridgemoduletypes.ModuleName,
		oraclemanagermoduletypes.ModuleName,
		multisigmoduletypes.ModuleName,
		vestingmintmoduletypes.ModuleName,
		identitymoduletypes.ModuleName,
		cscalisttypes.ModuleName,
		rootupdatermoduletypes.ModuleName,

		// this line is used by starport scaffolding # stargate/app/beginBlockers
	)

	app.mm.SetOrderEndBlockers(
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		tokenmanagermoduletypes.ModuleName,
		rarimocoremoduletypes.ModuleName,
		bridgemoduletypes.ModuleName,
		oraclemanagermoduletypes.ModuleName,
		multisigmoduletypes.ModuleName,
		evmtypes.ModuleName,
		feemarkettypes.ModuleName, // NOTE: fee market module must go last in order to retrieve the block gas used for all operational modules.
		ibctransfertypes.ModuleName,
		ibchost.ModuleName,
		icatypes.ModuleName,
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		minttypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		group.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		vestingmintmoduletypes.ModuleName,
		identitymoduletypes.ModuleName,
		cscalisttypes.ModuleName,
		rootupdatermoduletypes.ModuleName,

		// this line is used by starport scaffolding # stargate/app/endBlockers
	)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	app.mm.SetOrderInitGenesis(
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		// evm module denomination is used by the feemarket module, in AnteHandle
		evmtypes.ModuleName,
		// NOTE: feemarket need to be initialized before genutil module:
		// gentx transactions use MinGasPriceDecorator.AnteHandle
		feemarkettypes.ModuleName,
		genutiltypes.ModuleName,
		ibctransfertypes.ModuleName,
		ibchost.ModuleName,
		icatypes.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		group.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		tokenmanagermoduletypes.ModuleName,
		rarimocoremoduletypes.ModuleName,
		bridgemoduletypes.ModuleName,
		oraclemanagermoduletypes.ModuleName,
		multisigmoduletypes.ModuleName,
		// NOTE: crisis module must go at the end to check for invariants on each module
		crisistypes.ModuleName,
		vestingmintmoduletypes.ModuleName,
		identitymoduletypes.ModuleName,
		cscalisttypes.ModuleName,

		rootupdatermoduletypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/initGenesis
	)

	// Uncomment if you want to set a custom migration order here.
	// app.mm.SetOrderMigrations(custom order)

	app.mm.RegisterInvariants(&app.CrisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)

	app.configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterServices(app.configurator)

	// create the simulation manager and define the order of the modules for deterministic simulations
	app.sm = module.NewSimulationManager(
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts),
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		params.NewAppModule(app.ParamsKeeper),
		groupmodule.NewAppModule(appCodec, app.GroupKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		evidence.NewAppModule(app.EvidenceKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		transferModule,
		tokenmanagerModule,
		rarimocoreModule,
		bridgeModule,
		oraclemanagerModule,
		multisigModule,
		evmModule,
		feeMarketModule,
		rootupdaterModule,

		// this line is used by starport scaffolding # stargate/app/appModule
	)
	app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)

	app.setAnteHandler(encodingConfig.TxConfig, cast.ToUint64(appOpts.Get(srvflags.EVMMaxTxGasWanted)))

	// In v0.46, the SDK introduces _postHandlers_. PostHandlers are like
	// antehandlers, but are run _after_ the `runMsgs` execution. They are also
	// defined as a chain, and have the same signature as antehandlers.
	//
	// In baseapp, postHandlers are run in the same store branch as `runMsgs`,
	// meaning that both `runMsgs` and `postHandler` state will be committed if
	// both are successful, and both will be reverted if any of the two fails.
	//
	// The SDK exposes a default empty postHandlers chain.
	//
	// Please note that changing any of the anteHandler or postHandler chain is
	// likely to be a state-machine breaking change, which needs a coordinated
	// upgrade.
	app.setPostHandler()

	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Errorf("failed to read upgrade info from disk: %w", err))
	}

	if upgradeInfo.Name == "v1.0.4" && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storetypes.StoreUpgrades{
			Added: []string{identitymoduletypes.ModuleName},
		}))
	}

	app.UpgradeKeeper.SetUpgradeHandler(
		"v1.0.4",
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			params := app.FeeMarketKeeper.GetParams(ctx)
			params.BaseFee = sdk.NewIntFromUint64(0)
			app.FeeMarketKeeper.SetParams(ctx, params)

			app.IdentityKeeper.SetParams(ctx, identitymoduletypes.Params{
				LcgA:                    1664525,
				LcgB:                    1013904223,
				LcgMod:                  4294967296,
				LcgValue:                12345,
				IdentityContractAddress: "0x",
				ChainName:               "Rarimo",
				GISTHash:                "0x",
				GISTUpdatedTimestamp:    0,
				TreapRootKey:            "0x",
				StatesWaitingForSign:    []string{},
			})

			// Disabling repeated call of InitGenesis for new identity module
			fromVM[identitymoduletypes.ModuleName] = identitymodule.AppModule{}.ConsensusVersion()

			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		},
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		"v1.0.5",
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		},
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		"v1.0.6",
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		},
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		"v1.0.7",
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		},
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		"v1.0.8",
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			params := app.RarimocoreKeeper.GetParams(ctx)
			params.IsUpdateRequired = false
			params.MaxViolationsCount = 1000000000
			for _, p := range params.Parties {
				p.Status = rarimocoremoduletypes.PartyStatus_Active
				p.ViolationsCount = 0
				p.ReportedSessions = []string{}
				p.FreezeEndBlock = 0
			}

			app.RarimocoreKeeper.SetParams(ctx, params)
			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		},
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		"v1.1.0-rc0",
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		},
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		"v1.1.0-rc1",
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		},
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		"v1.1.0-rc2",
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		},
	)

	if upgradeInfo.Name == "v1.1.1-rc1" && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storetypes.StoreUpgrades{
			Added: []string{cscalisttypes.ModuleName},
		}))
	}

	app.UpgradeKeeper.SetUpgradeHandler(
		"v1.1.1-rc1",
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		},
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		"v1.1.1-rc2",
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		},
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		"v1.1.1-rc3",
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		},
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		"v1.1.2-rc0",
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			params := app.FeeMarketKeeper.GetParams(ctx)
			params.BaseFee = sdk.NewIntFromUint64(0)
			params.NoBaseFee = true
			app.FeeMarketKeeper.SetParams(ctx, params)
			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		},
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		"v1.1.3-rc1",
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		},
	)

	if upgradeInfo.Name == "v1.1.4-rc1" && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storetypes.StoreUpgrades{
			Added: []string{rootupdatermoduletypes.ModuleName},
		}))
	}

	app.UpgradeKeeper.SetUpgradeHandler(
		"v1.1.4-rc1",
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			app.RootupdaterKeeper.SetParams(ctx, rootupdatermoduletypes.Params{
				ContractAddress:     "0xBF926a23B4A0bcA301F97Ccd27358b55Dc4C7D3C",
				Root:                "0x00",
				LastSignedRoot:      "",
				LastSignedRootIndex: "",
				EventName:           "RootUpdated",
				RootTimestamp:       1724316208,
			})
			fromVM[rootupdatermoduletypes.ModuleName] = rootupdatermodule.AppModule{}.ConsensusVersion()

			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		},
	)

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}
	}

	app.ScopedIBCKeeper = scopedIBCKeeper
	app.ScopedTransferKeeper = scopedTransferKeeper
	// this line is used by starport scaffolding # stargate/app/beforeInitReturn

	return app
}

func (app *App) setAnteHandler(txConfig client.TxConfig, maxGasWanted uint64) {
	anteHandler, err := ante2.NewAnteHandler(ante2.HandlerOptions{
		AccountKeeper:          app.AccountKeeper,
		BankKeeper:             app.BankKeeper,
		SignModeHandler:        txConfig.SignModeHandler(),
		FeegrantKeeper:         app.FeeGrantKeeper,
		SigGasConsumer:         ante2.DefaultSigVerificationGasConsumer,
		IBCKeeper:              app.IBCKeeper,
		EvmKeeper:              app.EvmKeeper,
		FeeMarketKeeper:        app.FeeMarketKeeper,
		MaxTxGasWanted:         maxGasWanted,
		ExtensionOptionChecker: ethermint.HasDynamicFeeExtensionOption,
		TxFeeChecker:           ante2.NewDynamicFeeChecker(app.EvmKeeper),
		DisabledAuthzMsgs: []string{
			sdk.MsgTypeURL(&evmtypes.MsgEthereumTx{}),
			sdk.MsgTypeURL(&vestingtypes.MsgCreateVestingAccount{}),
		},
	})
	if err != nil {
		panic(err)
	}

	app.SetAnteHandler(anteHandler)
}

func (app *App) setPostHandler() {
	postHandler, err := posthandler.NewPostHandler(
		posthandler.HandlerOptions{},
	)
	if err != nil {
		panic(err)
	}

	app.SetPostHandler(postHandler)
}

// Name returns the name of the App
func (app *App) Name() string { return app.BaseApp.Name() }

// BeginBlocker application updates every begin block
func (app *App) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *App) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// InitChainer application update at chain initialization
func (app *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}
	app.UpgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())
	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height
func (app *App) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *App) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// BlockedModuleAccountAddrs returns all the app's blocked module account
// addresses.
func (app *App) BlockedModuleAccountAddrs() map[string]bool {
	modAccAddrs := app.ModuleAccountAddrs()
	delete(modAccAddrs, authtypes.NewModuleAddress(govtypes.ModuleName).String())

	return modAccAddrs
}

// LegacyAmino returns SimApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) LegacyAmino() *codec.LegacyAmino {
	return app.cdc
}

// AppCodec returns an app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns an InterfaceRegistry
func (app *App) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetKey(storeKey string) *storetypes.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetTKey(storeKey string) *storetypes.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *App) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *App) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register node gRPC service for grpc-gateway.
	nodeservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register grpc-gateway routes for all modules.
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register app's OpenAPI routes.
	docs.RegisterOpenAPIService(Name, apiSvr.Router)
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *App) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *App) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(
		clientCtx,
		app.BaseApp.GRPCQueryRouter(),
		app.interfaceRegistry,
		app.Query,
	)
}

// RegisterNodeService implements the Application.RegisterNodeService method.
func (app *App) RegisterNodeService(clientCtx client.Context) {
	nodeservice.RegisterNodeService(clientCtx, app.GRPCQueryRouter())
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey storetypes.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govv1.ParamKeyTable())
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(ibchost.ModuleName)
	paramsKeeper.Subspace(icacontrollertypes.SubModuleName)
	paramsKeeper.Subspace(icahosttypes.SubModuleName)
	//paramsKeeper.Subspace(monitoringptypes.ModuleName)
	paramsKeeper.Subspace(evmtypes.ModuleName)
	paramsKeeper.Subspace(feemarkettypes.ModuleName)
	paramsKeeper.Subspace(rootupdatermoduletypes.ModuleName)
	// this line is used by starport scaffolding # stargate/app/paramSubspace

	return paramsKeeper
}

// SimulationManager implements the SimulationApp interface
func (app *App) SimulationManager() *module.SimulationManager {
	return app.sm
}
