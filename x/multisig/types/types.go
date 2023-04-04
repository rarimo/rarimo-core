package types

const (
	EventTypeCreateGroup    = "create_group"
	EventTypeChangeGroup    = "change_group"
	EventTypeVote           = "vote"
	EventTypeSubmitProposal = "submit_proposal"

	AttributeKeyGroup      = "group"
	AttributeKeyProposal   = "proposal"
	AttributeKeyVoteOption = "vote_option"

	GroupAccountKey byte = 0x10
)
