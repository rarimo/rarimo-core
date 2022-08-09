package rarimocore

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"gitlab.com/rarify-protocol/rarimo-core/testutil/sample"
	rarimocoresimulation "gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/simulation"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = rarimocoresimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateDeposit = "op_weight_msg_deposit"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDeposit int = 100

	opWeightMsgUpdateDeposit = "op_weight_msg_deposit"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateDeposit int = 100

	opWeightMsgDeleteDeposit = "op_weight_msg_deposit"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteDeposit int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	rarimocoreGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		DepositList: []types.Deposit{
			{
				Creator: sample.AccAddress(),
				Tx:      "0",
			},
			{
				Creator: sample.AccAddress(),
				Tx:      "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&rarimocoreGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateDeposit int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateDeposit, &weightMsgCreateDeposit, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDeposit = defaultWeightMsgCreateDeposit
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDeposit,
		rarimocoresimulation.SimulateMsgCreateDeposit(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateDeposit int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateDeposit, &weightMsgUpdateDeposit, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateDeposit = defaultWeightMsgUpdateDeposit
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateDeposit,
		rarimocoresimulation.SimulateMsgUpdateDeposit(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteDeposit int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteDeposit, &weightMsgDeleteDeposit, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteDeposit = defaultWeightMsgDeleteDeposit
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteDeposit,
		rarimocoresimulation.SimulateMsgDeleteDeposit(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
