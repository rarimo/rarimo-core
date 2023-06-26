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
	cdc.RegisterConcrete(&MsgCreateTransferOp{}, "oraclemanager/CreateTransferOp", nil)
	cdc.RegisterConcrete(&MsgStake{}, "oraclemanager/Stake", nil)
	cdc.RegisterConcrete(&MsgUnstake{}, "oraclemanager/Unstake", nil)
	cdc.RegisterConcrete(&MsgUnjail{}, "oraclemanager/Unjail", nil)
	cdc.RegisterConcrete(&MsgVote{}, "oraclemanager/Vote", nil)
	cdc.RegisterConcrete(&OracleUnfreezeProposal{}, "oraclemanager/OracleUnfreezeProposal", nil)
	cdc.RegisterConcrete(&ChangeParamsProposal{}, "oraclemanager/ChangeParamsProposal", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateTransferOp{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgStake{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnstake{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnjail{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgVote{},
	)
	registry.RegisterImplementations(
		(*govv1beta1.Content)(nil),
		&OracleUnfreezeProposal{},
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
