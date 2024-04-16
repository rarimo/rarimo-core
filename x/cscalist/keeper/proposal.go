package keeper

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/rarimo/rarimo-core/x/cscalist/types"
)

func (k Keeper) EditCSCAListProposal(ctx sdk.Context, proposal *types.EditCSCAListProposal) error {
	tree := Treap{k}
	root := k.GetRootKey(ctx)

	for _, leaf := range proposal.ToRemove {
		tree.Remove(ctx, leaf)
	}
	for _, leaf := range proposal.ToAdd {
		_, ok := k.GetNode(ctx, leaf)
		if ok {
			return errors.Wrapf(govtypes.ErrInvalidProposalContent, "node already exists")
		}
		tree.Insert(ctx, leaf)
	}

	if root != k.GetRootKey(ctx) {
		params := k.GetParams(ctx)
		params.RootUpdated = true
		params.UpdatedAtBlock = uint64(ctx.BlockHeight())
		k.SetParams(ctx, params)
	}

	return nil
}

func (k Keeper) ReplaceCSCAListProposal(ctx sdk.Context, proposal *types.ReplaceCSCAListProposal) error {
	tree := Treap{k}
	root := k.GetRootKey(ctx)
	k.RemoveTree(ctx) // safe while no errors are expected, otherwise think about a backup

	for _, leaf := range proposal.Leaves {
		tree.Insert(ctx, leaf)
	}

	if root != k.GetRootKey(ctx) {
		params := k.GetParams(ctx)
		params.RootUpdated = true
		params.UpdatedAtBlock = uint64(ctx.BlockHeight())
		k.SetParams(ctx, params)
	}

	return nil
}
