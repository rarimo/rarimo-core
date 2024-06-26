package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/gogo/protobuf/proto"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&AddNetworkProposal{}, "tokenmanager/AddNetworkProposal", nil)
	cdc.RegisterConcrete(&RemoveNetworkProposal{}, "tokenmanager/RemoveNetworkProposal", nil)

	cdc.RegisterConcrete(&AddFeeTokenProposal{}, "tokenmanager/AddFeeTokenProposal", nil)
	cdc.RegisterConcrete(&UpdateFeeTokenProposal{}, "tokenmanager/UpdateFeeTokenProposal", nil)
	cdc.RegisterConcrete(&RemoveFeeTokenProposal{}, "tokenmanager/RemoveFeeTokenProposal", nil)
	cdc.RegisterConcrete(&WithdrawFeeProposal{}, "tokenmanager/WithdrawFeeProposal", nil)

	cdc.RegisterConcrete(&UpdateTokenItemProposal{}, "tokenmanager/UpdateTokenItemProposal", nil)
	cdc.RegisterConcrete(&RemoveTokenItemProposal{}, "tokenmanager/RemoveTokenItemProposal", nil)
	cdc.RegisterConcrete(&CreateCollectionProposal{}, "tokenmanager/CreateCollectionProposal", nil)
	cdc.RegisterConcrete(&UpdateCollectionDataProposal{}, "tokenmanager/UpdateCollectionDataProposal", nil)
	cdc.RegisterConcrete(&AddCollectionDataProposal{}, "tokenmanager/AddCollectionDataProposal", nil)
	cdc.RegisterConcrete(&RemoveCollectionDataProposal{}, "tokenmanager/RemoveCollectionDataProposal", nil)
	cdc.RegisterConcrete(&RemoveCollectionProposal{}, "tokenmanager/RemoveCollectionProposal", nil)

	cdc.RegisterConcrete(&BridgeNetworkParams{}, "tokenmanager/BridgeNetworkParams", nil)
	cdc.RegisterConcrete(&IdentityNetworkParams{}, "tokenmanager/IdentityNetworkParams", nil)
	cdc.RegisterConcrete(&FeeNetworkParams{}, "tokenmanager/FeeNetworkParams", nil)

	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&AddNetworkProposal{},
		&RemoveNetworkProposal{},

		&AddFeeTokenProposal{},
		&UpdateFeeTokenProposal{},
		&RemoveFeeTokenProposal{},
		&WithdrawFeeProposal{},

		&UpdateTokenItemProposal{},
		&RemoveTokenItemProposal{},
		&CreateCollectionProposal{},
		&UpdateCollectionDataProposal{},
		&AddCollectionDataProposal{},
		&RemoveCollectionDataProposal{},
		&RemoveCollectionProposal{},
	)

	registry.RegisterInterface(
		"rarimo.rarimocore.tokenmanager.BridgeNetworkParams",
		(*proto.Message)(nil),
		&BridgeNetworkParams{},
	)

	registry.RegisterInterface(
		"rarimo.rarimocore.tokenmanager.FeeNetworkParams",
		(*proto.Message)(nil),
		&FeeNetworkParams{},
	)

	registry.RegisterInterface(
		"rarimo.rarimocore.tokenmanager.IdentityNetworkParams",
		(*proto.Message)(nil),
		&IdentityNetworkParams{},
	)

	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
