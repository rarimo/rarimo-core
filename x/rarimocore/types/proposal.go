package types

import (
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	ProposalTypeAddSignerParty    = "add_signer_party"
	ProposalTypeRemoveSignerParty = "remove_signer_party"
	ProposalTypeReshareKeys       = "reshare_keys"
	ProposalTypeChangeThreshold   = "change_threshold"
)

func init() {
	gov.RegisterProposalType(ProposalTypeAddSignerParty)
	gov.RegisterProposalType(ProposalTypeRemoveSignerParty)
	gov.RegisterProposalType(ProposalTypeReshareKeys)
	gov.RegisterProposalType(ProposalTypeChangeThreshold)
	gov.RegisterProposalTypeCodec(&AddSignerPartyProposal{}, "rarimocore/AddSignerPartyProposal")
	gov.RegisterProposalTypeCodec(&RemoveSignerPartyProposal{}, "rarimocore/RemoveSignerPartyProposal")
	gov.RegisterProposalTypeCodec(&ReshareKeysProposal{}, "rarimocore/ReshareKeysProposal")
	gov.RegisterProposalTypeCodec(&ChangeThresholdProposal{}, "rarimocore/ChangeThresholdProposal")
}

// Implements Proposal Interface
var _ gov.Content = &AddSignerPartyProposal{}

func (m *AddSignerPartyProposal) ProposalRoute() string { return RouterKey }
func (m *AddSignerPartyProposal) ProposalType() string  { return ProposalTypeAddSignerParty }

func (m *AddSignerPartyProposal) ValidateBasic() error {
	return gov.ValidateAbstract(m)
}

// Implements Proposal Interface
var _ gov.Content = &RemoveSignerPartyProposal{}

func (m *RemoveSignerPartyProposal) ProposalRoute() string { return RouterKey }
func (m *RemoveSignerPartyProposal) ProposalType() string  { return ProposalTypeRemoveSignerParty }

func (m *RemoveSignerPartyProposal) ValidateBasic() error {
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
var _ gov.Content = &ChangeThresholdProposal{}

func (m *ChangeThresholdProposal) ProposalRoute() string { return RouterKey }
func (m *ChangeThresholdProposal) ProposalType() string  { return ProposalTypeChangeThreshold }

func (m *ChangeThresholdProposal) ValidateBasic() error {
	return gov.ValidateAbstract(m)
}
