package identity

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/identity/keeper"
	"gitlab.com/rarimo/rarimo-core/x/identity/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)

	for _, v := range genState.Node {
		k.SetNode(ctx, v)
	}

	for _, v := range genState.StateInfo {
		k.SetStateInfo(ctx, v)
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.Node = k.GetAllNode(ctx)
	genesis.StateInfo = k.GetAllStateInfo(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
