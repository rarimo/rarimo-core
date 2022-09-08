package types

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	tokentypes "gitlab.com/rarify-protocol/rarimo-core/x/tokenmanager/types"
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
	tokenType tokentypes.Type,
) *MsgCreateDeposit {
	return &MsgCreateDeposit{
		Creator:   creator,
		Tx:        tx,
		EventId:   eventId,
		FromChain: fromChain,
		ToChain:   toChain,
		Receiver:  receiver,
		TokenType: tokenType,
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

	if _, err = hexutil.Decode(msg.Receiver); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid hex receiver address (%s)", err)
	}

	if _, ok := new(big.Int).SetString(msg.Amount, 10); !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid dec amount")
	}

	if msg.Index != hexutil.Encode(crypto.Keccak256([]byte(msg.Tx), []byte(msg.EventId), []byte(msg.FromChain))) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid index: not a hash of tx|event|chain")
	}

	return nil
}
