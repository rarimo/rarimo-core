package types

const (
	EventTypeCreateGroup           = "create_group"
	EventTypeChangeGroup           = "change_group"
	EventTypeVote                  = "vote"
	EventTypeSubmitProposal        = "submit_proposal"
	EventTypeProposalStatusChanged = "proposal_status_changed"
	EventTypeProposalPruned        = "proposal_pruned"
	EventTypeProposalExecuted      = "proposal_executed"

	AttributeKeyGroup                 = "group"
	AttributeKeyProposal              = "proposal"
	AttributeKeyProposalStatus        = "proposal_status"
	AttributeKeyProposalExecutionLogs = "proposal_execution_logs"
	AttributeKeyVoteOption            = "vote_option"

	GroupAccountKey byte = 0x10
)
