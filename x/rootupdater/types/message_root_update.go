package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgUpdateRoot = "update_root"
)

var _ sdk.Msg = &MsgUpdateRoot{}

func NewMsgUpdateRoot(
	creator string,
	address string,
) *MsgUpdateRoot {
	return &MsgUpdateRoot{
		Creator: creator,
		Root:    address,
	}
}

func (msg *MsgUpdateRoot) Route() string {
	return RouterKey
}

func (msg *MsgUpdateRoot) Type() string {
	return TypeMsgUpdateRoot
}

func (msg *MsgUpdateRoot) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateRoot) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateRoot) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Root == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid root")
	}
	return nil
}
