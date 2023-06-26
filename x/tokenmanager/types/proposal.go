package types

import (
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	ProposalTypeSetNetwork           = "tokenmanager/SetNetwork"
	ProposalTypeUpdateTokenItem      = "tokenmanager/UpdateTokenItem"
	ProposalTypeRemoveTokenItem      = "tokenmanager/RemoveTokenItem"
	ProposalTypeCreateCollection     = "tokenmanager/CreateCollection"
	ProposalTypeUpdateCollectionData = "tokenmanager/UpdateCollectionData"
	ProposalTypeAddCollectionData    = "tokenmanager/AddCollectionData"
	ProposalTypeRemoveCollectionData = "tokenmanager/RemoveCollectionData"
	ProposalTypeRemoveCollection     = "tokenmanager/RemoveCollection"
)

func init() {
	govv1beta1.RegisterProposalType(ProposalTypeSetNetwork)
	govv1beta1.RegisterProposalType(ProposalTypeUpdateTokenItem)
	govv1beta1.RegisterProposalType(ProposalTypeRemoveTokenItem)
	govv1beta1.RegisterProposalType(ProposalTypeCreateCollection)
	govv1beta1.RegisterProposalType(ProposalTypeUpdateCollectionData)
	govv1beta1.RegisterProposalType(ProposalTypeAddCollectionData)
	govv1beta1.RegisterProposalType(ProposalTypeRemoveCollectionData)
	govv1beta1.RegisterProposalType(ProposalTypeRemoveCollection)

	govv1beta1.ModuleCdc.RegisterConcrete(&SetNetworkProposal{}, "tokenmanager/SetNetworkProposal", nil)
	govv1beta1.ModuleCdc.RegisterConcrete(&UpdateTokenItemProposal{}, "tokenmanager/UpdateTokenItemProposal", nil)
	govv1beta1.ModuleCdc.RegisterConcrete(&RemoveTokenItemProposal{}, "tokenmanager/RemoveTokenItemProposal", nil)
	govv1beta1.ModuleCdc.RegisterConcrete(&CreateCollectionProposal{}, "tokenmanager/CreateCollectionProposal", nil)
	govv1beta1.ModuleCdc.RegisterConcrete(&UpdateCollectionDataProposal{}, "tokenmanager/UpdateCollectionDataProposal", nil)
	govv1beta1.ModuleCdc.RegisterConcrete(&AddCollectionDataProposal{}, "tokenmanager/AddCollectionDataProposal", nil)
	govv1beta1.ModuleCdc.RegisterConcrete(&RemoveCollectionDataProposal{}, "tokenmanager/RemoveCollectionDataProposal", nil)
	govv1beta1.ModuleCdc.RegisterConcrete(&RemoveCollectionProposal{}, "tokenmanager/RemoveCollectionProposal", nil)
}

// Implements Proposal Interface
var _ govv1beta1.Content = &SetNetworkProposal{}

func (m *SetNetworkProposal) ProposalRoute() string { return RouterKey }
func (m *SetNetworkProposal) ProposalType() string  { return ProposalTypeSetNetwork }

func (m *SetNetworkProposal) ValidateBasic() error {
	if err := govv1beta1.ValidateAbstract(m); err != nil {
		return err
	}

	return validateNetwork(m.NetworkParams)
}

// Implements Proposal Interface
var _ govv1beta1.Content = &UpdateTokenItemProposal{}

func (m *UpdateTokenItemProposal) ProposalRoute() string { return RouterKey }
func (m *UpdateTokenItemProposal) ProposalType() string  { return ProposalTypeUpdateTokenItem }

func (m *UpdateTokenItemProposal) ValidateBasic() error {
	if err := govv1beta1.ValidateAbstract(m); err != nil {
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
var _ govv1beta1.Content = &RemoveTokenItemProposal{}

func (m *RemoveTokenItemProposal) ProposalRoute() string { return RouterKey }
func (m *RemoveTokenItemProposal) ProposalType() string  { return ProposalTypeRemoveTokenItem }

func (m *RemoveTokenItemProposal) ValidateBasic() error {
	if err := govv1beta1.ValidateAbstract(m); err != nil {
		return err
	}

	return nil
}

// Implements Proposal Interface
var _ govv1beta1.Content = &CreateCollectionProposal{}

func (m *CreateCollectionProposal) ProposalRoute() string { return RouterKey }
func (m *CreateCollectionProposal) ProposalType() string  { return ProposalTypeCreateCollection }

func (m *CreateCollectionProposal) ValidateBasic() error {
	if err := govv1beta1.ValidateAbstract(m); err != nil {
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
var _ govv1beta1.Content = &UpdateCollectionDataProposal{}

func (m *UpdateCollectionDataProposal) ProposalRoute() string { return RouterKey }
func (m *UpdateCollectionDataProposal) ProposalType() string  { return ProposalTypeUpdateCollectionData }

func (m *UpdateCollectionDataProposal) ValidateBasic() error {
	if err := govv1beta1.ValidateAbstract(m); err != nil {
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
var _ govv1beta1.Content = &AddCollectionDataProposal{}

func (m *AddCollectionDataProposal) ProposalRoute() string { return RouterKey }
func (m *AddCollectionDataProposal) ProposalType() string  { return ProposalTypeAddCollectionData }

func (m *AddCollectionDataProposal) ValidateBasic() error {
	if err := govv1beta1.ValidateAbstract(m); err != nil {
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
var _ govv1beta1.Content = &RemoveCollectionDataProposal{}

func (m *RemoveCollectionDataProposal) ProposalRoute() string { return RouterKey }
func (m *RemoveCollectionDataProposal) ProposalType() string  { return ProposalTypeRemoveCollectionData }

func (m *RemoveCollectionDataProposal) ValidateBasic() error {
	if err := govv1beta1.ValidateAbstract(m); err != nil {
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
var _ govv1beta1.Content = &RemoveCollectionProposal{}

func (m *RemoveCollectionProposal) ProposalRoute() string { return RouterKey }
func (m *RemoveCollectionProposal) ProposalType() string  { return ProposalTypeRemoveCollection }

func (m *RemoveCollectionProposal) ValidateBasic() error {
	return govv1beta1.ValidateAbstract(m)
}
