package types

import (
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	ProposalTypeSetNetwork           = "SetNetwork"
	ProposalTypeUpdateTokenItem      = "UpdateTokenItem"
	ProposalTypeRemoveTokenItem      = "RemoveTokenItem"
	ProposalTypeCreateCollection     = "CreateCollection"
	ProposalTypeUpdateCollectionData = "UpdateCollectionData"
	ProposalTypeAddCollectionData    = "AddCollectionData"
	ProposalTypeRemoveCollectionData = "RemoveCollectionData"
	ProposalTypeRemoveCollection     = "RemoveCollection"
)

func init() {
	gov.RegisterProposalType(ProposalTypeSetNetwork)
	gov.RegisterProposalType(ProposalTypeUpdateTokenItem)
	gov.RegisterProposalType(ProposalTypeRemoveTokenItem)
	gov.RegisterProposalType(ProposalTypeCreateCollection)
	gov.RegisterProposalType(ProposalTypeUpdateCollectionData)
	gov.RegisterProposalType(ProposalTypeAddCollectionData)
	gov.RegisterProposalType(ProposalTypeRemoveCollectionData)
	gov.RegisterProposalType(ProposalTypeRemoveCollection)

	gov.RegisterProposalTypeCodec(&SetNetworkProposal{}, "rarimocore/SetNetworkProposal")
	gov.RegisterProposalTypeCodec(&UpdateTokenItemProposal{}, "rarimocore/UpdateTokenItemProposal")
	gov.RegisterProposalTypeCodec(&RemoveTokenItemProposal{}, "rarimocore/RemoveTokenItemProposal")
	gov.RegisterProposalTypeCodec(&CreateCollectionProposal{}, "rarimocore/CreateCollectionProposal")
	gov.RegisterProposalTypeCodec(&UpdateCollectionDataProposal{}, "rarimocore/UpdateCollectionDataProposal")
	gov.RegisterProposalTypeCodec(&AddCollectionDataProposal{}, "rarimocore/AddCollectionDataProposal")
	gov.RegisterProposalTypeCodec(&RemoveCollectionDataProposal{}, "rarimocore/RemoveCollectionDataProposal")
	gov.RegisterProposalTypeCodec(&RemoveCollectionProposal{}, "rarimocore/RemoveCollectionProposal")
}

// Implements Proposal Interface
var _ gov.Content = &SetNetworkProposal{}

func (m *SetNetworkProposal) ProposalRoute() string { return RouterKey }
func (m *SetNetworkProposal) ProposalType() string  { return ProposalTypeSetNetwork }

func (m *SetNetworkProposal) ValidateBasic() error {
	if err := gov.ValidateAbstract(m); err != nil {
		return err
	}

	return validateNetwork(m.NetworkParams)
}

// Implements Proposal Interface
var _ gov.Content = &UpdateTokenItemProposal{}

func (m *UpdateTokenItemProposal) ProposalRoute() string { return RouterKey }
func (m *UpdateTokenItemProposal) ProposalType() string  { return ProposalTypeUpdateTokenItem }

func (m *UpdateTokenItemProposal) ValidateBasic() error {
	if err := gov.ValidateAbstract(m); err != nil {
		return err
	}

	for _, i := range m.Item {
		if err := validateItem(i); err != nil {
			return err
		}
	}

	return nil
}

// Implements Proposal Interface
var _ gov.Content = &RemoveTokenItemProposal{}

func (m *RemoveTokenItemProposal) ProposalRoute() string { return RouterKey }
func (m *RemoveTokenItemProposal) ProposalType() string  { return ProposalTypeRemoveTokenItem }

func (m *RemoveTokenItemProposal) ValidateBasic() error {
	if err := gov.ValidateAbstract(m); err != nil {
		return err
	}

	return nil
}

// Implements Proposal Interface
var _ gov.Content = &CreateCollectionProposal{}

func (m *CreateCollectionProposal) ProposalRoute() string { return RouterKey }
func (m *CreateCollectionProposal) ProposalType() string  { return ProposalTypeCreateCollection }

func (m *CreateCollectionProposal) ValidateBasic() error {
	if err := gov.ValidateAbstract(m); err != nil {
		return err
	}

	if err := validateCollection(&Collection{Index: m.Index, Meta: m.Metadata}); err != nil {
		return err
	}

	for _, d := range m.Data {
		if err := validateCollectionData(d); err != nil {
			return err
		}
	}

	for _, i := range m.Item {
		if err := validateItem(i); err != nil {
			return err
		}
	}

	return nil
}

// Implements Proposal Interface
var _ gov.Content = &UpdateCollectionDataProposal{}

func (m *UpdateCollectionDataProposal) ProposalRoute() string { return RouterKey }
func (m *UpdateCollectionDataProposal) ProposalType() string  { return ProposalTypeUpdateCollectionData }

func (m *UpdateCollectionDataProposal) ValidateBasic() error {
	if err := gov.ValidateAbstract(m); err != nil {
		return err
	}

	for _, d := range m.Data {
		if err := validateCollectionData(d); err != nil {
			return err
		}
	}

	return nil
}

// Implements Proposal Interface
var _ gov.Content = &AddCollectionDataProposal{}

func (m *AddCollectionDataProposal) ProposalRoute() string { return RouterKey }
func (m *AddCollectionDataProposal) ProposalType() string  { return ProposalTypeAddCollectionData }

func (m *AddCollectionDataProposal) ValidateBasic() error {
	if err := gov.ValidateAbstract(m); err != nil {
		return err
	}

	for _, d := range m.Data {
		if err := validateCollectionData(d); err != nil {
			return err
		}
	}

	return nil
}

// Implements Proposal Interface
var _ gov.Content = &RemoveCollectionDataProposal{}

func (m *RemoveCollectionDataProposal) ProposalRoute() string { return RouterKey }
func (m *RemoveCollectionDataProposal) ProposalType() string  { return ProposalTypeRemoveCollectionData }

func (m *RemoveCollectionDataProposal) ValidateBasic() error {
	if err := gov.ValidateAbstract(m); err != nil {
		return err
	}

	for _, d := range m.Index {
		if err := validateCollectionDataIndex(d); err != nil {
			return err
		}
	}

	return nil
}

// Implements Proposal Interface
var _ gov.Content = &RemoveCollectionProposal{}

func (m *RemoveCollectionProposal) ProposalRoute() string { return RouterKey }
func (m *RemoveCollectionProposal) ProposalType() string  { return ProposalTypeRemoveCollection }

func (m *RemoveCollectionProposal) ValidateBasic() error {
	return gov.ValidateAbstract(m)
}
