package types

import (
	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	TypeMsgCreateAddSignerPartyProposal = "create_add_signer_party_proposal"
)

func NewMsgCreateAddSignerPartyProposal(title, description string, creator, account types.AccAddress, address, trialPublicKey string, deposit sdk.Coins) sdk.Msg {
	return &MsgCreateAddSignerPartyProposal{
		Creator:        creator.String(),
		Title:          title,
		Description:    description,
		Account:        account.String(),
		Address:        address,
		TrialPublicKey: trialPublicKey,
		Deposit:        deposit,
	}
}

var _ sdk.Msg = &MsgCreateConfirmation{}

func (msg *MsgCreateAddSignerPartyProposal) Route() string {
	return RouterKey
}

func (msg *MsgCreateAddSignerPartyProposal) Type() string {
	return TypeMsgCreateAddSignerPartyProposal
}

func (msg *MsgCreateAddSignerPartyProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAddSignerPartyProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAddSignerPartyProposal) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.Account); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid account address (%s)", err)
	}

	if _, err := hexutil.Decode(msg.TrialPublicKey); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid hex pub key address (%s)", err)
	}

	return nil
}
