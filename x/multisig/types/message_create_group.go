package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateGroup = "create_group"
)

var _ sdk.Msg = &MsgCreateGroup{}

func NewMsgCreateGroup(
	creator string,
	members []string,
	threshold uint64,
) *MsgCreateGroup {
	return &MsgCreateGroup{
		Creator:   creator,
		Members:   members,
		Threshold: threshold,
	}
}

func (msg *MsgCreateGroup) Route() string {
	return RouterKey
}

func (msg *MsgCreateGroup) Type() string {
	return TypeMsgCreateGroup
}

func (msg *MsgCreateGroup) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateGroup) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateGroup) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	for _, member := range msg.Members {
		if _, err := sdk.AccAddressFromBech32(member); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid member address (%s)", err)
		}
	}

	if msg.Threshold == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "threshold must be greater than 0")
	}

	return nil
}
