package types

import gov "github.com/cosmos/cosmos-sdk/x/gov/types"

const (
	ProposalTypeSetTokenInfo    = "set_token_info"
	ProposalTypeSetTokenItem    = "set_token_item"
	ProposalTypeRemoveTokenItem = "remove_token_item"
	ProposalTypeRemoveTokenInfo = "remove_token_info"
	ProposalTypeSetNetwork      = "set_network"
)

func init() {
	gov.RegisterProposalType(ProposalTypeSetTokenInfo)
	gov.RegisterProposalType(ProposalTypeSetTokenItem)
	gov.RegisterProposalType(ProposalTypeRemoveTokenItem)
	gov.RegisterProposalType(ProposalTypeRemoveTokenInfo)
	gov.RegisterProposalType(ProposalTypeSetNetwork)
	gov.RegisterProposalTypeCodec(&SetTokenInfoProposal{}, "rarimocore/SetTokenInfoProposal")
	gov.RegisterProposalTypeCodec(&SetTokenItemProposal{}, "rarimocore/SetTokenItemProposal")
	gov.RegisterProposalTypeCodec(&RemoveTokenItemProposal{}, "rarimocore/RemoveTokenItemProposal")
	gov.RegisterProposalTypeCodec(&RemoveTokenInfoProposal{}, "rarimocore/RemoveTokenInfoProposal")
	gov.RegisterProposalTypeCodec(&SetNetworkProposal{}, "rarimocore/SetNetworkProposal")
}

// Implements Proposal Interface
var _ gov.Content = &SetTokenInfoProposal{}

func (m *SetTokenInfoProposal) ProposalRoute() string { return RouterKey }
func (m *SetTokenInfoProposal) ProposalType() string  { return ProposalTypeSetTokenInfo }

func (m *SetTokenInfoProposal) ValidateBasic() error {
	return gov.ValidateAbstract(m)
}

// Implements Proposal Interface
var _ gov.Content = &SetTokenItemProposal{}

func (m *SetTokenItemProposal) ProposalRoute() string { return RouterKey }
func (m *SetTokenItemProposal) ProposalType() string  { return ProposalTypeSetTokenItem }

func (m *SetTokenItemProposal) ValidateBasic() error {
	return gov.ValidateAbstract(m)
}

// Implements Proposal Interface
var _ gov.Content = &RemoveTokenItemProposal{}

func (m *RemoveTokenItemProposal) ProposalRoute() string { return RouterKey }
func (m *RemoveTokenItemProposal) ProposalType() string  { return ProposalTypeRemoveTokenItem }

func (m *RemoveTokenItemProposal) ValidateBasic() error {
	return gov.ValidateAbstract(m)
}

// Implements Proposal Interface
var _ gov.Content = &RemoveTokenInfoProposal{}

func (m *RemoveTokenInfoProposal) ProposalRoute() string { return RouterKey }
func (m *RemoveTokenInfoProposal) ProposalType() string  { return ProposalTypeRemoveTokenInfo }

func (m *RemoveTokenInfoProposal) ValidateBasic() error {
	return gov.ValidateAbstract(m)
}

// Implements Proposal Interface
var _ gov.Content = &SetNetworkProposal{}

func (m *SetNetworkProposal) ProposalRoute() string { return RouterKey }
func (m *SetNetworkProposal) ProposalType() string  { return ProposalTypeSetNetwork }

func (m *SetNetworkProposal) ValidateBasic() error {
	return gov.ValidateAbstract(m)
}
