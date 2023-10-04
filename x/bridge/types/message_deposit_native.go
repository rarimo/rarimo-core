package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rarimo/rarimo-core/x/tokenmanager/types"
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
	to types.OnChainItemIndex,
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
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if _, err := hexutil.Decode(msg.Receiver); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid receiver address (%s)", err)
	}

	if _, err := hexutil.Decode(msg.BundleData); len(msg.BundleData) != 0 && err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid bundle data (%s)", err)
	}

	if _, err := hexutil.Decode(msg.BundleSalt); len(msg.BundleSalt) != 0 && err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid bundle salt (%s)", err)
	}

	if msg.To.Chain == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid to")
	}

	return nil
}
