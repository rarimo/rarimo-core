package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/rarimo/rarimo-core/x/bridge/types"
	"github.com/rarimo/rarimo-core/x/rarimocore/crypto/pkg"
	rarimocoretypes "github.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k msgServer) WithdrawFee(goCtx context.Context, msg *types.MsgWithdrawFee) (*types.MsgWithdrawFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checking for operation validness

	op, found := k.rarimocoreKeeper.GetOperation(ctx, msg.Origin)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "operation not found (%s)", msg.Origin)
	}

	if op.OperationType != rarimocoretypes.OpType_FEE_TOKEN_MANAGEMENT {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "operation is not a transfer (%s)", msg.Origin)
	}

	if op.Status != rarimocoretypes.OpStatus_SIGNED {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "operation is not signed (%s)", msg.Origin)
	}

	// Checking that current operation has not been executed yet
	if _, found := k.GetHash(ctx, msg.Origin); found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "transfer already withdrawn (%s)", msg.Origin)
	}

	manage, err := pkg.GetFeeTokenManagement(op)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to get transfer (%s)", msg.Origin)
	}

	if manage.Chain != types.NetworkName {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "target chain id does not match to the current chain (%s)", msg.Origin)
	}

	if manage.OpType != rarimocoretypes.FeeTokenManagementType_WITHDRAW_FEE_TOKEN {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid operation type: should be: %s", rarimocoretypes.FeeTokenManagementType_WITHDRAW_FEE_TOKEN.String())
	}

	amount, _ := sdk.NewIntFromString(manage.Token.Amount)
	receiverAddress, err := sdk.AccAddressFromHexUnsafe(getAddressWithoutLeading0x(manage.Receiver))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to parse receiver address (%s)", manage.Receiver)
	}

	// Transferring tokens
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddress, sdk.NewCoins(sdk.NewCoin(k.GetParams(ctx).WithdrawDenom, amount)))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to transfer tokens to address (%s)", receiverAddress.String())
	}

	k.SetHash(ctx, types.Hash{Index: msg.Origin})

	return &types.MsgWithdrawFeeResponse{}, nil
}
