package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/cscalist/types"
	rarimo "github.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k Keeper) EndBlocker(ctx sdk.Context) {
	params := k.GetParams(ctx)
	if !params.RootUpdated {
		return
	}

	root, ok := k.GetNode(ctx, params.RootKey)
	if !ok {
		k.Logger(ctx).Error("Root node not found: ", params.RootKey)
		return
	}

	// Creating operation to be signed by TSS parties
	index, err := k.rarimo.CreateCSCARootUpdateOperation(ctx, types.ModuleName, &rarimo.CSCARootUpdate{
		Root:      root.Hash,
		Timestamp: uint64(ctx.BlockTime().Unix()),
	})
	if err != nil {
		k.Logger(ctx).Error("Failed to create CSCA root update operation: " + err.Error())
		return
	}

	params.RootUpdated = false
	params.LastUpdateOperationIndex = index
	k.SetParams(ctx, params)
}
