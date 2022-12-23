package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	currentId string,
	currentAddress string,
	currentChain string,
	currentType Type,
) *MsgCreateInfo {
	return &MsgCreateInfo{
		Creator:        creator,
		Index:          index,
		CurrentChain:   currentChain,
		CurrentId:      currentId,
		CurrentAddress: currentAddress,
		CurrentType:    currentType,
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

	if msg.CurrentType != Type_NATIVE {
		if _, err = hexutil.Decode(msg.CurrentAddress); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid hex token address (%s)", err)
		}

		if _, err = hexutil.Decode(msg.CurrentId); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid hex token id (%s)", err)
		}
	}

	if msg.CurrentSeed != "" {
		if _, err = hexutil.Decode(msg.CurrentSeed); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid hex token seed (%s)", err)
		}
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
	tokenType Type,
) *MsgAddChain {
	return &MsgAddChain{
		Creator:      creator,
		Index:        index,
		ChainName:    chain,
		TokenAddress: tokenAddress,
		TokenId:      tokenId,
		TokenType:    tokenType,
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

	if _, err = hexutil.Decode(msg.TokenAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid hex token address (%s)", err)
	}

	if _, err = hexutil.Decode(msg.TokenId); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid hex token id (%s)", err)
	}

	return nil
}
