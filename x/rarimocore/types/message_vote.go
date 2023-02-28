package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgVote = "vote"
)

var _ sdk.Msg = &MsgVote{}

func NewMsgVote(creator, operation string, vote VoteType) *MsgVote {
	return &MsgVote{
		Creator:   creator,
		Operation: operation,
		Vote:      vote,
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
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.Operation) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid operation")
	}

	return nil
}
