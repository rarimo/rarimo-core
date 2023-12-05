package oraclemanager

import (
	"fmt"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
	"github.com/rarimo/rarimo-core/x/oraclemanager/keeper"
	"github.com/rarimo/rarimo-core/x/oraclemanager/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		var res proto.Message
		var err error

		switch msg := msg.(type) {
		case *types.MsgStake:
			res, err = msgServer.Stake(sdk.WrapSDKContext(ctx), msg)
		case *types.MsgUnstake:
			res, err = msgServer.Unstake(sdk.WrapSDKContext(ctx), msg)
		case *types.MsgUnjail:
			res, err = msgServer.Unjail(sdk.WrapSDKContext(ctx), msg)
		case *types.MsgCreateTransferOp:
			res, err = msgServer.CreateTransferOperation(sdk.WrapSDKContext(ctx), msg)
		case *types.MsgCreateIdentityDefaultTransferOp:
			res, err = msgServer.CreateIdentityDefaultTransferOperation(sdk.WrapSDKContext(ctx), msg)
		case *types.MsgCreateIdentityGISTTransferOp:
			res, err = msgServer.CreateIdentityGISTTransferOperation(sdk.WrapSDKContext(ctx), msg)
		case *types.MsgCreateIdentityStateTransferOp:
			res, err = msgServer.CreateIdentityStateTransferOperation(sdk.WrapSDKContext(ctx), msg)
		case *types.MsgCreateWorldCoinIdentityTransferOp:
			res, err = msgServer.CreateWorldCoinIdentityTransferOperation(sdk.WrapSDKContext(ctx), msg)
		case *types.MsgVote:
			res, err = msgServer.Vote(sdk.WrapSDKContext(ctx), msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, errors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}

		return sdk.WrapServiceResult(ctx, res, err)
	}
}
