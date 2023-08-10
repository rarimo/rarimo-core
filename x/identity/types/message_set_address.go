package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	TypeMsgChangeGroup = "set_identity_contract_address"
)

var _ sdk.Msg = &MsgSetIdentityContractAddress{}

func NewMsgChangeGroup(
	creator string,
	address string,
) *MsgSetIdentityContractAddress {
	return &MsgSetIdentityContractAddress{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgSetIdentityContractAddress) Route() string {
	return RouterKey
}

func (msg *MsgSetIdentityContractAddress) Type() string {
	return TypeMsgChangeGroup
}

func (msg *MsgSetIdentityContractAddress) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetIdentityContractAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetIdentityContractAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if _, err = hexutil.Decode(msg.Address); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid hex address (%s)", err)
	}

	return nil
}
