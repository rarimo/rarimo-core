package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/cscalist/types"
)

func (k Keeper) EditCSCAListProposal(ctx sdk.Context, proposal *types.EditCSCAListProposal) error {
	params := k.GetParams(ctx)
	k.RemoveTree(ctx) // safe while no errors are expected, otherwise think about a backup

	tree := Treap{k}
	for _, leaf := range proposal.GetLeaves() {
		tree.Insert(ctx, leaf)
	}

	k.SetParams(ctx, params)
	return nil
}
