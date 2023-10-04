package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/rarimo/rarimo-core/x/rarimocore/types"
)

const (
	TypeMsgVote = "vote"
)

var _ sdk.Msg = &MsgVote{}

func NewMsgVote(oracle, chain, operation string, vote types.VoteType) *MsgVote {
	return &MsgVote{
		Index: &OracleIndex{
			Chain:   chain,
			Account: oracle,
		},
		Operation: operation,
		Vote:      vote,
	}
}

func (msg *MsgVote) Route() string {
	return RouterKey
}

func (msg *MsgVote) Type() string {
	return TypeMsgVote
}

func (msg *MsgVote) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Index.Account)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgVote) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgVote) ValidateBasic() error {
	if msg.Index == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid index: nil")
	}

	_, err := sdk.AccAddressFromBech32(msg.Index.Account)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.Operation) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid operation")
	}

	return nil
}
