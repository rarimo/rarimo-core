package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/rarimo/rarimo-core/x/oraclemanager/types"
)

func (k msgServer) Unjail(goCtx context.Context, msg *types.MsgUnjail) (*types.MsgUnjailResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// REQUIRES: validated index
	oracle, ok := k.GetOracle(ctx, msg.Index)
	if !ok {
		return nil, types.ErrOracleNotFound
	}

	if oracle.Status != types.OracleStatus_Jailed {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "can not unjail: is not jailed")
	}

	oracle.Status = types.OracleStatus_Active
	oracle.MissedCount = 0
	k.SetOracle(ctx, oracle)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeOracleActivated,
		sdk.NewAttribute(types.AttributeKeyChain, msg.Index.Chain),
		sdk.NewAttribute(types.AttributeKeyAccount, msg.Index.Account),
	))

	return &types.MsgUnjailResponse{}, nil
}
