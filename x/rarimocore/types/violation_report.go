package types

func NewViolationReportIndex(sessionId, offender string, violationType ViolationType) *ViolationReportIndex {
	return &ViolationReportIndex{
		SessionId:     sessionId,
		Offender:      offender,
		ViolationType: violationType,
	}
}
