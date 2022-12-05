package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

func (k msgServer) CreateAddSignerPartyProposal(goCtx context.Context, msg *types.MsgCreateAddSignerPartyProposal) (*types.MsgCreateAddSignerPartyProposalResponse, error) {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	content := &types.AddSignerPartyProposal{
		Title:          msg.Title,
		Description:    msg.Description,
		Account:        msg.Account,
		Address:        msg.Address,
		TrialPublicKey: msg.TrialPublicKey,
	}

	proposal, err := govtypes.NewMsgSubmitProposal(content, msg.Deposit, creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "can not create proposal %v", err)
	}

	_, err = govkeeper.NewMsgServerImpl(*k.gov).SubmitProposal(goCtx, proposal)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateAddSignerPartyProposalResponse{}, nil
}
