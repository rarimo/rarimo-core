package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

const (
	TypeMsgUnjail = "unjail"
)

var _ sdk.Msg = &MsgUnjail{}

func NewMsgUnjail(oracle, chain string) *MsgUnjail {
	return &MsgUnjail{
		Index: &OracleIndex{
			Chain:   chain,
			Account: oracle,
		},
	}
}

func (msg *MsgUnjail) Route() string {
	return types.RouterKey
}

func (msg *MsgUnjail) Type() string {
	return TypeMsgUnjail
}

func (msg *MsgUnjail) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Index.Account)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUnjail) GetSignBytes() []byte {
	bz := types.ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUnjail) ValidateBasic() error {
	if msg.Index == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid index: nil")
	}

	_, err := sdk.AccAddressFromBech32(msg.Index.Account)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
