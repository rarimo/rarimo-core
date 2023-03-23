package types

import (
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	ProposalTypeOracleUnfreeze = "OracleUnfreeze"
	ProposalTypeChangeParams   = "ChangeParams"
)

func init() {
	gov.RegisterProposalType(ProposalTypeOracleUnfreeze)
	gov.RegisterProposalType(ProposalTypeChangeParams)
	gov.RegisterProposalTypeCodec(&OracleUnfreezeProposal{}, "oraclemanager/OracleUnfreezeProposal")
	gov.RegisterProposalTypeCodec(&ChangeParamsProposal{}, "oraclemanager/ChangeParamsProposal")
}

// Implements Proposal Interface
var _ gov.Content = &OracleUnfreezeProposal{}

func (m *OracleUnfreezeProposal) ProposalRoute() string { return RouterKey }
func (m *OracleUnfreezeProposal) ProposalType() string  { return ProposalTypeOracleUnfreeze }

func (m *OracleUnfreezeProposal) ValidateBasic() error {
	return gov.ValidateAbstract(m)
}

// Implements Proposal Interface
var _ gov.Content = &ChangeParamsProposal{}

func (m *ChangeParamsProposal) ProposalRoute() string { return RouterKey }
func (m *ChangeParamsProposal) ProposalType() string  { return ProposalTypeChangeParams }

func (m *ChangeParamsProposal) ValidateBasic() error {
	return gov.ValidateAbstract(m)
}
