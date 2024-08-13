package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	rarimocoremoduletypes "github.com/rarimo/rarimo-core/x/rarimocore/types"
	"github.com/rarimo/rarimo-core/x/rootupdater/types"
)

func (k Keeper) EndBlocker(ctx sdk.Context) {
	params := k.GetParams(ctx)

	if params.Root == params.LastRoot {
		return
	}
	// Creating operation to be signed by TSS parties
	index, err := k.rarimo.CreateRootUpdateOperation(ctx, types.ModuleName, &rarimocoremoduletypes.RootUpdate{
		ContractAddress: params.ContractAddress,
		Root:            params.Root,
	})

	if err != nil {
		k.Logger(ctx).Error("Failed to create identity aggregated transfer: " + err.Error())
		return
	}

	params.LastRoot = params.Root
	params.Index = index
	k.SetParams(ctx, params)
}
