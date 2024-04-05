package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/cscalist/types"
)

func (k Keeper) EditCSCAListProposal(ctx sdk.Context, proposal *types.EditCSCAListProposal) error {
	tree := Treap{k}

	for _, leaf := range proposal.GetToRemove() {
		tree.Remove(ctx, leaf)
	}
	for _, leaf := range proposal.GetToAdd() {
		tree.Insert(ctx, leaf)
	}

	return nil
}

func (k Keeper) ReplaceCSCAListProposal(ctx sdk.Context, proposal *types.ReplaceCSCAListProposal) error {
	k.RemoveTree(ctx) // safe while no errors are expected, otherwise think about a backup

	tree := Treap{k}
	for _, leaf := range proposal.GetLeaves() {
		tree.Insert(ctx, leaf)
	}

	return nil
}
