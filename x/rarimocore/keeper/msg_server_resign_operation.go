package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k msgServer) ResignOperation(goCtx context.Context, msg *types.MsgResignOperation) (*types.MsgResignOperationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	changeOp, ok := k.GetOperation(ctx, msg.ChangeOpIndex)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("operation %s not found", msg.ChangeOpIndex))
	}

	if changeOp.OperationType != types.OpType_CHANGE_PARTIES {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("operation %s is not %s", changeOp.Index, types.OpType_CHANGE_PARTIES.String()))
	}

	op, ok := k.GetOperation(ctx, msg.OpIndex)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("operation %s not found", msg.OpIndex))
	}

	params := k.GetParams(ctx)

	if op.Timestamp < changeOp.Timestamp-params.AvailableResignBlockDelta || op.Timestamp > changeOp.Timestamp {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("can not resign operation: invalid timestamp"))
	}

	op.Signed = false
	k.SetOperation(ctx, op)
	return &types.MsgResignOperationResponse{}, nil
}
