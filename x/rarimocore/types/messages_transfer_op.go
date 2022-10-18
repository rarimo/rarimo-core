package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	tokentypes "gitlab.com/rarify-protocol/rarimo-core/x/tokenmanager/types"
)

const (
	TypeMsgCreateDeposit = "create_transfer_operation"
)

var _ sdk.Msg = &MsgCreateTransferOp{}

func NewMsgCreateTransferOp(
	creator string,
	tx string,
	eventId string,
	fromChain string,
	tokenType tokentypes.Type,
) *MsgCreateTransferOp {
	return &MsgCreateTransferOp{
		Creator:   creator,
		Tx:        tx,
		EventId:   eventId,
		FromChain: fromChain,
		TokenType: tokenType,
	}
}

func (msg *MsgCreateTransferOp) Route() string {
	return RouterKey
}

func (msg *MsgCreateTransferOp) Type() string {
	return TypeMsgCreateDeposit
}

func (msg *MsgCreateTransferOp) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTransferOp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTransferOp) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
