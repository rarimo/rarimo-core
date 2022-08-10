package rarimocore

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/keeper"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the deposit
	for _, elem := range genState.DepositList {
		k.SetDeposit(ctx, elem)
	}
	// Set all the confirmation
	for _, elem := range genState.ConfirmationList {
		k.SetConfirmation(ctx, elem)
	}
	// Set all the changeKeyECDSA
	for _, elem := range genState.ChangeKeyECDSAList {
		k.SetChangeKeyECDSA(ctx, elem)
	}
	// Set all the changeKeyEdDSA
	for _, elem := range genState.ChangeKeyEdDSAList {
		k.SetChangeKeyEdDSA(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.DepositList = k.GetAllDeposit(ctx)
	genesis.ConfirmationList = k.GetAllConfirmation(ctx)
	genesis.ChangeKeyECDSAList = k.GetAllChangeKeyECDSA(ctx)
	genesis.ChangeKeyEdDSAList = k.GetAllChangeKeyEdDSA(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
