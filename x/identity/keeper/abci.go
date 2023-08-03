package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/identity/types"
	rarimocoremoduletypes "gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k Keeper) EndBlocker(ctx sdk.Context) {
	params := k.GetParams(ctx)
	if len(params.StatesWaitingForSign) == 0 {
		// Nothing to do
		return
	}

	index, err := k.rarimo.CreateIdentityAggregatedTransferOperation(ctx, types.ModuleName, &rarimocoremoduletypes.IdentityAggregatedTransfer{
		Contract:      params.IdentityContractAddress,
		Chain:         params.ChainName,
		GISTHash:      params.GISTHash,
		StateRootHash: k.GetRootKey(ctx),
		Timestamp:     uint64(ctx.BlockTime().UTC().Unix()),
	})

	if err != nil {
		k.Logger(ctx).Error("Failed to create identity aggregated transfer: " + err.Error())
		return
	}

	for _, id := range params.StatesWaitingForSign {
		state, ok := k.GetStateInfo(ctx, id)
		if ok {
			state.LastUpdateOperationIndex = index
			k.SetStateInfo(ctx, state)
		}
	}

	params.StatesWaitingForSign = []string{}
	k.SetParams(ctx, params)
}
