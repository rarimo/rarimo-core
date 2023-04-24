package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/oraclemanager/types"
)

func (k Keeper) OracleUnfreezeProposal(ctx sdk.Context, proposal *types.OracleUnfreezeProposal) error {
	oracle, ok := k.GetOracle(ctx, &proposal.Index)
	if !ok {
		return types.ErrOracleNotFound
	}

	if oracle.Status != types.OracleStatus_Freezed {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid oracle status: is not status")
	}

	oracle.ViolationsCount = 0
	oracle.Status = types.OracleStatus_Active
	k.SetOracle(ctx, oracle)

	return nil
}

func (k Keeper) ChangeParamsProposal(ctx sdk.Context, proposal *types.ChangeParamsProposal) error {
	k.SetParams(ctx, proposal.Params)
	return nil
}
