package types

import (
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	ProposalTypeUnfreezeSignerParty = "rarimocore/UnfreezeSignerParty"
	ProposalTypeReshareKeys         = "rarimocore/ReshareKeys"
	ProposalTypeSlash               = "rarimocore/SlashProposal"
	ProposalTypeDropParties         = "rarimocore/DropPartiesProposal"
)

func init() {
	govv1beta1.RegisterProposalType(ProposalTypeUnfreezeSignerParty)
	govv1beta1.RegisterProposalType(ProposalTypeReshareKeys)
	govv1beta1.RegisterProposalType(ProposalTypeSlash)
	govv1beta1.RegisterProposalType(ProposalTypeDropParties)
	govv1beta1.ModuleCdc.RegisterConcrete(&UnfreezeSignerPartyProposal{}, "rarimocore/UnfreezeSignerPartyProposal", nil)
	govv1beta1.ModuleCdc.RegisterConcrete(&ReshareKeysProposal{}, "rarimocore/ReshareKeysProposal", nil)
	govv1beta1.ModuleCdc.RegisterConcrete(&SlashProposal{}, "rarimocore/SlashProposal", nil)
	govv1beta1.ModuleCdc.RegisterConcrete(&DropPartiesProposal{}, "rarimocore/DropPartiesProposal", nil)
}

// Implements Proposal Interface
var _ govv1beta1.Content = &UnfreezeSignerPartyProposal{}

func (m *UnfreezeSignerPartyProposal) ProposalRoute() string { return RouterKey }
func (m *UnfreezeSignerPartyProposal) ProposalType() string  { return ProposalTypeUnfreezeSignerParty }

func (m *UnfreezeSignerPartyProposal) ValidateBasic() error {
	return govv1beta1.ValidateAbstract(m)
}

// Implements Proposal Interface
var _ govv1beta1.Content = &ReshareKeysProposal{}

func (m *ReshareKeysProposal) ProposalRoute() string { return RouterKey }
func (m *ReshareKeysProposal) ProposalType() string  { return ProposalTypeReshareKeys }

func (m *ReshareKeysProposal) ValidateBasic() error {
	return govv1beta1.ValidateAbstract(m)
}

// Implements Proposal Interface
var _ govv1beta1.Content = &SlashProposal{}

func (m *SlashProposal) ProposalRoute() string { return RouterKey }
func (m *SlashProposal) ProposalType() string  { return ProposalTypeSlash }

func (m *SlashProposal) ValidateBasic() error {
	return govv1beta1.ValidateAbstract(m)
}

// Implements Proposal Interface
var _ govv1beta1.Content = &DropPartiesProposal{}

func (m *DropPartiesProposal) ProposalRoute() string { return RouterKey }
func (m *DropPartiesProposal) ProposalType() string  { return ProposalTypeDropParties }

func (m *DropPartiesProposal) ValidateBasic() error {
	return govv1beta1.ValidateAbstract(m)
}
