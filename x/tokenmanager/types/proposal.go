package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	ProposalTypeCreateTokenItem                = "set_token_item"
	ProposalTypeRemoveTokenItem                = "remove_token_item"
	ProposalTypeRemoveTokenInfo                = "remove_token_info"
	ProposalTypeSetNetwork                     = "set_network"
	ProposalTypeCreateCollection               = "create_collection"
	ProposalTypePutCollectionNetworkAddress    = "put_collection_network_address"
	ProposalTypeRemoveCollectionNetworkAddress = "remove_collection_network_address"
	ProposalTypeRemoveCollection               = "remove_collection"
)

func init() {
	gov.RegisterProposalType(ProposalTypeCreateTokenItem)
	gov.RegisterProposalType(ProposalTypeRemoveTokenItem)
	gov.RegisterProposalType(ProposalTypeRemoveTokenInfo)
	gov.RegisterProposalType(ProposalTypeSetNetwork)
	gov.RegisterProposalType(ProposalTypeCreateCollection)
	gov.RegisterProposalType(ProposalTypePutCollectionNetworkAddress)
	gov.RegisterProposalType(ProposalTypeRemoveCollectionNetworkAddress)
	gov.RegisterProposalType(ProposalTypeRemoveCollection)
	gov.RegisterProposalTypeCodec(&CreateTokenItemProposal{}, "rarimocore/CreateTokenItemProposal")
	gov.RegisterProposalTypeCodec(&RemoveTokenItemProposal{}, "rarimocore/RemoveTokenItemProposal")
	gov.RegisterProposalTypeCodec(&SetNetworkProposal{}, "rarimocore/SetNetworkProposal")
	gov.RegisterProposalTypeCodec(&CreateCollectionProposal{}, "rarimocore/CreateCollectionProposal")
	gov.RegisterProposalTypeCodec(&PutCollectionNetworkAddressProposal{}, "rarimocore/PutCollectionNetworkAddressProposal")
	gov.RegisterProposalTypeCodec(&RemoveCollectionNetworkAddressProposal{}, "rarimocore/RemoveCollectionNetworkAddressProposal")
	gov.RegisterProposalTypeCodec(&RemoveCollectionProposal{}, "rarimocore/RemoveCollectionProposal")
}

// Implements Proposal Interface
var _ gov.Content = &CreateTokenItemProposal{}

func (m *CreateTokenItemProposal) ProposalRoute() string { return RouterKey }
func (m *CreateTokenItemProposal) ProposalType() string  { return ProposalTypeCreateTokenItem }

func (m *CreateTokenItemProposal) ValidateBasic() error {
	if m.Item == nil {
		return sdkerrors.Wrapf(gov.ErrInvalidProposalContent, "token item should not be nil")
	}

	if m.Item.Index != "" {
		return sdkerrors.Wrapf(gov.ErrInvalidProposalContent, "token item index should be empty as it set by the system")
	}

	if m.Item.Collection == "" {
		return sdkerrors.Wrapf(gov.ErrInvalidProposalContent, "token item collection should not be empty")
	}

	return gov.ValidateAbstract(m)
}

// Implements Proposal Interface
var _ gov.Content = &RemoveTokenItemProposal{}

func (m *RemoveTokenItemProposal) ProposalRoute() string { return RouterKey }
func (m *RemoveTokenItemProposal) ProposalType() string  { return ProposalTypeRemoveTokenItem }

func (m *RemoveTokenItemProposal) ValidateBasic() error {
	return gov.ValidateAbstract(m)
}

// Implements Proposal Interface
var _ gov.Content = &SetNetworkProposal{}

func (m *SetNetworkProposal) ProposalRoute() string { return RouterKey }
func (m *SetNetworkProposal) ProposalType() string  { return ProposalTypeSetNetwork }

func (m *SetNetworkProposal) ValidateBasic() error {
	return gov.ValidateAbstract(m)
}

// Implements Proposal Interface
var _ gov.Content = &CreateCollectionProposal{}

func (m *CreateCollectionProposal) ProposalRoute() string { return RouterKey }
func (m *CreateCollectionProposal) ProposalType() string  { return ProposalTypeCreateCollection }

func (m *CreateCollectionProposal) ValidateBasic() error {
	return gov.ValidateAbstract(m)
}

// Implements Proposal Interface
var _ gov.Content = &PutCollectionNetworkAddressProposal{}

func (m *PutCollectionNetworkAddressProposal) ProposalRoute() string { return RouterKey }
func (m *PutCollectionNetworkAddressProposal) ProposalType() string {
	return ProposalTypePutCollectionNetworkAddress
}

func (m *PutCollectionNetworkAddressProposal) ValidateBasic() error {
	return gov.ValidateAbstract(m)
}

// Implements Proposal Interface
var _ gov.Content = &RemoveCollectionNetworkAddressProposal{}

func (m *RemoveCollectionNetworkAddressProposal) ProposalRoute() string { return RouterKey }
func (m *RemoveCollectionNetworkAddressProposal) ProposalType() string {
	return ProposalTypeRemoveCollectionNetworkAddress
}

func (m *RemoveCollectionNetworkAddressProposal) ValidateBasic() error {
	return gov.ValidateAbstract(m)
}

// Implements Proposal Interface
var _ gov.Content = &RemoveCollectionProposal{}

func (m *RemoveCollectionProposal) ProposalRoute() string { return RouterKey }
func (m *RemoveCollectionProposal) ProposalType() string  { return ProposalTypeRemoveCollection }

func (m *RemoveCollectionProposal) ValidateBasic() error {
	return gov.ValidateAbstract(m)
}
