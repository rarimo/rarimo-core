package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	TypeMsgCreateChangeKeyECDSA = "create_change_key_ecdsa"
)

var _ sdk.Msg = &MsgCreateChangeKeyECDSA{}

func NewMsgCreateChangeKeyECDSA(
	creator string,
	newKey string,
) *MsgCreateChangeKeyECDSA {
	return &MsgCreateChangeKeyECDSA{
		Creator: creator,
		NewKey:  newKey,
	}
}

func (msg *MsgCreateChangeKeyECDSA) Route() string {
	return RouterKey
}

func (msg *MsgCreateChangeKeyECDSA) Type() string {
	return TypeMsgCreateChangeKeyECDSA
}

func (msg *MsgCreateChangeKeyECDSA) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateChangeKeyECDSA) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateChangeKeyECDSA) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if _, err = hexutil.Decode(msg.NewKey); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid hex receiver address (%s)", err)
	}

	return nil
}
