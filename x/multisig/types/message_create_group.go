package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgCreateGroup = "create_group"
)

var _ sdk.Msg = &MsgCreateGroup{}

func NewMsgCreateGroup(
	creator string,
	members []string,
	threshold uint64,
) *MsgCreateGroup {
	return &MsgCreateGroup{
		Creator:   creator,
		Members:   members,
		Threshold: threshold,
	}
}

func (msg *MsgCreateGroup) Route() string {
	return RouterKey
}

func (msg *MsgCreateGroup) Type() string {
	return TypeMsgCreateGroup
}

func (msg *MsgCreateGroup) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateGroup) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateGroup) ValidateBasic() error {
	return validateGroupMessage(msg)
}
