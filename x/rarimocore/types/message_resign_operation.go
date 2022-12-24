package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgResignOperation = "resign_operation"
)

var _ sdk.Msg = &MsgResignOperation{}

func NewMsgResignOperation(
	creator string,
	change string,
	op string,
) *MsgResignOperation {
	return &MsgResignOperation{
		Creator:       creator,
		ChangeOpIndex: change,
		OpIndex:       op,
	}
}

func (msg *MsgResignOperation) Route() string {
	return RouterKey
}

func (msg *MsgResignOperation) Type() string {
	return TypeMsgResignOperation
}

func (msg *MsgResignOperation) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgResignOperation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgResignOperation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
