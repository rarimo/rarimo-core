package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/gogo/protobuf/proto"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateTransferOp{}, "rarimocore/CreateTransfer", nil)
	cdc.RegisterConcrete(&MsgCreateConfirmation{}, "rarimocore/CreateConfirmation", nil)
	cdc.RegisterConcrete(&MsgCreateChangeKeyECDSA{}, "rarimocore/CreateChangeKeyECDSA", nil)
	cdc.RegisterConcrete(&Transfer{}, "rarimocore/Transfer", nil)
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
		&MsgCreateChangeKeyECDSA{},
	)

	registry.RegisterInterface(
		"rarifyprotocol.rarimocore.rarimocore.Transfer",
		(*proto.Message)(nil),
		&Transfer{},
	)

	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
