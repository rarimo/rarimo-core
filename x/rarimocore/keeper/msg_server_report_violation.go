package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (t msgServer) ReportViolation(goCtx context.Context, msg *types.MsgCreateViolationReport) (*types.MsgCreateViolationReportResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Only Active party can submit report
	if err := t.checkIsAnActiveParty(ctx, msg.Creator); err != nil {
		return nil, sdkerrors.Wrap(err, "only active party can report the violation")
	}

	// Only Active party can be reported
	if err := t.checkIsAnActiveParty(ctx, msg.Offender); err != nil {
		return nil, sdkerrors.Wrap(err, "only active party can be reported")
	}

	index := types.NewViolationReportIndex(msg.SessionId, msg.Offender, msg.Creator, msg.ViolationType)

	// Report can be submitted only once for same `sessionId|offender|creator|violationType`
	if _, found := t.GetViolationReport(ctx, index); found {
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

	// Iterating over existing reports for our sessionId|offender to calculate violations count.

	params := t.GetParams(ctx)
	reports := make(map[string]struct{})

	t.IterateViolationReports(ctx, msg.SessionId, msg.Offender, func(report types.ViolationReport) (stop bool) {
		reports[report.Index.Sender] = struct{}{}
		return false
	})

	party := getPartyByAccount(msg.Offender, params.Parties)

	// Check if party violations count was already incremented for that session.
	alreadyIncremented := false
	for _, sessionReported := range party.ReportedSessions {
		alreadyIncremented = alreadyIncremented || (sessionReported == msg.SessionId)
	}

	// Trying to apply new violation if was not applied yet.
	if !alreadyIncremented {
		// Check that already have enough reports to confirm violation
		if uint64(len(reports)) >= params.Threshold {
			confirmViolation(party, msg.SessionId)
		}

		// If violations reaches threshold then party becomes frozen.
		if party.ViolationsCount == params.MaxViolationsCount {
			party.Status = types.PartyStatus_Frozen
			party.FreezeEndBlock = uint64(ctx.BlockHeight()) + params.FreezeBlocksPeriod

			ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypePartyFrozen,
				sdk.NewAttribute(types.AttributeKeyPartyAccount, party.Account),
			))

			params.IsUpdateRequired = true
		}

		t.SetParams(ctx, params)
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeNewViolationReport,
		sdk.NewAttribute(types.AttributeKeySessionId, msg.SessionId),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Offender),
		sdk.NewAttribute(types.AttributeKeyPartyAccount, msg.Offender),
		sdk.NewAttribute(types.AttributeKeyViolationType, msg.ViolationType.String()),
	))

	return &types.MsgCreateViolationReportResponse{}, nil
}

// confirmViolation updates party entry with:
// 1. increments violations counter
// 2. adds new sessionId to the reported sessions
// 3. removes old reports with corresponding decreasing of counter
func confirmViolation(party *types.Party, sessionId string) {
	party.ReportedSessions = append(party.ReportedSessions, sessionId)

	// we are assuming that session ids is an integer values encoded into string
	sessionIdInt, _ := strconv.Atoi(sessionId)

	// dropReportSessionDelta is an amount of block when report is considered to old and can be removed
	const dropReportSessionDelta = 1000

	actualReportedSessions := make([]string, 0, len(party.ReportedSessions))
	actualCounter := 0

	for _, ssi := range party.ReportedSessions {
		ssiInt, _ := strconv.Atoi(ssi)
		if ssiInt+dropReportSessionDelta > sessionIdInt {
			actualReportedSessions = append(actualReportedSessions, ssi)
			actualCounter++
		}
	}

	party.ReportedSessions = actualReportedSessions
	party.ViolationsCount = uint64(actualCounter)
}
