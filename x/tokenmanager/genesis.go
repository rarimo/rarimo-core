package tokenmanager

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarify-protocol/rarimo-core/x/tokenmanager/keeper"
	"gitlab.com/rarify-protocol/rarimo-core/x/tokenmanager/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the item
	for _, elem := range genState.ItemList {
		k.SetItem(ctx, elem)
	}
	// Set all the info
	for _, elem := range genState.InfoList {
		k.SetInfo(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ItemList = k.GetAllItem(ctx)
	genesis.InfoList = k.GetAllInfo(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
