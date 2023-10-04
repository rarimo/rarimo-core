package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k msgServer) CreateChangePartiesOperation(goCtx context.Context, msg *types.MsgCreateChangePartiesOp) (*types.MsgCreateChangePartiesOpResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Only active party can submit such operation
	if err := k.checkIsAnActiveParty(ctx, msg.Creator); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "only active parties can submit that transaction")
	}

	var changeOp = &types.ChangeParties{
		Parties:      msg.NewSet,
		Signature:    msg.Signature,
		NewPublicKey: msg.NewPublicKey,
	}

	if err := k.Keeper.CreateChangePartiesOperation(ctx, msg.Creator, changeOp); err != nil {
		return nil, err
	}

	return &types.MsgCreateChangePartiesOpResponse{}, nil
}
