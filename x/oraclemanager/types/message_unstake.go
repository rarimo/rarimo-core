package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

const (
	TypeMsgUnstake = "unstake"
)

var _ sdk.Msg = &MsgUnstake{}

func NewMsgUnstake(oracle, chain string) *MsgUnstake {
	return &MsgUnstake{
		Index: &OracleIndex{
			Chain:   chain,
			Account: oracle,
		},
	}
}

func (msg *MsgUnstake) Route() string {
	return types.RouterKey
}

func (msg *MsgUnstake) Type() string {
	return TypeMsgUnstake
}

func (msg *MsgUnstake) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Index.Account)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUnstake) GetSignBytes() []byte {
	bz := types.ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUnstake) ValidateBasic() error {
	if msg.Index == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid index: nil")
	}

	_, err := sdk.AccAddressFromBech32(msg.Index.Account)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
