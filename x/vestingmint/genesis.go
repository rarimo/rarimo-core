package vestingmint

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/vestingmint/keeper"
	"gitlab.com/rarimo/rarimo-core/x/vestingmint/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)

	for _, v := range genState.Vestings {
		k.SetVesting(ctx, v)
		if v.NextDepositBlock >= uint64(ctx.BlockHeight()) {
			k.AddToActiveVestingsQueue(ctx, v.NextDepositBlock, v.Index)
		}
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.Vestings = k.GetAllVesting(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
