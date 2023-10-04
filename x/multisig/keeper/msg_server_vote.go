package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/rarimo/rarimo-core/x/multisig/types"
)

func (k msgServer) Vote(goCtx context.Context, msg *types.MsgVote) (*types.MsgVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	currentBlockHeight := uint64(ctx.BlockHeight())

	proposal, found := k.GetProposal(ctx, msg.ProposalId)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "proposal (%d) not found", msg.ProposalId)
	}

	// Ensure that we can still accept votes for this proposal.
	if proposal.Status != types.ProposalStatus_SUBMITTED {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "proposal not open for voting")
	}

	// Ensure that the voter is a member of the group.
	group, found := k.GetGroup(ctx, proposal.Group)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "group (%s) not found", proposal.Group)
	}
	if !group.HasMember(msg.Creator) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "voter (%s) not a member of group (%s)", msg.Creator, proposal.Group)
	}

	vote := types.Vote{
		ProposalId:  proposal.Id,
		Voter:       msg.Creator,
		Option:      msg.Option,
		SubmitBlock: currentBlockHeight,
	}

	k.SetVote(ctx, vote)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeVote,
			sdk.NewAttribute(types.AttributeKeyProposal, fmt.Sprintf("%d", proposal.Id)),
			sdk.NewAttribute(types.AttributeKeyGroup, proposal.Group),
			sdk.NewAttribute(types.AttributeKeyVoteOption, vote.Option.String()),
		),
	)

	return &types.MsgVoteResponse{}, nil
}
