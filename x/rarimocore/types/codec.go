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
	cdc.RegisterConcrete(&MsgCreateConfirmation{}, "rarimocore/CreateConfirmation", nil)
	cdc.RegisterConcrete(&MsgCreateChangePartiesOp{}, "rarimocore/CreateChangeParties", nil)
	cdc.RegisterConcrete(&MsgSetupInitial{}, "rarimocore/SetupInitial", nil)
	cdc.RegisterConcrete(&Transfer{}, "rarimocore/Transfer", nil)
	cdc.RegisterConcrete(&ChangeParties{}, "rarimocore/ChangeParties", nil)
	cdc.RegisterConcrete(&AddSignerPartyProposal{}, "rarimocore/AddSignerPartyProposal", nil)
	cdc.RegisterConcrete(&RemoveSignerPartyProposal{}, "rarimocore/RemoveSignerPartyProposal", nil)
	cdc.RegisterConcrete(&ReshareKeysProposal{}, "rarimocore/ReshareKeysProposal", nil)
	cdc.RegisterConcrete(&ChangeThresholdProposal{}, "rarimocore/ChangeThresholdProposal", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateConfirmation{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateChangePartiesOp{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetupInitial{},
	)
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&AddSignerPartyProposal{},
	)
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&RemoveSignerPartyProposal{},
	)
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&ReshareKeysProposal{},
	)
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&ChangeThresholdProposal{},
	)

	registry.RegisterInterface(
		"rarimo.rarimocore.rarimocore.Transfer",
		(*proto.Message)(nil),
		&Transfer{},
	)

	registry.RegisterInterface(
		"rarimo.rarimocore.rarimocore.ChangeParties",
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
