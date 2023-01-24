package types

const (
	EventTypeNewOperation      = "new_operation"
	EventTypeNewConfirmation   = "new_confirmation"
	EventTypeOperationSigned   = "operation_signed"
	EventTypeParamsUpdated     = "params_updated"
	EventTypeOperationApproved = "operation_approved"

	AttributeKeyParamsUpdateType = "params_update_type"
	AttributeKeyOperationId      = "operation_id"
	AttributeKeyOperationType    = "operation_type"
	AttributeKeyConfirmationId   = "confirmation_id"
)
