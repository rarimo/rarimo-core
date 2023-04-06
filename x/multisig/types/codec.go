package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgChangeGroup{}, "multisig/ChangeGroup", nil)
	cdc.RegisterConcrete(&MsgCreateGroup{}, "multisig/CreateGroup", nil)
	cdc.RegisterConcrete(&MsgSubmitProposal{}, "multisig/MsgSubmitProposal", nil)
	cdc.RegisterConcrete(&MsgVote{}, "multisig/MsgVote", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgChangeGroup{},
		&MsgCreateGroup{},
		&MsgSubmitProposal{},
		&MsgVote{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
