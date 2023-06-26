package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&UpgradeContractProposal{}, "tokenmanager/UpgradeContractProposal", nil)
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
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&UpgradeContractProposal{},
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

	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
