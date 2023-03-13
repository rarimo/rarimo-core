package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

const (
	TypeMsgWithdrawNative = "withdraw_native"
)

var _ sdk.Msg = &MsgWithdrawNative{}

func NewMsgWithdrawNative(
	creator string,
	origin string,
	path string,
	signatures []string,
	recoveryId uint32,
	receiver string,
	amount *sdk.Coin,
	to *types.OnChainItemIndex,
) *MsgWithdrawNative {
	return &MsgWithdrawNative{
		Creator:    creator,
		Origin:     origin,
		Path:       path,
		Signatures: signatures,
		RecoveryId: recoveryId,
		Receiver:   receiver,
		Amount:     amount,
		To:         to,
	}
}

func (msg *MsgWithdrawNative) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawNative) Type() string {
	return TypeMsgWithdrawNative
}

func (msg *MsgWithdrawNative) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWithdrawNative) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawNative) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
