package types

import (
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	ProposalTypeAddSignerParty = "add_signer_party"
)

// Implements Proposal Interface
var _ gov.Content = &AddSignerPartyProposal{}

func init() {
	gov.RegisterProposalType(ProposalTypeAddSignerParty)
	gov.RegisterProposalTypeCodec(&AddSignerPartyProposal{}, "rarimocore/AddSignerPartyProposal")
}

func (m *AddSignerPartyProposal) ProposalRoute() string { return RouterKey }
func (m *AddSignerPartyProposal) ProposalType() string  { return ProposalTypeAddSignerParty }

func (m *AddSignerPartyProposal) ValidateBasic() error {
	return gov.ValidateAbstract(m)
}
