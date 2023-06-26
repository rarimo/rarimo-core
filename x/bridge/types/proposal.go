package types

import (
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	ProposalTypeChangeParams = "bridge/ChangeParams"
)

func init() {
	govv1beta1.RegisterProposalType(ProposalTypeChangeParams)
	govv1beta1.ModuleCdc.RegisterConcrete(&ChangeParamsProposal{}, "bridge/ChangeParamsProposal", nil)
}

// Implements Proposal Interface
var _ govv1beta1.Content = &ChangeParamsProposal{}

func (m *ChangeParamsProposal) ProposalRoute() string { return RouterKey }
func (m *ChangeParamsProposal) ProposalType() string  { return ProposalTypeChangeParams }

func (m *ChangeParamsProposal) ValidateBasic() error {
	return govv1beta1.ValidateAbstract(m)
}
