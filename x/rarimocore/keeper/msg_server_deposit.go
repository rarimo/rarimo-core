package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
	"gitlab.com/rarify-protocol/rarimo-core/x/tokenmanager/saver"
	savermsg "gitlab.com/rarify-protocol/saver-grpc-lib/grpc"
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

	infoRequest := &savermsg.MsgTransactionInfoRequest{
		Hash:    msg.Tx,
		EventId: msg.EventId,
		Type:    networks[msg.FromChain].Types[uint32(msg.TokenType)],
	}

	infoResp, err := saver.GetClient(msg.FromChain).GetDepositInfo(goCtx, infoRequest)
	if err != nil {
		k.Logger(ctx).Error("error calling saver service", err)
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error searching deposit %s", err.Error())
	}

	if _, ok := networks[infoResp.TargetNetwork]; !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "network not found: %s", infoResp.TargetNetwork)
	}

	item, ok := k.tm.GetItem(ctx, infoResp.TokenAddress, infoResp.TokenId, msg.FromChain)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "token not found")
	}

	var deposit = types.Deposit{
		Index:      msg.Index,
		Creator:    msg.Creator,
		Tx:         msg.Tx,
		EventId:    msg.EventId,
		FromChain:  msg.FromChain,
		ToChain:    infoResp.TargetNetwork,
		Receiver:   infoResp.Receiver,
		Amount:     infoResp.Amount,
		Signed:     false,
		TokenIndex: item.Index,
		Timestamp:  uint64(time.Now().UTC().UnixMilli()),
	}

	k.SetDeposit(
		ctx,
		deposit,
	)
	return &types.MsgCreateDepositResponse{}, nil
}
