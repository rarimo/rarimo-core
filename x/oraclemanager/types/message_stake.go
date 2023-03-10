package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

const (
	TypeMsgStake = "stake"
)

var _ sdk.Msg = &MsgStake{}

func NewMsgStake(oracle, chain, amount string) *MsgStake {
	return &MsgStake{
		Index: &OracleIndex{
			Chain:   chain,
			Account: oracle,
		},
		Amount: amount,
	}
}

func (msg *MsgStake) Route() string {
	return types.RouterKey
}

func (msg *MsgStake) Type() string {
	return TypeMsgStake
}

func (msg *MsgStake) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Index.Account)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgStake) GetSignBytes() []byte {
	bz := types.ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgStake) ValidateBasic() error {
	if msg.Index == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid index: nil")
	}

	_, err := sdk.AccAddressFromBech32(msg.Index.Account)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if _, ok := sdk.NewIntFromString(msg.Amount); !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid amount")
	}

	return nil
}
