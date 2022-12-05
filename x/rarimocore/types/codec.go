package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/gogo/protobuf/proto"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateTransferOp{}, "rarimocore/CreateTransfer", nil)
	cdc.RegisterConcrete(&MsgCreateConfirmation{}, "rarimocore/CreateConfirmation", nil)
	cdc.RegisterConcrete(&MsgCreateChangePartiesOp{}, "rarimocore/CreateChangeParties", nil)
	cdc.RegisterConcrete(&MsgSetupInitial{}, "rarimocore/SetupInitial", nil)
	cdc.RegisterConcrete(&Transfer{}, "rarimocore/Transfer", nil)
	cdc.RegisterConcrete(&ChangeParties{}, "rarimocore/ChangeParties", nil)
	cdc.RegisterConcrete(&MsgCreateAddSignerPartyProposal{}, "rarimocore/CreateAddSignerPartyProposal", nil)
	cdc.RegisterConcrete(&AddSignerPartyProposal{}, "rarimocore/AddSignerPartyProposal", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateTransferOp{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateConfirmation{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateChangePartiesOp{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetupInitial{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateAddSignerPartyProposal{},
	)

	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&AddSignerPartyProposal{},
	)

	registry.RegisterInterface(
		"rarifyprotocol.rarimocore.rarimocore.Transfer",
		(*proto.Message)(nil),
		&Transfer{},
	)

	registry.RegisterInterface(
		"rarifyprotocol.rarimocore.rarimocore.ChangeParties",
		(*proto.Message)(nil),
		&ChangeParties{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
