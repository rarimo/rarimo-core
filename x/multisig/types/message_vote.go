package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgVote = "vote"
)

var _ sdk.Msg = &MsgVote{}

func NewMsgVote(
	creator string,
	proposalId uint64,
	option VoteOption,
) *MsgVote {
	return &MsgVote{
		Creator:    creator,
		ProposalId: proposalId,
		Option:     option,
	}
}

func (msg *MsgVote) Route() string {
	return RouterKey
}

func (msg *MsgVote) Type() string {
	return TypeMsgVote
}

func (msg *MsgVote) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgVote) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgVote) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.ProposalId == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "proposal id cannot be 0")
	}

	if _, ok := VoteOption_name[int32(msg.Option)]; !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidType, "vote option")
	}

	return nil
}
