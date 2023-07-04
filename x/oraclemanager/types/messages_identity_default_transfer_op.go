package types

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	TypeMsgCreateIdentityDefaultTransfer = "create_identity_default_transfer_operation"
)

var _ sdk.Msg = &MsgCreateIdentityDefaultTransferOp{}

func (msg *MsgCreateIdentityDefaultTransferOp) Route() string {
	return RouterKey
}

func (msg *MsgCreateIdentityDefaultTransferOp) Type() string {
	return TypeMsgCreateIdentityDefaultTransfer
}

func (msg *MsgCreateIdentityDefaultTransferOp) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateIdentityDefaultTransferOp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateIdentityDefaultTransferOp) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return msg.ValidateBody()
}

func (msg *MsgCreateIdentityDefaultTransferOp) ValidateBody() error {
	if _, err := hexutil.Decode(msg.Contract); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid contract address (%s)", err)
	}

	if _, err := hexutil.Decode(msg.GISTHash); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid gist hash (%s)", err)
	}

	if _, err := hexutil.Decode(msg.Id); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid id (%s)", err)
	}

	if _, err := hexutil.Decode(msg.StateHash); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid state hash (%s)", err)
	}

	if _, err := hexutil.Decode(msg.StateReplacedBy); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid state replaced by (%s)", err)
	}

	if _, err := hexutil.Decode(msg.GISTReplacedBy); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid gist replaced by (%s)", err)
	}

	if _, ok := new(big.Int).SetString(msg.StateCreatedAtTimestamp, 10); !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid timestamp")
	}

	if _, ok := new(big.Int).SetString(msg.StateCreatedAtBlock, 10); !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid timestamp")
	}

	if _, ok := new(big.Int).SetString(msg.StateReplacedAtTimestamp, 10); !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid timestamp")
	}

	if _, ok := new(big.Int).SetString(msg.StateReplacedAtBlock, 10); !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid timestamp")
	}

	if _, ok := new(big.Int).SetString(msg.GISTCreatedAtTimestamp, 10); !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid timestamp")
	}

	if _, ok := new(big.Int).SetString(msg.GISTCreatedAtBlock, 10); !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid timestamp")
	}

	if _, ok := new(big.Int).SetString(msg.GISTReplacedAtTimestamp, 10); !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid timestamp")
	}

	if _, ok := new(big.Int).SetString(msg.GISTReplacedAtBlock, 10); !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid timestamp")
	}

	return nil
}
