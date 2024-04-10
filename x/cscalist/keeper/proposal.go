package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/cscalist/types"
)

func (k Keeper) EditCSCAListProposal(ctx sdk.Context, proposal *types.EditCSCAListProposal) error {
	tree := Treap{k}
	root := k.GetRootKey(ctx)

	for _, leaf := range proposal.GetToRemove() {
		tree.Remove(ctx, leaf)
	}
	for _, leaf := range proposal.GetToAdd() {
		tree.Insert(ctx, leaf)
	}

	if root != k.GetRootKey(ctx) {
		params := k.GetParams(ctx)
		params.RootUpdated = true
		k.SetParams(ctx, params)
	}

	return nil
}

func (k Keeper) ReplaceCSCAListProposal(ctx sdk.Context, proposal *types.ReplaceCSCAListProposal) error {
	tree := Treap{k}
	root := k.GetRootKey(ctx)
	k.RemoveTree(ctx) // safe while no errors are expected, otherwise think about a backup

	for _, leaf := range proposal.GetLeaves() {
		tree.Insert(ctx, leaf)
	}

	if root != k.GetRootKey(ctx) {
		params := k.GetParams(ctx)
		params.RootUpdated = true
		k.SetParams(ctx, params)
	}

	return nil
}
