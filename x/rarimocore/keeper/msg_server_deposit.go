package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

func (k msgServer) CreateDeposit(goCtx context.Context, msg *types.MsgCreateDeposit) (*types.MsgCreateDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, isFound := k.GetDeposit(
		ctx,
		msg.Tx,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	networks := k.tm.GetParams(ctx).Networks

	if _, ok := networks[msg.FromChain]; !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "network not found: %s", msg.FromChain)
	}

	if _, ok := networks[msg.ToChain]; !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "network not found: %s", msg.ToChain)
	}

	// TODO validate with saver and fill
	var tokenAddress, tokenId, tokenChain = "", "", ""

	item, ok := k.tm.GetItem(ctx, tokenAddress, tokenId, tokenChain)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "token not found")
	}

	var deposit = types.Deposit{
		Index:      msg.Index,
		Creator:    msg.Creator,
		Tx:         msg.Tx,
		EventId:    msg.EventId,
		FromChain:  msg.FromChain,
		ToChain:    msg.ToChain,
		Receiver:   msg.Receiver,
		Amount:     msg.Amount,
		Signed:     false,
		TokenIndex: item.Index,
	}

	k.SetDeposit(
		ctx,
		deposit,
	)
	return &types.MsgCreateDepositResponse{}, nil
}
