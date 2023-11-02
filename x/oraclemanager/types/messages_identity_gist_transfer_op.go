package types

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	TypeMsgCreateIdentityGISTTransfer = "create_identity_gist_transfer_operation"
)

var _ sdk.Msg = &MsgCreateIdentityGISTTransferOp{}

func (msg *MsgCreateIdentityGISTTransferOp) Route() string {
	return RouterKey
}

func (msg *MsgCreateIdentityGISTTransferOp) Type() string {
	return TypeMsgCreateIdentityGISTTransfer
}

func (msg *MsgCreateIdentityGISTTransferOp) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateIdentityGISTTransferOp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateIdentityGISTTransferOp) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return msg.ValidateBody()
}

func (msg *MsgCreateIdentityGISTTransferOp) ValidateBody() error {
	if _, err := hexutil.Decode(msg.Contract); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid contract address (%s)", err)
	}

	if _, err := hexutil.Decode(msg.GISTHash); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid gist hash (%s)", err)
	}

	if _, err := hexutil.Decode(msg.GISTReplacedBy); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid gist replaced by (%s)", err)
	}

	if _, ok := new(big.Int).SetString(msg.GISTCreatedAtTimestamp, 10); !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid timestamp")
	}

	if _, ok := new(big.Int).SetString(msg.GISTCreatedAtBlock, 10); !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid timestamp")
	}

	if _, err := hexutil.Decode(msg.ReplacedGISTtHash); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid  replaced gist hash (%s)", err)
	}

	return nil
}
