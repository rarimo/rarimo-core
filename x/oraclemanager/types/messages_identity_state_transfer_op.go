package types

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	TypeMsgCreateIdentityStateTransfer = "create_identity_state_transfer_operation"
)

var _ sdk.Msg = &MsgCreateIdentityStateTransferOp{}

func (msg *MsgCreateIdentityStateTransferOp) Route() string {
	return RouterKey
}

func (msg *MsgCreateIdentityStateTransferOp) Type() string {
	return TypeMsgCreateIdentityStateTransfer
}

func (msg *MsgCreateIdentityStateTransferOp) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateIdentityStateTransferOp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateIdentityStateTransferOp) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return msg.ValidateBody()
}

func (msg *MsgCreateIdentityStateTransferOp) ValidateBody() error {
	if _, err := hexutil.Decode(msg.Contract); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid contract address (%s)", err)
	}

	if _, err := hexutil.Decode(msg.Id); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid id (%s)", err)
	}

	if _, err := hexutil.Decode(msg.StateHash); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid state hash (%s)", err)
	}

	if _, ok := new(big.Int).SetString(msg.StateCreatedAtTimestamp, 10); !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid timestamp")
	}

	if _, ok := new(big.Int).SetString(msg.StateCreatedAtBlock, 10); !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid timestamp")
	}

	if _, err := hexutil.Decode(msg.ReplacedStateHash); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid  replaced state hash (%s)", err)
	}

	return nil
}
