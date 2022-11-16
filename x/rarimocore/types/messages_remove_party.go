package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateRemovePartyOp = "create_remove_party"
)

var _ sdk.Msg = &MsgCreateRemovePartyOp{}

func NewMsgCreateRemovePartyOp(
	creator string,
	index uint32,
) *MsgCreateRemovePartyOp {
	return &MsgCreateRemovePartyOp{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgCreateRemovePartyOp) Route() string {
	return RouterKey
}

func (msg *MsgCreateRemovePartyOp) Type() string {
	return TypeMsgCreateRemovePartyOp
}

func (msg *MsgCreateRemovePartyOp) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateRemovePartyOp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateRemovePartyOp) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
