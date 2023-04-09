package types

import (
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	ProposalTypeUnfreezeSignerParty = "rarimocore/UnfreezeSignerParty"
	ProposalTypeReshareKeys         = "rarimocore/ReshareKeys"
	ProposalTypeSlash               = "rarimocore/SlashProposal"
	ProposalTypeDropParties         = "rarimocore/DropPartiesProposal"
)

func init() {
	gov.RegisterProposalType(ProposalTypeUnfreezeSignerParty)
	gov.RegisterProposalType(ProposalTypeReshareKeys)
	gov.RegisterProposalType(ProposalTypeSlash)
	gov.RegisterProposalType(ProposalTypeDropParties)
	gov.RegisterProposalTypeCodec(&UnfreezeSignerPartyProposal{}, "rarimocore/UnfreezeSignerPartyProposal")
	gov.RegisterProposalTypeCodec(&ReshareKeysProposal{}, "rarimocore/ReshareKeysProposal")
	gov.RegisterProposalTypeCodec(&SlashProposal{}, "rarimocore/SlashProposal")
	gov.RegisterProposalTypeCodec(&DropPartiesProposal{}, "rarimocore/DropPartiesProposal")
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

// Implements Proposal Interface
var _ gov.Content = &SlashProposal{}

func (m *SlashProposal) ProposalRoute() string { return RouterKey }
func (m *SlashProposal) ProposalType() string  { return ProposalTypeSlash }

func (m *SlashProposal) ValidateBasic() error {
	return gov.ValidateAbstract(m)
}

// Implements Proposal Interface
var _ gov.Content = &DropPartiesProposal{}

func (m *DropPartiesProposal) ProposalRoute() string { return RouterKey }
func (m *DropPartiesProposal) ProposalType() string  { return ProposalTypeDropParties }

func (m *DropPartiesProposal) ValidateBasic() error {
	return gov.ValidateAbstract(m)
}
