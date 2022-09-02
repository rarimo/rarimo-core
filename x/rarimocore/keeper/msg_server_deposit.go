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

	// TODO validate

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
