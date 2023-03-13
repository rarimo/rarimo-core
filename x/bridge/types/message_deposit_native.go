package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

const (
	TypeMsgDepositNative = "deposit_native"
)

var _ sdk.Msg = &MsgDepositNative{}

func NewMsgDepositNative(
	creator string,
	seed string,
	receiver string,
	amount *sdk.Coin,
	bundleData string,
	bundleSalt string,
	to *types.OnChainItemIndex,
) *MsgDepositNative {
	return &MsgDepositNative{
		Creator:    creator,
		Seed:       seed,
		Receiver:   receiver,
		Amount:     amount,
		BundleData: bundleData,
		BundleSalt: bundleSalt,
		To:         to,
	}
}

func (msg *MsgDepositNative) Route() string {
	return RouterKey
}

func (msg *MsgDepositNative) Type() string {
	return TypeMsgDepositNative
}

func (msg *MsgDepositNative) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDepositNative) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDepositNative) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
