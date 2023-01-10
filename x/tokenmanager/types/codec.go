package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&CreateTokenItemProposal{}, "tokenmanager/CreateTokenItemProposal", nil)
	cdc.RegisterConcrete(&RemoveTokenItemProposal{}, "tokenmanager/RemoveTokenItemProposal", nil)
	cdc.RegisterConcrete(&SetNetworkProposal{}, "tokenmanager/SetNetworkProposal", nil)
	cdc.RegisterConcrete(&CreateCollectionProposal{}, "tokenmanager/CreateCollectionProposal", nil)
	cdc.RegisterConcrete(&PutCollectionNetworkAddressProposal{}, "tokenmanager/PutCollectionNetworkAddressProposal", nil)
	cdc.RegisterConcrete(&RemoveCollectionNetworkAddressProposal{}, "tokenmanager/RemoveCollectionNetworkAddressProposal", nil)
	cdc.RegisterConcrete(&RemoveCollectionProposal{}, "tokenmanager/RemoveCollectionProposal", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&CreateTokenItemProposal{},
		&RemoveTokenItemProposal{},
		&SetNetworkProposal{},
		&CreateCollectionProposal{},
		&PutCollectionNetworkAddressProposal{},
		&RemoveCollectionNetworkAddressProposal{},
		&RemoveCollectionProposal{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
