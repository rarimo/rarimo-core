package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateDeposit{}, "rarimocore/CreateDeposit", nil)
	cdc.RegisterConcrete(&MsgUpdateDeposit{}, "rarimocore/UpdateDeposit", nil)
	cdc.RegisterConcrete(&MsgDeleteDeposit{}, "rarimocore/DeleteDeposit", nil)
	cdc.RegisterConcrete(&MsgCreateConfirmation{}, "rarimocore/CreateConfirmation", nil)
	cdc.RegisterConcrete(&MsgUpdateConfirmation{}, "rarimocore/UpdateConfirmation", nil)
	cdc.RegisterConcrete(&MsgDeleteConfirmation{}, "rarimocore/DeleteConfirmation", nil)
	cdc.RegisterConcrete(&MsgCreateChangeKeyECDSA{}, "rarimocore/CreateChangeKeyECDSA", nil)
	cdc.RegisterConcrete(&MsgUpdateChangeKeyECDSA{}, "rarimocore/UpdateChangeKeyECDSA", nil)
	cdc.RegisterConcrete(&MsgDeleteChangeKeyECDSA{}, "rarimocore/DeleteChangeKeyECDSA", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDeposit{},
		&MsgUpdateDeposit{},
		&MsgDeleteDeposit{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateConfirmation{},
		&MsgUpdateConfirmation{},
		&MsgDeleteConfirmation{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateChangeKeyECDSA{},
		&MsgUpdateChangeKeyECDSA{},
		&MsgDeleteChangeKeyECDSA{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
