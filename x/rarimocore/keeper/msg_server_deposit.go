package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

func (k msgServer) CreateDeposit(goCtx context.Context, msg *types.MsgCreateDeposit) (*types.MsgCreateDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetDeposit(
		ctx,
		msg.Tx,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var deposit = types.Deposit{
		Creator:      msg.Creator,
		Tx:           msg.Tx,
		FromChain:    msg.FromChain,
		ToChain:      msg.ToChain,
		Receiver:     msg.Receiver,
		TokenAddress: msg.TokenAddress,
		TokenId:      msg.TokenId,
		Signed:       false,
		TokenType:    msg.TokenType,
	}

	k.SetDeposit(
		ctx,
		deposit,
	)
	return &types.MsgCreateDepositResponse{}, nil
}

func (k msgServer) UpdateDeposit(goCtx context.Context, msg *types.MsgUpdateDeposit) (*types.MsgUpdateDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDeposit(
		ctx,
		msg.Tx,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var deposit = types.Deposit{
		Creator:      msg.Creator,
		Tx:           msg.Tx,
		FromChain:    msg.FromChain,
		ToChain:      msg.ToChain,
		Receiver:     msg.Receiver,
		TokenAddress: msg.TokenAddress,
		TokenId:      msg.TokenId,
		Signed:       false,
		TokenType:    msg.TokenType,
	}

	k.SetDeposit(ctx, deposit)

	return &types.MsgUpdateDepositResponse{}, nil
}

func (k msgServer) DeleteDeposit(goCtx context.Context, msg *types.MsgDeleteDeposit) (*types.MsgDeleteDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDeposit(
		ctx,
		msg.Tx,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveDeposit(
		ctx,
		msg.Tx,
	)

	return &types.MsgDeleteDepositResponse{}, nil
}
