package multisig

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/multisig/keeper"
	"github.com/rarimo/rarimo-core/x/multisig/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)

	for _, group := range genState.GroupList {
		k.SetGroup(ctx, group)
	}
	for _, proposal := range genState.ProposalList {
		k.SetProposal(ctx, proposal)
	}
	for _, vote := range genState.VoteList {
		k.SetVote(ctx, vote)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.GroupList = k.GetAllGroup(ctx)
	genesis.ProposalList = k.GetAllProposal(ctx)
	genesis.VoteList = k.GetAllVote(ctx)

	// this line is used by starport scaffolding # genesis/module/export
	return genesis
}
