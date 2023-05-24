package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&SetNetworkProposal{}, "tokenmanager/SetNetworkProposal", nil)
	cdc.RegisterConcrete(&UpdateTokenItemProposal{}, "tokenmanager/UpdateTokenItemProposal", nil)
	cdc.RegisterConcrete(&RemoveTokenItemProposal{}, "tokenmanager/RemoveTokenItemProposal", nil)
	cdc.RegisterConcrete(&CreateCollectionProposal{}, "tokenmanager/CreateCollectionProposal", nil)
	cdc.RegisterConcrete(&UpdateCollectionDataProposal{}, "tokenmanager/UpdateCollectionDataProposal", nil)
	cdc.RegisterConcrete(&AddCollectionDataProposal{}, "tokenmanager/AddCollectionDataProposal", nil)
	cdc.RegisterConcrete(&RemoveCollectionDataProposal{}, "tokenmanager/RemoveCollectionDataProposal", nil)
	cdc.RegisterConcrete(&RemoveCollectionProposal{}, "tokenmanager/RemoveCollectionProposal", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*govv1beta1.Content)(nil),
		&SetNetworkProposal{},
		&UpdateTokenItemProposal{},
		&RemoveTokenItemProposal{},
		&CreateCollectionProposal{},
		&UpdateCollectionDataProposal{},
		&AddCollectionDataProposal{},
		&RemoveCollectionDataProposal{},
		&RemoveCollectionProposal{},
	)

	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
