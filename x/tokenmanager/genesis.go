package tokenmanager

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/keeper"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the info
	for _, collection := range genState.GetCollections() {
		k.SetCollection(ctx, collection)
	}

	for _, item := range genState.Items {
		k.SetItem(ctx, item)
	}

	for _, data := range genState.Datas {
		k.SetCollectionData(ctx, data)
	}

	for _, item := range genState.OnChainItems {
		k.SetOnChainItem(ctx, item)
	}

	for _, seed := range genState.Seeds {
		k.SetSeed(ctx, seed)
	}

	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.Collections = k.GetAllCollections(ctx)
	genesis.Datas = k.GetAllCollectionData(ctx)
	genesis.Items = k.GetAllItem(ctx)
	genesis.OnChainItems = k.GetAllOnChainItem(ctx)
	genesis.Seeds = k.GetAllSeed(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
