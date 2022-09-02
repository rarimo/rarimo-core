package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateConfirmation = "create_confirmation"
)

var _ sdk.Msg = &MsgCreateConfirmation{}

func NewMsgCreateConfirmation(
	creator string,
	height string,
	root string,
	hashes []string,
	sigECDSA string,

) *MsgCreateConfirmation {
	return &MsgCreateConfirmation{
		Creator:  creator,
		Height:   height,
		Root:     root,
		Hashes:   hashes,
		SigECDSA: sigECDSA,
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
