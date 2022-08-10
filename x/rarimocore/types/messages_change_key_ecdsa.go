package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateChangeKeyECDSA = "create_change_key_ecdsa"
	TypeMsgUpdateChangeKeyECDSA = "update_change_key_ecdsa"
	TypeMsgDeleteChangeKeyECDSA = "delete_change_key_ecdsa"
)

var _ sdk.Msg = &MsgCreateChangeKeyECDSA{}

func NewMsgCreateChangeKeyECDSA(
	creator string,
	newKey string,
	signature string,

) *MsgCreateChangeKeyECDSA {
	return &MsgCreateChangeKeyECDSA{
		Creator:   creator,
		NewKey:    newKey,
		Signature: signature,
	}
}

func (msg *MsgCreateChangeKeyECDSA) Route() string {
	return RouterKey
}

func (msg *MsgCreateChangeKeyECDSA) Type() string {
	return TypeMsgCreateChangeKeyECDSA
}

func (msg *MsgCreateChangeKeyECDSA) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateChangeKeyECDSA) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateChangeKeyECDSA) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateChangeKeyECDSA{}

func NewMsgUpdateChangeKeyECDSA(
	creator string,
	newKey string,
	signature string,

) *MsgUpdateChangeKeyECDSA {
	return &MsgUpdateChangeKeyECDSA{
		Creator:   creator,
		NewKey:    newKey,
		Signature: signature,
	}
}

func (msg *MsgUpdateChangeKeyECDSA) Route() string {
	return RouterKey
}

func (msg *MsgUpdateChangeKeyECDSA) Type() string {
	return TypeMsgUpdateChangeKeyECDSA
}

func (msg *MsgUpdateChangeKeyECDSA) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateChangeKeyECDSA) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateChangeKeyECDSA) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteChangeKeyECDSA{}

func NewMsgDeleteChangeKeyECDSA(
	creator string,
	newKey string,

) *MsgDeleteChangeKeyECDSA {
	return &MsgDeleteChangeKeyECDSA{
		Creator: creator,
		NewKey:  newKey,
	}
}
func (msg *MsgDeleteChangeKeyECDSA) Route() string {
	return RouterKey
}

func (msg *MsgDeleteChangeKeyECDSA) Type() string {
	return TypeMsgDeleteChangeKeyECDSA
}

func (msg *MsgDeleteChangeKeyECDSA) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteChangeKeyECDSA) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteChangeKeyECDSA) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
