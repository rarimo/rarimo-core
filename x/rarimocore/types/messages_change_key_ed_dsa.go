package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateChangeKeyEdDSA = "create_change_key_ed_dsa"
	TypeMsgUpdateChangeKeyEdDSA = "update_change_key_ed_dsa"
	TypeMsgDeleteChangeKeyEdDSA = "delete_change_key_ed_dsa"
)

var _ sdk.Msg = &MsgCreateChangeKeyEdDSA{}

func NewMsgCreateChangeKeyEdDSA(
	creator string,
	newKey string,
	signature string,

) *MsgCreateChangeKeyEdDSA {
	return &MsgCreateChangeKeyEdDSA{
		Creator:   creator,
		NewKey:    newKey,
		Signature: signature,
	}
}

func (msg *MsgCreateChangeKeyEdDSA) Route() string {
	return RouterKey
}

func (msg *MsgCreateChangeKeyEdDSA) Type() string {
	return TypeMsgCreateChangeKeyEdDSA
}

func (msg *MsgCreateChangeKeyEdDSA) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateChangeKeyEdDSA) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateChangeKeyEdDSA) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateChangeKeyEdDSA{}

func NewMsgUpdateChangeKeyEdDSA(
	creator string,
	newKey string,
	signature string,

) *MsgUpdateChangeKeyEdDSA {
	return &MsgUpdateChangeKeyEdDSA{
		Creator:   creator,
		NewKey:    newKey,
		Signature: signature,
	}
}

func (msg *MsgUpdateChangeKeyEdDSA) Route() string {
	return RouterKey
}

func (msg *MsgUpdateChangeKeyEdDSA) Type() string {
	return TypeMsgUpdateChangeKeyEdDSA
}

func (msg *MsgUpdateChangeKeyEdDSA) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateChangeKeyEdDSA) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateChangeKeyEdDSA) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteChangeKeyEdDSA{}

func NewMsgDeleteChangeKeyEdDSA(
	creator string,
	newKey string,

) *MsgDeleteChangeKeyEdDSA {
	return &MsgDeleteChangeKeyEdDSA{
		Creator: creator,
		NewKey:  newKey,
	}
}
func (msg *MsgDeleteChangeKeyEdDSA) Route() string {
	return RouterKey
}

func (msg *MsgDeleteChangeKeyEdDSA) Type() string {
	return TypeMsgDeleteChangeKeyEdDSA
}

func (msg *MsgDeleteChangeKeyEdDSA) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteChangeKeyEdDSA) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteChangeKeyEdDSA) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
