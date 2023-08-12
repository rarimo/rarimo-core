package rarimocore

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/keeper"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the deposit
	for _, elem := range genState.OperationList {
		k.SetOperation(ctx, elem)
	}
	// Set all the confirmation
	for _, elem := range genState.ConfirmationList {
		k.SetConfirmation(ctx, elem)

		for _, op := range elem.Indexes {
			k.SetOperationConfirmationId(ctx, op, elem.Root)
		}
	}

	// Set all the vote
	for _, elem := range genState.VoteList {
		k.SetVote(ctx, elem)
	}

	// Set all the vote
	for _, elem := range genState.ViolationReportList {
		k.SetViolationReport(ctx, elem)
	}

	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.OperationList = k.GetAllOperation(ctx)
	genesis.ConfirmationList = k.GetAllConfirmation(ctx)
	genesis.VoteList = k.GetAllVote(ctx)
	genesis.ViolationReportList = k.GetAllViolationReport(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
