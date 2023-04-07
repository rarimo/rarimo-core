package types

import (
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	ProposalTypeUnfreezeSignerParty = "rarimocore/UnfreezeSignerParty"
	ProposalTypeReshareKeys         = "rarimocore/ReshareKeys"
)

func init() {
	gov.RegisterProposalType(ProposalTypeUnfreezeSignerParty)
	gov.RegisterProposalType(ProposalTypeReshareKeys)
	gov.RegisterProposalTypeCodec(&UnfreezeSignerPartyProposal{}, "rarimocore/UnfreezeSignerPartyProposal")
	gov.RegisterProposalTypeCodec(&ReshareKeysProposal{}, "rarimocore/ReshareKeysProposal")
}

// Implements Proposal Interface
var _ gov.Content = &UnfreezeSignerPartyProposal{}

func (m *UnfreezeSignerPartyProposal) ProposalRoute() string { return RouterKey }
func (m *UnfreezeSignerPartyProposal) ProposalType() string  { return ProposalTypeUnfreezeSignerParty }

func (m *UnfreezeSignerPartyProposal) ValidateBasic() error {
	return gov.ValidateAbstract(m)
}

// Implements Proposal Interface
var _ gov.Content = &ReshareKeysProposal{}

func (m *ReshareKeysProposal) ProposalRoute() string { return RouterKey }
func (m *ReshareKeysProposal) ProposalType() string  { return ProposalTypeReshareKeys }

func (m *ReshareKeysProposal) ValidateBasic() error {
	return gov.ValidateAbstract(m)
}
