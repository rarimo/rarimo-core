package types

import (
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	ProposalTypeEditCSCAList = "cscalist/EditCSCAList"
)

func init() {
	govv1beta1.RegisterProposalType(ProposalTypeEditCSCAList)
	govv1beta1.ModuleCdc.RegisterConcrete(&EditCSCAListProposal{}, "rarimocore/EditCSCAListProposal", nil)
}

// Implements Proposal Interface
var _ govv1beta1.Content = &EditCSCAListProposal{}

func (m *EditCSCAListProposal) ProposalRoute() string { return RouterKey }
func (m *EditCSCAListProposal) ProposalType() string  { return ProposalTypeEditCSCAList }

func (m *EditCSCAListProposal) ValidateBasic() error {
	return govv1beta1.ValidateAbstract(m)
}
