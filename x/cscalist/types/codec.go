package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&EditCSCAListProposal{}, "cscalist/EditCSCAListProposal", nil)
	cdc.RegisterConcrete(&ReplaceCSCAListProposal{}, "cscalist/ReplaceCSCAListProposal", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*govv1beta1.Content)(nil),
		&EditCSCAListProposal{},
	)
	registry.RegisterImplementations(
		(*govv1beta1.Content)(nil),
		&ReplaceCSCAListProposal{},
	)
	// this line is used by starport scaffolding # 3
}
