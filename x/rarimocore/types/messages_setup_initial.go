package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	TypeMsgSetupInitial = "setup_initial"
)

var _ sdk.Msg = &MsgSetupInitial{}

func NewMsgSetupInitial(
	creator string,
	set []*Party,
	signature string,
) *MsgSetupInitial {
	return &MsgSetupInitial{
		Creator:   creator,
		Set:       set,
		Signature: signature,
	}
}

func (msg *MsgSetupInitial) Route() string {
	return RouterKey
}

func (msg *MsgSetupInitial) Type() string {
	return TypeMsgSetupInitial
}

func (msg *MsgSetupInitial) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetupInitial) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetupInitial) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.Set) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid set size (0)")
	}

	if _, err = hexutil.Decode(msg.Signature); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid hex receiver address (%s)", err)
	}

	return nil
}
