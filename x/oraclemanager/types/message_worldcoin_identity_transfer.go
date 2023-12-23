package types

import (
	"math/big"

	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const TypeMsgCreateWorldCoinIdentityTransferOp = "create_worldcoin_identity_transfer_operation"

var _ types.Msg = &MsgCreateWorldCoinIdentityTransferOp{}

func (m *MsgCreateWorldCoinIdentityTransferOp) Route() string {
	return RouterKey
}

func (m *MsgCreateWorldCoinIdentityTransferOp) Type() string {
	return TypeMsgCreateWorldCoinIdentityTransferOp
}

func (m *MsgCreateWorldCoinIdentityTransferOp) GetSigners() []types.AccAddress {
	creator, err := types.AccAddressFromBech32(m.Creator)
	if err != nil {
		panic(err)
	}
	return []types.AccAddress{creator}
}

func (m *MsgCreateWorldCoinIdentityTransferOp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return types.MustSortJSON(bz)
}

func (m *MsgCreateWorldCoinIdentityTransferOp) ValidateBasic() error {
	if _, err := types.AccAddressFromBech32(m.Creator); err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return m.ValidateBody()
}

func (m *MsgCreateWorldCoinIdentityTransferOp) ValidateBody() error {
	if _, err := hexutil.Decode(m.Contract); err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid contract address (%s)", err)
	}
	if m.Chain == "" {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "chain required")
	}
	if _, err := hexutil.Decode(m.PrevState); err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid previous state (%s)", err)
	}
	if _, err := hexutil.Decode(m.State); err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid state (%s)", err)
	}
	if _, ok := new(big.Int).SetString(m.Timestamp, 10); !ok {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid timestamp %s", m.Timestamp)
	}
	if m.BlockNumber == 0 {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "block number required")
	}

	return nil
}
