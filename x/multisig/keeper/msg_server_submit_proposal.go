package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pkg/errors"
	"github.com/rarimo/rarimo-core/x/multisig/types"
)

func (k msgServer) SubmitProposal(goCtx context.Context, msg *types.MsgSubmitProposal) (*types.MsgSubmitProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	group, found := k.GetGroup(ctx, msg.Group)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "group (%s) not found", msg.Group)
	}

	msgs, err := msg.GetMsgs()
	if err != nil {
		return nil, sdkerrors.Wrap(err, "request msgs")
	}

	if !group.HasMember(msg.Creator) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "not a member of group (%s)", msg.Group)
	}

	// Check that if the messages require signers, they are all equal to the given account address of group.
	if err = ensureMsgAuthZ(msgs, group.Account); err != nil {
		return nil, err
	}

	params := k.GetParams(ctx)

	proposal := types.Proposal{
		Proposer:       msg.Creator,
		Id:             params.ProposalSequence + 1,
		Group:          group.Account,
		SubmitBlock:    uint64(ctx.BlockHeight()),
		Status:         types.ProposalStatus_SUBMITTED,
		VotingEndBlock: uint64(ctx.BlockHeight()) + params.VotingPeriod,
	}

	err = proposal.SetMsgs(msgs)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to set proposal msgs")
	}

	params.ProposalSequence++
	k.SetProposal(ctx, proposal)
	k.SetParams(ctx, params)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSubmitProposal,
			sdk.NewAttribute(types.AttributeKeyProposal, fmt.Sprintf("%d", proposal.Id)),
			sdk.NewAttribute(types.AttributeKeyGroup, proposal.Group),
		),
	)

	return &types.MsgSubmitProposalResponse{
		ProposalId: proposal.Id,
	}, nil
}

// ensureMsgAuthZ checks that if a message requires signers that all of them
// are equal to the given account address of group.
func ensureMsgAuthZ(msgs []sdk.Msg, group string) error {
	groupAcc, _ := sdk.AccAddressFromBech32(group)
	for i := range msgs {
		// In practice, GetSigners() should return a non-empty array without
		// duplicates, so the code below is equivalent to:
		// `msgs[i].GetSigners()[0] == groupAcc`
		// but we prefer to loop through all GetSigners just to be sure.
		for _, acct := range msgs[i].GetSigners() {
			if !groupAcc.Equals(acct) {
				return errors.Wrapf(sdkerrors.ErrUnauthorized, "msg does not have group authorization; expected %s, got %s", groupAcc.String(), acct.String())
			}
		}
	}
	return nil
}
