package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateChangePartiesOp = "create_remove_party"
)

var _ sdk.Msg = &MsgCreateChangePartiesOp{}

func NewMsgCreateChangePartiesOp(
	creator string,
	set []*Party,
) *MsgCreateChangePartiesOp {
	return &MsgCreateChangePartiesOp{
		Creator: creator,
		NewSet:  set,
	}
}

func (msg *MsgCreateChangePartiesOp) Route() string {
	return RouterKey
}

func (msg *MsgCreateChangePartiesOp) Type() string {
	return TypeMsgCreateChangePartiesOp
}

func (msg *MsgCreateChangePartiesOp) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateChangePartiesOp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateChangePartiesOp) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.NewSet) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid set size (0)")
	}

	return nil
}
