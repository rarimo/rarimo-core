package types

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgSubmitOperation = "create_group"
)

var _ sdk.Msg = &MsgSubmitOperation{}

func NewMsgSubmitOperation(
	creator string,
	group string,
	messages []*types.Any,
) *MsgSubmitOperation {
	return &MsgSubmitOperation{
		Creator:  creator,
		Group:    group,
		Messages: messages,
	}
}

func (msg *MsgSubmitOperation) Route() string {
	return RouterKey
}

func (msg *MsgSubmitOperation) Type() string {
	return TypeMsgSubmitOperation
}

func (msg *MsgSubmitOperation) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitOperation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitOperation) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.Group); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid group address (%s)", err)
	}

	if len(msg.Messages) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "no messages to submit")
	}

	return nil
}
