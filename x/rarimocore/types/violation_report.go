package types

func NewViolationReportIndex(sessionId, offender, sender string, violationType ViolationType) *ViolationReportIndex {
	return &ViolationReportIndex{
		SessionId:     sessionId,
		Offender:      offender,
		Sender:        sender,
		ViolationType: violationType,
	}
}
