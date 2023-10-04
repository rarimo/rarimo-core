package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	tokentypes "github.com/rarimo/rarimo-core/x/tokenmanager/types"
)

const (
	TypeMsgCreateTransfer = "create_transfer_operation"
)

var _ sdk.Msg = &MsgCreateTransferOp{}

func NewMsgCreateTransferOp(
	creator string,
	tx string,
	eventId string,
	receiver string,
	amount string,
	bundleData string,
	bundleSalt string,
	from tokentypes.OnChainItemIndex,
	to tokentypes.OnChainItemIndex,
	meta *tokentypes.ItemMetadata,
) *MsgCreateTransferOp {
	return &MsgCreateTransferOp{
		Creator:    creator,
		Tx:         tx,
		EventId:    eventId,
		Receiver:   receiver,
		Amount:     amount,
		BundleData: bundleData,
		BundleSalt: bundleSalt,
		From:       from,
		To:         to,
		Meta:       meta,
	}
}

func (msg *MsgCreateTransferOp) Route() string {
	return RouterKey
}

func (msg *MsgCreateTransferOp) Type() string {
	return TypeMsgCreateTransfer
}

func (msg *MsgCreateTransferOp) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTransferOp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTransferOp) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return msg.ValidateBody()
}

func (msg *MsgCreateTransferOp) ValidateBody() error {
	if _, err := hexutil.Decode(msg.Receiver); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid receiver address (%s)", err)
	}

	if _, err := hexutil.Decode(msg.BundleData); len(msg.BundleData) != 0 && err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid bundle data (%s)", err)
	}

	if _, err := hexutil.Decode(msg.BundleSalt); len(msg.BundleSalt) != 0 && err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid bundle salt (%s)", err)
	}

	if msg.From.Chain == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid from")
	}

	if msg.To.Chain == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid to")
	}

	return nil
}
