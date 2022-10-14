package keeper

import (
	"context"

	cosmostypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

func (k msgServer) CreateChangeKeyECDSA(goCtx context.Context, msg *types.MsgCreateChangeKeyECDSA) (*types.MsgCreateChangeKeyECDSAResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.NewKey == k.GetKeyECDSA(ctx) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "key already set")
	}

	var changeOp = types.ChangeKey{
		NewKey: msg.NewKey,
	}

	details, err := cosmostypes.NewAnyWithValue(&changeOp)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error parsing details %s", err.Error())
	}

	var operation = types.Operation{
		Index:         uuid.NewString(),
		OperationType: types.OpType_CHANGE_KEY,
		Details:       details,
		Signed:        false,
		Creator:       msg.Creator,
		Timestamp:     ctx.BlockTime().Unix(),
	}

	k.SetOperation(
		ctx,
		operation,
	)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeNewOperation,
		sdk.NewAttribute(types.AttributeKeyOperationId, operation.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, types.OpType_CHANGE_KEY.String()),
	))

	return &types.MsgCreateChangeKeyECDSAResponse{}, nil
}
