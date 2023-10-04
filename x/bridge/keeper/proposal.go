package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/bridge/types"
)

func (k Keeper) ChangeParamsProposal(ctx sdk.Context, proposal *types.ChangeParamsProposal) error {
	k.SetParams(ctx, proposal.Params)
	return nil
}
