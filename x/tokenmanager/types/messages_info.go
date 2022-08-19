package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateInfo = "create_info"
	TypeMsgUpdateInfo = "update_info"
	TypeMsgDeleteInfo = "delete_info"
	TypeMsgAddChain   = "add_chain"
)

var _ sdk.Msg = &MsgCreateInfo{}

func NewMsgCreateInfo(
	creator string,
	index string,
	name string,
	symbol string,
	image string,
	description string,
	animationUrl string,
	externalUrl string,
	attributes string,
	currentChain string,

) *MsgCreateInfo {
	return &MsgCreateInfo{
		Creator:      creator,
		Index:        index,
		Name:         name,
		Symbol:       symbol,
		Image:        image,
		Description:  description,
		AnimationUrl: animationUrl,
		ExternalUrl:  externalUrl,
		Attributes:   attributes,
		CurrentChain: currentChain,
	}
}

func (msg *MsgCreateInfo) Route() string {
	return RouterKey
}

func (msg *MsgCreateInfo) Type() string {
	return TypeMsgCreateInfo
}

func (msg *MsgCreateInfo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateInfo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateInfo{}

func NewMsgUpdateInfo(
	creator string,
	index string,
	name string,
	symbol string,
	image string,
	description string,
	animationUrl string,
	externalUrl string,
	attributes string,
	currentChain string,

) *MsgUpdateInfo {
	return &MsgUpdateInfo{
		Creator:      creator,
		Index:        index,
		Name:         name,
		Symbol:       symbol,
		Image:        image,
		Description:  description,
		AnimationUrl: animationUrl,
		ExternalUrl:  externalUrl,
		Attributes:   attributes,
		CurrentChain: currentChain,
	}
}

func (msg *MsgUpdateInfo) Route() string {
	return RouterKey
}

func (msg *MsgUpdateInfo) Type() string {
	return TypeMsgUpdateInfo
}

func (msg *MsgUpdateInfo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateInfo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteInfo{}

func NewMsgDeleteInfo(
	creator string,
	index string,

) *MsgDeleteInfo {
	return &MsgDeleteInfo{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteInfo) Route() string {
	return RouterKey
}

func (msg *MsgDeleteInfo) Type() string {
	return TypeMsgDeleteInfo
}

func (msg *MsgDeleteInfo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteInfo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgAddChain{}

func NewMsgAddChain(
	creator string,
	index string,
	chain string,
	tokenAddress string,
	tokenId string,
) *MsgAddChain {
	return &MsgAddChain{
		Creator:      creator,
		Index:        index,
		Chain:        chain,
		TokenAddress: tokenAddress,
		TokenId:      tokenId,
	}
}

func (msg *MsgAddChain) Route() string {
	return RouterKey
}

func (msg *MsgAddChain) Type() string {
	return TypeMsgAddChain
}

func (msg *MsgAddChain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddChain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddChain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
