package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateViolationReport = "create_violation_report"
)

var _ sdk.Msg = &MsgCreateViolationReport{}

func NewMsgCreateViolationReport(
	creator,
	sessionId,
	offender string,
	violationType ViolationType,
) *MsgCreateViolationReport {
	return &MsgCreateViolationReport{
		Creator:       creator,
		SessionId:     sessionId,
		ViolationType: violationType,
		Offender:      offender,
	}
}

func (msg *MsgCreateViolationReport) Route() string {
	return RouterKey
}

func (msg *MsgCreateViolationReport) Type() string {
	return TypeMsgCreateViolationReport
}

func (msg *MsgCreateViolationReport) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateViolationReport) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateViolationReport) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Offender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid offender address (%s)", err)
	}

	return nil
}
