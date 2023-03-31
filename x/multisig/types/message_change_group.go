package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgChangeGroup = "change_group"
)

var _ sdk.Msg = &MsgChangeGroup{}

func NewMsgChangeGroup(
	creator string,
	group string,
	members []string,
	threshold uint64,
) *MsgChangeGroup {
	return &MsgChangeGroup{
		Creator:   creator,
		Group:     group,
		Members:   members,
		Threshold: threshold,
	}
}

func (msg *MsgChangeGroup) Route() string {
	return RouterKey
}

func (msg *MsgChangeGroup) Type() string {
	return TypeMsgChangeGroup
}

func (msg *MsgChangeGroup) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgChangeGroup) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgChangeGroup) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.Group); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid group address (%s)", err)
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
