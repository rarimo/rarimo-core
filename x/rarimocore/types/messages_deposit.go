package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	TypeMsgCreateDeposit = "create_deposit"
)

var _ sdk.Msg = &MsgCreateDeposit{}

func NewMsgCreateDeposit(
	creator string,
	tx string,
	eventId string,
	fromChain string,
	toChain string,
	receiver string,
	tokenAddress string,
	tokenId string,
	tokenType Type,
) *MsgCreateDeposit {
	return &MsgCreateDeposit{
		Creator:      creator,
		Tx:           tx,
		EventId:      eventId,
		FromChain:    fromChain,
		ToChain:      toChain,
		Receiver:     receiver,
		TokenAddress: tokenAddress,
		TokenId:      tokenId,
		TokenType:    tokenType,
	}
}

func (msg *MsgCreateDeposit) Route() string {
	return RouterKey
}

func (msg *MsgCreateDeposit) Type() string {
	return TypeMsgCreateDeposit
}

func (msg *MsgCreateDeposit) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDeposit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDeposit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if _, err = hexutil.Decode(msg.TokenAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid hex token address (%s)", err)
	}

	if _, err = hexutil.Decode(msg.TokenId); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid hex token id (%s)", err)
	}

	if _, err = hexutil.Decode(msg.Receiver); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid hex receiver address (%s)", err)
	}

	return nil
}
