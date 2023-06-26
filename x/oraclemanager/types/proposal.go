package types

import (
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	ProposalTypeOracleUnfreeze = "oraclemanager/OracleUnfreeze"
	ProposalTypeChangeParams   = "oraclemanager/ChangeParams"
)

func init() {
	govv1beta1.RegisterProposalType(ProposalTypeOracleUnfreeze)
	govv1beta1.RegisterProposalType(ProposalTypeChangeParams)
	govv1beta1.ModuleCdc.RegisterConcrete(&OracleUnfreezeProposal{}, "oraclemanager/OracleUnfreezeProposal", nil)
	govv1beta1.ModuleCdc.RegisterConcrete(&ChangeParamsProposal{}, "oraclemanager/ChangeParamsProposal", nil)
}

// Implements Proposal Interface
var _ govv1beta1.Content = &OracleUnfreezeProposal{}

func (m *OracleUnfreezeProposal) ProposalRoute() string { return RouterKey }
func (m *OracleUnfreezeProposal) ProposalType() string  { return ProposalTypeOracleUnfreeze }

func (m *OracleUnfreezeProposal) ValidateBasic() error {
	return govv1beta1.ValidateAbstract(m)
}

// Implements Proposal Interface
var _ govv1beta1.Content = &ChangeParamsProposal{}

func (m *ChangeParamsProposal) ProposalRoute() string { return RouterKey }
func (m *ChangeParamsProposal) ProposalType() string  { return ProposalTypeChangeParams }

func (m *ChangeParamsProposal) ValidateBasic() error {
	return govv1beta1.ValidateAbstract(m)
}
