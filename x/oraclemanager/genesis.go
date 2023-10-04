package oraclemanager

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/oraclemanager/keeper"
	"github.com/rarimo/rarimo-core/x/oraclemanager/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	for _, elem := range genState.Oracles {
		k.SetOracle(ctx, elem)
		if elem.Status == types.OracleStatus_Freezed {
			k.AddToFreezedQueue(ctx, elem.FreezeEndBlock, elem.Index)
		}
	}

	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.Oracles = k.GetAllOracle(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
