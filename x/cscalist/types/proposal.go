package types

import (
	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	ProposalTypeEditCSCAList    = "cscalist/EditCSCAList"
	ProposalTypeReplaceCSCAList = "cscalist/ReplaceCSCAList"
)

func init() {
	govv1beta1.RegisterProposalType(ProposalTypeReplaceCSCAList)
	govv1beta1.ModuleCdc.RegisterConcrete(&ReplaceCSCAListProposal{}, "rarimocore/ReplaceCSCAListProposal", nil)
	govv1beta1.RegisterProposalType(ProposalTypeEditCSCAList)
	govv1beta1.ModuleCdc.RegisterConcrete(&EditCSCAListProposal{}, "rarimocore/EditCSCAListProposal", nil)
}

// Implements Proposal Interface
var _ govv1beta1.Content = &EditCSCAListProposal{}
var _ govv1beta1.Content = &ReplaceCSCAListProposal{}

func (m *EditCSCAListProposal) ProposalRoute() string { return RouterKey }

func (m *EditCSCAListProposal) ProposalType() string { return ProposalTypeEditCSCAList }

func (m *EditCSCAListProposal) ValidateBasic() error {
	leaves := append(m.GetToAdd(), m.GetToRemove()...)
	if len(leaves) == 0 {
		return errors.Wrap(types.ErrInvalidProposalContent, "nothing to add or remove")
	}

	for _, leaf := range leaves {
		if _, err := hexutil.Decode(leaf); err != nil {
			return errors.Wrapf(types.ErrInvalidProposalContent, "invalid hex string: %s", leaf)
		}
	}

	return govv1beta1.ValidateAbstract(m)
}

func (m *ReplaceCSCAListProposal) ProposalRoute() string { return RouterKey }

func (m *ReplaceCSCAListProposal) ProposalType() string { return ProposalTypeReplaceCSCAList }

func (m *ReplaceCSCAListProposal) ValidateBasic() error {
	if len(m.GetLeaves()) == 0 {
		return errors.Wrap(types.ErrInvalidProposalContent, "empty leaves")
	}

	for _, leaf := range m.GetLeaves() {
		if _, err := hexutil.Decode(leaf); err != nil {
			return errors.Wrapf(types.ErrInvalidProposalContent, "invalid hex string: %s", leaf)
		}
	}

	return govv1beta1.ValidateAbstract(m)
}
