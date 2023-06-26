package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgWithdrawFee = "withdraw_fee"
)

var _ sdk.Msg = &MsgWithdrawFee{}

func NewMsgWithdrawFee(
	creator string,
	origin string,
) *MsgWithdrawFee {
	return &MsgWithdrawFee{
		Creator: creator,
		Origin:  origin,
	}
}

func (msg *MsgWithdrawFee) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawFee) Type() string {
	return TypeMsgWithdrawFee
}

func (msg *MsgWithdrawFee) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWithdrawFee) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawFee) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
