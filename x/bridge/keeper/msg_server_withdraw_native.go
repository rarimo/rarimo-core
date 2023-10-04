package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/bridge/types"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/pkg"
	rarimocoretypes "gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k msgServer) WithdrawNative(goCtx context.Context, msg *types.MsgWithdrawNative) (*types.MsgWithdrawNativeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, _ := sdk.AccAddressFromBech32(msg.Creator)

	// Checking for operation validness

	op, found := k.rarimocoreKeeper.GetOperation(ctx, msg.Origin)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "operation not found (%s)", msg.Origin)
	}

	if op.OperationType != rarimocoretypes.OpType_TRANSFER {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "operation is not a transfer (%s)", msg.Origin)
	}

	if op.Status != rarimocoretypes.OpStatus_SIGNED {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "operation is not signed (%s)", msg.Origin)
	}

	// Checking that current operation has not been executed yet

	if _, found := k.GetHash(ctx, msg.Origin); found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "transfer already withdrawn (%s)", msg.Origin)
	}

	transfer, err := pkg.GetTransfer(op)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to get transfer (%s)", msg.Origin)
	}

	if types.NetworkName != transfer.To.Chain {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "target chain id does not match to the current chain (%s)", msg.Origin)
	}

	amount, ok := sdk.NewIntFromString(transfer.Amount)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "failed to parse amount (%s)", transfer.Amount)
	}

	params := k.GetParams(ctx)

	receiverAddress, err := sdk.AccAddressFromHexUnsafe(getAddressWithoutLeading0x(transfer.Receiver))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to parse receiver address (%s)", transfer.Receiver)
	}

	// Minting native token to the receiver
	err = k.bankKeeper.MintTokens(ctx, receiverAddress, sdk.Coins{{
		Amount: amount,
		Denom:  params.GetWithdrawDenom(),
	}})
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to mint tokens for address (%s)", creatorAddr.String())
	}

	k.SetHash(ctx, types.Hash{Index: msg.Origin})

	return &types.MsgWithdrawNativeResponse{}, nil
}

func getAddressWithoutLeading0x(addr string) string {
	return strings.ReplaceAll(addr, "0x", "")
}
