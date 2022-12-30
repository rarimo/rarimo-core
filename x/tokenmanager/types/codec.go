package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&SetTokenInfoProposal{}, "tokenmanager/SetTokenInfoProposal", nil)
	cdc.RegisterConcrete(&SetTokenItemProposal{}, "tokenmanager/SetTokenItemProposal", nil)
	cdc.RegisterConcrete(&RemoveTokenItemProposal{}, "tokenmanager/RemoveTokenItemProposal", nil)
	cdc.RegisterConcrete(&RemoveTokenInfoProposal{}, "tokenmanager/RemoveTokenInfoProposal", nil)
	cdc.RegisterConcrete(&SetNetworkProposal{}, "tokenmanager/SetNetworkProposal", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&SetTokenInfoProposal{},
		&SetTokenItemProposal{},
		&RemoveTokenItemProposal{},
		&RemoveTokenInfoProposal{},
		&SetNetworkProposal{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
