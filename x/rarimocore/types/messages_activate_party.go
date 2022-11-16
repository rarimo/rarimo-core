package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto"
)

const (
	TypeMsgActivateParty = "create_change_key_ecdsa"
)

var _ sdk.Msg = &MsgActivateParty{}

func NewMsgActivateParty(
	creator string,
	key string,
) *MsgActivateParty {
	return &MsgActivateParty{
		Creator: creator,
		PubKey:  key,
	}
}

func (msg *MsgActivateParty) Route() string {
	return RouterKey
}

func (msg *MsgActivateParty) Type() string {
	return TypeMsgActivateParty
}

func (msg *MsgActivateParty) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgActivateParty) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgActivateParty) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if err := crypto.ValidateECDSAKey(msg.PubKey); err != nil {
		return err
	}

	return nil
}
