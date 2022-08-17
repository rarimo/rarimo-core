package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateConfirmation = "create_confirmation"
	TypeMsgUpdateConfirmation = "update_confirmation"
	TypeMsgDeleteConfirmation = "delete_confirmation"
)

var _ sdk.Msg = &MsgCreateConfirmation{}

func NewMsgCreateConfirmation(
	creator string,
	height string,
	root string,
	hashes []string,
	sigECDSA string,
	sigEdDSA string,

) *MsgCreateConfirmation {
	return &MsgCreateConfirmation{
		Creator:  creator,
		Height:   height,
		Root:     root,
		Hashes:   hashes,
		SigECDSA: sigECDSA,
		SigEdDSA: sigEdDSA,
	}
}

func (msg *MsgCreateConfirmation) Route() string {
	return RouterKey
}

func (msg *MsgCreateConfirmation) Type() string {
	return TypeMsgCreateConfirmation
}

func (msg *MsgCreateConfirmation) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateConfirmation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateConfirmation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateConfirmation{}

func NewMsgUpdateConfirmation(
	creator string,
	height string,
	root string,
	hashes []string,
	sigECDSA string,
	sigEdDSA string,

) *MsgUpdateConfirmation {
	return &MsgUpdateConfirmation{
		Creator:  creator,
		Height:   height,
		Root:     root,
		Hashes:   hashes,
		SigECDSA: sigECDSA,
		SigEdDSA: sigEdDSA,
	}
}

func (msg *MsgUpdateConfirmation) Route() string {
	return RouterKey
}

func (msg *MsgUpdateConfirmation) Type() string {
	return TypeMsgUpdateConfirmation
}

func (msg *MsgUpdateConfirmation) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateConfirmation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateConfirmation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteConfirmation{}

func NewMsgDeleteConfirmation(
	creator string,
	height string,

) *MsgDeleteConfirmation {
	return &MsgDeleteConfirmation{
		Creator: creator,
		Height:  height,
	}
}
func (msg *MsgDeleteConfirmation) Route() string {
	return RouterKey
}

func (msg *MsgDeleteConfirmation) Type() string {
	return TypeMsgDeleteConfirmation
}

func (msg *MsgDeleteConfirmation) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteConfirmation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteConfirmation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
