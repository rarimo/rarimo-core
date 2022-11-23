package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	TypeMsgCreateConfirmation = "create_confirmation"
)

var _ sdk.Msg = &MsgCreateConfirmation{}

func NewMsgCreateConfirmation(
	creator string,
	root string,
	indexes []string,
	sigECDSA string,

) *MsgCreateConfirmation {
	return &MsgCreateConfirmation{
		Creator:        creator,
		Root:           root,
		Indexes:        indexes,
		SignatureECDSA: sigECDSA,
	}
}

func (msg *MsgCreateConfirmation) Route() string {
	return RouterKey
}

func (msg *MsgCreateConfirmation) Type() string {
	return TypeMsgCreateConfirmation
}

func (msg *MsgCreateConfirmation) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateConfirmation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateConfirmation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if _, err = hexutil.Decode(msg.SignatureECDSA); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid hex receiver address (%s)", err)
	}

	if _, err = hexutil.Decode(msg.Root); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid hex receiver address (%s)", err)
	}

	return nil
}
