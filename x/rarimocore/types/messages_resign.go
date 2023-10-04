package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgResign = "resign"
)

var _ sdk.Msg = &MsgResign{}

func NewMsgResign(creator, operation string) *MsgResign {
	return &MsgResign{
		Creator:   creator,
		Operation: operation,
	}
}

func (msg *MsgResign) Route() string {
	return RouterKey
}

func (msg *MsgResign) Type() string {
	return TypeMsgResign
}

func (msg *MsgResign) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgResign) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgResign) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
