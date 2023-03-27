package types

const (
	EventTypeNewOperation       = "new_operation"
	EventTypeNewConfirmation    = "new_confirmation"
	EventTypeOperationSigned    = "operation_signed"
	EventTypeParamsUpdated      = "params_updated"
	EventTypeOperationApproved  = "operation_approved"
	EventTypeOperationRejected  = "operation_rejected"
	EventTypeVoted              = "voted"
	EventTypeStaked             = "staked"
	EventTypeUnstaked           = "unstaked"
	EventTypePartyFrozen        = "party_frozen"
	EventTypeNewViolationReport = "new_violation_report"

	AttributeKeyParamsUpdateType = "params_update_type"
	AttributeKeyOperationId      = "operation_id"
	AttributeKeyOperationType    = "operation_type"
	AttributeKeyConfirmationId   = "confirmation_id"
	AttributeKeyVotingChoice     = "voting_choice"
	AttributeKeyPartyAccount     = "party_account"
	AttributeKeyViolationType    = "violation_type"
	AttributeKeySessionId        = "session_id"
)
