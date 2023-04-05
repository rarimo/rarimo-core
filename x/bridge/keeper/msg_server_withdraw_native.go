package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"gitlab.com/rarimo/rarimo-core/x/bridge/types"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/pkg"
	rarimocoretypes "gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k msgServer) WithdrawNative(goCtx context.Context, msg *types.MsgWithdrawNative) (*types.MsgWithdrawNativeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, _ := sdk.AccAddressFromBech32(msg.Creator)

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

	if _, found := k.GetHash(ctx, msg.Origin); found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "transfer already withdrawn (%s)", msg.Origin)
	}

	transfer, err := pkg.GetTransfer(op)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to get transfer (%s)", msg.Origin)
	}

	if ctx.ChainID() != transfer.To.Chain {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "target chain id does not match to the current chain (%s)", msg.Origin)
	}

	amount, ok := sdk.NewIntFromString(transfer.Amount)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "failed to parse amount (%s)", transfer.Amount)
	}

	params := k.GetParams(ctx)

	receiver, err := hexutil.Decode(transfer.Receiver)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to decode receiver (%s)", transfer.Receiver)
	}

	receiverAddress, err := sdk.AccAddressFromBech32(string(receiver))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to parse receiver address (%s)", string(receiver))
	}

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
