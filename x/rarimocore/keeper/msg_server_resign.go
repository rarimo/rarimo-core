package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k msgServer) Resign(goCtx context.Context, msg *types.MsgResign) (*types.MsgResignResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	op, ok := k.GetOperation(ctx, msg.Operation)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("operation %s not found", msg.Operation))
	}

	if op.Status != types.OpStatus_SIGNED {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("operation %s is not signed", msg.Operation))
	}

	k.ApproveOperationDefault(ctx, op)
	return &types.MsgResignResponse{}, nil
}
