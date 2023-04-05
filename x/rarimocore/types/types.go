package types

const (
	EventTypeNewOperation      = "new_operation"      // emitted only from transactions
	EventTypeNewConfirmation   = "new_confirmation"   // emitted only from transactions
	EventTypeOperationSigned   = "operation_signed"   // emitted only from transactions
	EventTypeParamsUpdated     = "params_updated"     // emitted only from transactions
	EventTypeOperationApproved = "operation_approved" // emitted only from transactions
	EventTypeOperationRejected = "operation_rejected" // emitted only from transactions
	EventTypeVoted             = "voted"              // emitted only from transactions

	AttributeKeyParamsUpdateType = "params_update_type"
	AttributeKeyOperationId      = "operation_id"
	AttributeKeyOperationType    = "operation_type"
	AttributeKeyConfirmationId   = "confirmation_id"
	AttributeKeyVotingChoice     = "voting_choice"
)
