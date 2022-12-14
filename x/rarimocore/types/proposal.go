package types

import (
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	ProposalTypeAddSignerParty    = "add_signer_party"
	ProposalTypeRemoveSignerParty = "remove_signer_party"
)

func init() {
	gov.RegisterProposalType(ProposalTypeAddSignerParty)
	gov.RegisterProposalType(ProposalTypeRemoveSignerParty)
	gov.RegisterProposalTypeCodec(&AddSignerPartyProposal{}, "rarimocore/AddSignerPartyProposal")
	gov.RegisterProposalTypeCodec(&RemoveSignerPartyProposal{}, "rarimocore/RemoveSignerPartyProposal")
}

// Implements Proposal Interface
var _ gov.Content = &AddSignerPartyProposal{}

func (m *AddSignerPartyProposal) ProposalRoute() string { return RouterKey }
func (m *AddSignerPartyProposal) ProposalType() string  { return ProposalTypeAddSignerParty }

func (m *AddSignerPartyProposal) ValidateBasic() error {
	return gov.ValidateAbstract(m)
}

// Implements Proposal Interface
var _ gov.Content = &RemoveSignerPartyProposal{}

func (m *RemoveSignerPartyProposal) ProposalRoute() string { return RouterKey }
func (m *RemoveSignerPartyProposal) ProposalType() string  { return ProposalTypeRemoveSignerParty }

func (m *RemoveSignerPartyProposal) ValidateBasic() error {
	return gov.ValidateAbstract(m)
}
