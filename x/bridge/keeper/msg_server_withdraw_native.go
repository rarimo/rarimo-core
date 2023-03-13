package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/bridge/types"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/pkg"
	rarimocoretypes "gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k msgServer) WithdrawNative(goCtx context.Context, msg *types.MsgWithdrawNative) (*types.MsgDepositNativeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	op, found := k.rarimocoreKeeper.GetOperation(ctx, msg.Origin)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "operation not found (%s)", msg.Origin)
	}

	if op.OperationType != rarimocoretypes.OpType_TRANSFER {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "operation is not a transfer (%s)", msg.Origin)
	}

	if op.Status != rarimocoretypes.OpStatus_SIGNED {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "operation is not approved (%s)", msg.Origin)
	}

	if _, found := k.GetHash(ctx, msg.Origin); found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "transfer already withdrawn (%s)", msg.Origin)
	}

	transfer, err := pkg.GetTransfer(op)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to get transfer (%s)", msg.Origin)
	}

	amount, ok := sdk.NewIntFromString(transfer.Amount)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "failed to parse amount (%s)", transfer.Amount)
	}
	// TODO: test this
	denom, err := sdktypes.GetBaseDenom()
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to get base denom")
	}

	err = k.bankKeeper.MintTokens(ctx, creatorAddr, sdk.Coins{{
		Amount: amount,
		Denom:  denom,
	}})
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to mint tokens for address (%s)", creatorAddr.String())
	}

	k.SetHash(ctx, types.Hash{Index: msg.Origin})

	return &types.MsgDepositNativeResponse{}, nil
}
