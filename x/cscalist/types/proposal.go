package types

import (
	"encoding/hex"

	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	ProposalTypeEditCSCAList = "cscalist/EditCSCAList"
)

func init() {
	govv1beta1.RegisterProposalType(ProposalTypeEditCSCAList)
	govv1beta1.ModuleCdc.RegisterConcrete(&EditCSCAListProposal{}, "rarimocore/EditCSCAListProposal", nil)
}

// Implements Proposal Interface
var _ govv1beta1.Content = &EditCSCAListProposal{}

func (m *EditCSCAListProposal) ProposalRoute() string { return RouterKey }
func (m *EditCSCAListProposal) ProposalType() string  { return ProposalTypeEditCSCAList }

func (m *EditCSCAListProposal) ValidateBasic() error {
	for _, leaf := range m.Leaves {
		if _, err := hex.DecodeString(leaf); err != nil {
			return errors.Wrapf(types.ErrInvalidProposalContent, "invalid hex string: %s", leaf)
		}
	}

	return govv1beta1.ValidateAbstract(m)
}
