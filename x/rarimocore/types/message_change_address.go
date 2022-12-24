package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgChangePartyAddress = "change_party_address"
)

var _ sdk.Msg = &MsgChangePartyAddress{}

func NewMsgChangePartyAddress(
	creator string,
	address string,
) *MsgChangePartyAddress {
	return &MsgChangePartyAddress{
		Creator:    creator,
		NewAddress: address,
	}
}

func (msg *MsgChangePartyAddress) Route() string {
	return RouterKey
}

func (msg *MsgChangePartyAddress) Type() string {
	return TypeMsgChangePartyAddress
}

func (msg *MsgChangePartyAddress) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgChangePartyAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgChangePartyAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
