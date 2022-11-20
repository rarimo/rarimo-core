package types

import (
	"crypto/elliptic"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	eth "github.com/ethereum/go-ethereum/crypto"
)

const (
	TypeMsgCreateConfirmation = "create_confirmation"
)

var _ sdk.Msg = &MsgCreateConfirmation{}

func NewMsgCreateConfirmation(
	creator string,
	root string,
	indexes []string,
	sigECDSA string,

) *MsgCreateConfirmation {
	return &MsgCreateConfirmation{
		Creator:        creator,
		Root:           root,
		Indexes:        indexes,
		SignatureECDSA: sigECDSA,
	}
}

func (msg *MsgCreateConfirmation) Route() string {
	return RouterKey
}

func (msg *MsgCreateConfirmation) Type() string {
	return TypeMsgCreateConfirmation
}

func (msg *MsgCreateConfirmation) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateConfirmation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateConfirmation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if _, err = hexutil.Decode(msg.SignatureECDSA); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid hex receiver address (%s)", err)
	}

	if _, err = hexutil.Decode(msg.Root); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid hex receiver address (%s)", err)
	}

	if msg.Meta.NewKeyECDSA != getECDSAKey(msg.Meta.PartyKey) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid keys")
	}

	return nil
}

func getECDSAKey(keys []string) string {
	if len(keys) == 0 {
		return ""
	}

	curve := eth.S256()
	x, y := big.NewInt(0), big.NewInt(0)

	for _, key := range keys {
		kx, ky := elliptic.Unmarshal(curve, hexutil.MustDecode(key))
		x, y = curve.Add(x, y, kx, ky)
	}

	return hexutil.Encode(elliptic.Marshal(curve, x, y))
}
