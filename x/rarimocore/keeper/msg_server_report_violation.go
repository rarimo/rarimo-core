package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (t msgServer) ReportViolation(goCtx context.Context, msg *types.MsgCreateViolationReport) (*types.MsgCreateViolationReportResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := t.checkIsAnActiveParty(ctx, msg.Creator); err != nil {
		return nil, sdkerrors.Wrap(err, "only active party can report the violation")
	}

	if err := t.checkIsAnActiveParty(ctx, msg.Offender); err != nil {
		return nil, sdkerrors.Wrap(err, "only active party can be reported")
	}

	index := types.NewViolationReportIndex(msg.SessionId, msg.Offender, msg.Creator, msg.ViolationType)

	_, found := t.GetViolationReport(ctx, index)
	if found {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrConflict,
			"session_id: %s, offender: %s, sender: %s violation_type: %s",
			msg.SessionId,
			msg.Offender,
			msg.Creator,
			msg.ViolationType,
		)
	}

	t.SetViolationReport(ctx, types.ViolationReport{
		Index: index,
		Msg:   msg.Msg,
	})

	reports := 0
	params := t.GetParams(ctx)

	t.IterateViolationReports(ctx, msg.SessionId, msg.Offender, msg.ViolationType, func(report types.ViolationReport) (stop bool) {
		reports++

		if uint64(reports) == params.Threshold {
			for _, party := range params.Parties {
				if party.Account != msg.Offender {
					continue
				}

				if party.ViolationsCount+1 == params.MaxViolationsCount {
					t.SetPartyStatus(ctx, msg.Offender, types.PartyStatus_Frozen)
				}

				return t.IncrementPartyViolationsCounter(ctx, msg.Offender)
			}
		}
		return false
	})

	return &types.MsgCreateViolationReportResponse{}, nil
}
