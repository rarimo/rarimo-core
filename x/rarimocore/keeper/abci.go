package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k Keeper) EndBlocker(ctx sdk.Context) {
	param := k.GetParams(ctx)
	currentBlock := uint64(ctx.BlockHeight())

	for _, party := range param.Parties {
		if party.FreezeEndBlock == 0 || party.FreezeEndBlock != currentBlock {
			continue
		}

		party.Status = types.PartyStatus_Slashed
		party.FreezeEndBlock = 0

		k.UpdateParams(ctx, param)
	}
}
