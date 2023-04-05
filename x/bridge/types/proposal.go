package types

import (
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	ProposalTypeChangeParams = "bridge/ChangeParams"
)

func init() {
	gov.RegisterProposalType(ProposalTypeChangeParams)
	gov.RegisterProposalTypeCodec(&ChangeParamsProposal{}, "bridge/ChangeParamsProposal")
}

// Implements Proposal Interface
var _ gov.Content = &ChangeParamsProposal{}

func (m *ChangeParamsProposal) ProposalRoute() string { return RouterKey }
func (m *ChangeParamsProposal) ProposalType() string  { return ProposalTypeChangeParams }

func (m *ChangeParamsProposal) ValidateBasic() error {
	return gov.ValidateAbstract(m)
}
