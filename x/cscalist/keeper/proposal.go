package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/cscalist/types"
)

func (k Keeper) EditCSCAListProposal(ctx sdk.Context, proposal *types.EditCSCAListProposal) error {
	params := k.GetParams(ctx)
	k.RemoveTree(ctx) // safe while no errors are expected, otherwise think about a backup

	// TODO @violog Use a different tree implementation to allow passport-related ZKPs
	// Ensure that the tree root always matches the one in related smart-contract
	var (
		tree = Treap{k}
		lcg  = LCG{k}
	)
	for _, leaf := range proposal.Leaves {
		tree.Insert(ctx, leaf, lcg.Next(ctx))
	}

	k.SetParams(ctx, params)
	return nil
}
