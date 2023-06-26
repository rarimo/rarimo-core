package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgDepositNative{}, "bridge/DepositNative", nil)
	cdc.RegisterConcrete(&MsgWithdrawNative{}, "bridge/WithdrawNative", nil)
	cdc.RegisterConcrete(&ChangeParamsProposal{}, "bridge/ChangeParamsProposal", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDepositNative{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdrawNative{},
	)
	registry.RegisterImplementations(
		(*govv1beta1.Content)(nil),
		&ChangeParamsProposal{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
