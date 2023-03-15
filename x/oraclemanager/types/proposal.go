package types

import (
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	ProposalTypeOracleUnfreeze = "OracleUnfreeze"
)

func init() {
	gov.RegisterProposalType(ProposalTypeOracleUnfreeze)
	gov.RegisterProposalTypeCodec(&OracleUnfreezeProposal{}, "rarimocore/OracleUnfreezeProposal")
}

// Implements Proposal Interface
var _ gov.Content = &OracleUnfreezeProposal{}

func (m *OracleUnfreezeProposal) ProposalRoute() string { return RouterKey }
func (m *OracleUnfreezeProposal) ProposalType() string  { return ProposalTypeOracleUnfreeze }

func (m *OracleUnfreezeProposal) ValidateBasic() error {
	return gov.ValidateAbstract(m)
}
