package keeper

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/rarimo/rarimo-core/x/cscalist/types"
)

func (k Keeper) EditCSCAListProposal(ctx sdk.Context, proposal *types.EditCSCAListProposal) error {
	var (
		tree       = Treap{k}
		oldRootKey = k.GetRootKey(ctx)
		oldRoot, _ = k.GetNode(ctx, oldRootKey)
	)

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

	k.updateParamsOnNeed(ctx, oldRoot)
	return nil
}

func (k Keeper) ReplaceCSCAListProposal(ctx sdk.Context, proposal *types.ReplaceCSCAListProposal) error {
	var (
		tree       = Treap{k}
		oldRootKey = k.GetRootKey(ctx)
		oldRoot, _ = k.GetNode(ctx, oldRootKey)
	)
	k.RemoveTree(ctx) // safe while no errors are expected, otherwise think about a backup

	for _, leaf := range proposal.Leaves {
		tree.Insert(ctx, leaf)
	}

	k.updateParamsOnNeed(ctx, oldRoot)
	return nil
}

func (k Keeper) updateParamsOnNeed(ctx sdk.Context, oldRoot types.Node) {
	newRootKey := k.GetRootKey(ctx)
	newRoot, _ := k.GetNode(ctx, newRootKey)

	if oldRoot.Hash != newRoot.Hash {
		params := k.GetParams(ctx)
		params.RootUpdated = true
		params.UpdatedAtBlock = uint64(ctx.BlockHeight())
		k.SetParams(ctx, params)
	}
}
