package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k msgServer) Vote(goCtx context.Context, msg *types.MsgVote) (*types.MsgVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.checkCreatorIsValidator(ctx, msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "sender is not a validator")
	}

	if _, ok := k.GetVote(ctx, &types.VoteIndex{Operation: msg.Operation, Validator: msg.Creator}); ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "vote already exists")
	}

	k.SetVote(ctx, types.Vote{
		Index: &types.VoteIndex{
			Operation: msg.Operation,
			Validator: msg.Creator,
		},
		Vote: msg.Vote,
	})

	// TODO calculate vote power

	return &types.MsgVoteResponse{}, nil
}
