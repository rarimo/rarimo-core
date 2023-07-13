package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/gogo/protobuf/proto"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgStake{}, "rarimocore/Stake", nil)
	cdc.RegisterConcrete(&MsgUnstake{}, "rarimocore/Unstake", nil)
	cdc.RegisterConcrete(&MsgCreateViolationReport{}, "rarimocore/CreateViolationReport", nil)
	cdc.RegisterConcrete(&MsgCreateConfirmation{}, "rarimocore/CreateConfirmation", nil)
	cdc.RegisterConcrete(&MsgCreateChangePartiesOp{}, "rarimocore/CreateChangeParties", nil)
	cdc.RegisterConcrete(&MsgSetupInitial{}, "rarimocore/SetupInitial", nil)
	cdc.RegisterConcrete(&Transfer{}, "rarimocore/Transfer", nil)
	cdc.RegisterConcrete(&ChangeParties{}, "rarimocore/ChangeParties", nil)
	cdc.RegisterConcrete(&FeeTokenManagement{}, "rarimocore/FeeTokenManagement", nil)
	cdc.RegisterConcrete(&ContractUpgrade{}, "rarimocore/ContractUpgrade", nil)
	cdc.RegisterConcrete(&IdentityDefaultTransfer{}, "rarimocore/IdentityDefaultTransfer", nil)
	cdc.RegisterConcrete(&UnfreezeSignerPartyProposal{}, "rarimocore/UnfreezeSignerPartyProposal", nil)
	cdc.RegisterConcrete(&ReshareKeysProposal{}, "rarimocore/ReshareKeysProposal", nil)
	cdc.RegisterConcrete(&SlashProposal{}, "rarimocore/SlashProposal", nil)
	cdc.RegisterConcrete(&DropPartiesProposal{}, "rarimocore/DropPartiesProposal", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgStake{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateViolationReport{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnstake{},
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
	registry.RegisterImplementations(
		(*govv1beta1.Content)(nil),
		&UnfreezeSignerPartyProposal{},
	)
	registry.RegisterImplementations(
		(*govv1beta1.Content)(nil),
		&ReshareKeysProposal{},
	)
	registry.RegisterImplementations(
		(*govv1beta1.Content)(nil),
		&SlashProposal{},
	)
	registry.RegisterImplementations(
		(*govv1beta1.Content)(nil),
		&DropPartiesProposal{},
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

	registry.RegisterInterface(
		"rarimo.rarimocore.rarimocore.FeeTokenManagement",
		(*proto.Message)(nil),
		&FeeTokenManagement{},
	)

	registry.RegisterInterface(
		"rarimo.rarimocore.rarimocore.ContractUpgrade",
		(*proto.Message)(nil),
		&ContractUpgrade{},
	)

	registry.RegisterInterface(
		"rarimo.rarimocore.rarimocore.IdentityDefaultTransfer",
		(*proto.Message)(nil),
		&IdentityDefaultTransfer{},
	)

	registry.RegisterInterface(
		"rarimo.rarimocore.rarimocore.ViolationReport",
		(*proto.Message)(nil),
		&ViolationReport{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
