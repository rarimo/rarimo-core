package types

import (
	gov "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	ProposalTypeAddNetwork            = "tokenmanager/AddNetwork"
	ProposalTypeRemoveNetwork         = "tokenmanager/RemoveNetwork"
	ProposalTypeUpdateContractAddress = "tokenmanager/UpdateContractAddressNetwork"
	ProposalTypeAddFeeToken           = "tokenmanager/AddFeeToken"
	ProposalTypeRemoveFeeToken        = "tokenmanager/RemoveFeeToken"
	ProposalTypeUpdateFeeToken        = "tokenmanager/UpdateFeeToken"
	ProposalTypeWithdrawFee           = "tokenmanager/WithdrawFee"

	ProposalTypeUpdateTokenItem      = "tokenmanager/UpdateTokenItem"
	ProposalTypeRemoveTokenItem      = "tokenmanager/RemoveTokenItem"
	ProposalTypeCreateCollection     = "tokenmanager/CreateCollection"
	ProposalTypeUpdateCollectionData = "tokenmanager/UpdateCollectionData"
	ProposalTypeAddCollectionData    = "tokenmanager/AddCollectionData"
	ProposalTypeRemoveCollectionData = "tokenmanager/RemoveCollectionData"
	ProposalTypeRemoveCollection     = "tokenmanager/RemoveCollection"
)

func init() {
	gov.RegisterProposalType(ProposalTypeAddNetwork)
	gov.RegisterProposalType(ProposalTypeRemoveNetwork)
	gov.RegisterProposalType(ProposalTypeUpdateContractAddress)

	gov.RegisterProposalType(ProposalTypeAddFeeToken)
	gov.RegisterProposalType(ProposalTypeRemoveFeeToken)
	gov.RegisterProposalType(ProposalTypeUpdateFeeToken)
	gov.RegisterProposalType(ProposalTypeWithdrawFee)

	gov.RegisterProposalType(ProposalTypeUpdateTokenItem)
	gov.RegisterProposalType(ProposalTypeRemoveTokenItem)
	gov.RegisterProposalType(ProposalTypeCreateCollection)
	gov.RegisterProposalType(ProposalTypeUpdateCollectionData)
	gov.RegisterProposalType(ProposalTypeAddCollectionData)
	gov.RegisterProposalType(ProposalTypeRemoveCollectionData)
	gov.RegisterProposalType(ProposalTypeRemoveCollection)

	gov.ModuleCdc.RegisterConcrete(&AddNetworkProposal{}, "tokenmanager/AddNetworkProposal", nil)
	gov.ModuleCdc.RegisterConcrete(&RemoveNetworkProposal{}, "tokenmanager/RemoveNetworkProposal", nil)

	gov.ModuleCdc.RegisterConcrete(&AddFeeTokenProposal{}, "tokenmanager/AddFeeTokenProposal", nil)
	gov.ModuleCdc.RegisterConcrete(&RemoveFeeTokenProposal{}, "tokenmanager/RemoveFeeTokenProposal", nil)
	gov.ModuleCdc.RegisterConcrete(&UpdateFeeTokenProposal{}, "tokenmanager/UpdateFeeTokenProposal", nil)
	gov.ModuleCdc.RegisterConcrete(&WithdrawFeeProposal{}, "tokenmanager/WithdrawFeeProposal", nil)

	gov.ModuleCdc.RegisterConcrete(&UpdateTokenItemProposal{}, "tokenmanager/UpdateTokenItemProposal", nil)
	gov.ModuleCdc.RegisterConcrete(&RemoveTokenItemProposal{}, "tokenmanager/RemoveTokenItemProposal", nil)
	gov.ModuleCdc.RegisterConcrete(&CreateCollectionProposal{}, "tokenmanager/CreateCollectionProposal", nil)
	gov.ModuleCdc.RegisterConcrete(&UpdateCollectionDataProposal{}, "tokenmanager/UpdateCollectionDataProposal", nil)
	gov.ModuleCdc.RegisterConcrete(&AddCollectionDataProposal{}, "tokenmanager/AddCollectionDataProposal", nil)
	gov.ModuleCdc.RegisterConcrete(&RemoveCollectionDataProposal{}, "tokenmanager/RemoveCollectionDataProposal", nil)
	gov.ModuleCdc.RegisterConcrete(&RemoveCollectionProposal{}, "tokenmanager/RemoveCollectionProposal", nil)
}

// Implements Proposal Interface
var _ gov.Content = &AddNetworkProposal{}

func (m *AddNetworkProposal) ProposalRoute() string { return RouterKey }
func (m *AddNetworkProposal) ProposalType() string  { return ProposalTypeAddNetwork }

func (m *AddNetworkProposal) ValidateBasic() error {
	if err := gov.ValidateAbstract(m); err != nil {
		return err
	}

	return validateNetwork(&m.Network)
}

// Implements Proposal Interface
var _ gov.Content = &RemoveNetworkProposal{}

func (m *RemoveNetworkProposal) ProposalRoute() string { return RouterKey }
func (m *RemoveNetworkProposal) ProposalType() string  { return ProposalTypeRemoveNetwork }

func (m *RemoveNetworkProposal) ValidateBasic() error {
	if err := gov.ValidateAbstract(m); err != nil {
		return err
	}

	return nil
}

// Implements Proposal Interface
var _ gov.Content = &AddFeeTokenProposal{}

func (m *AddFeeTokenProposal) ProposalRoute() string { return RouterKey }
func (m *AddFeeTokenProposal) ProposalType() string {
	return ProposalTypeAddFeeToken
}

func (m *AddFeeTokenProposal) ValidateBasic() error {
	if err := gov.ValidateAbstract(m); err != nil {
		return err
	}

	return nil
}

// Implements Proposal Interface
var _ gov.Content = &RemoveFeeTokenProposal{}

func (m *RemoveFeeTokenProposal) ProposalRoute() string { return RouterKey }
func (m *RemoveFeeTokenProposal) ProposalType() string {
	return ProposalTypeRemoveFeeToken
}

func (m *RemoveFeeTokenProposal) ValidateBasic() error {
	if err := gov.ValidateAbstract(m); err != nil {
		return err
	}

	return nil
}

// Implements Proposal Interface
var _ gov.Content = &UpdateFeeTokenProposal{}

func (m *UpdateFeeTokenProposal) ProposalRoute() string { return RouterKey }
func (m *UpdateFeeTokenProposal) ProposalType() string {
	return ProposalTypeUpdateFeeToken
}

func (m *UpdateFeeTokenProposal) ValidateBasic() error {
	if err := gov.ValidateAbstract(m); err != nil {
		return err
	}

	return nil
}

// Implements Proposal Interface
var _ gov.Content = &WithdrawFeeProposal{}

func (m *WithdrawFeeProposal) ProposalRoute() string { return RouterKey }
func (m *WithdrawFeeProposal) ProposalType() string {
	return ProposalTypeWithdrawFee
}

func (m *WithdrawFeeProposal) ValidateBasic() error {
	if err := gov.ValidateAbstract(m); err != nil {
		return err
	}

	return nil
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
