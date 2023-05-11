package pkg

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
	tokentypes "gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func GetContractUpgrade(operation types.Operation) (*types.ContractUpgrade, error) {
	if operation.OperationType == types.OpType_FEE_TOKEN_MANAGEMENT {
		op := new(types.ContractUpgrade)
		return op, proto.Unmarshal(operation.Details.Value, op)
	}
	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "invalid operation type")
}

func GetContractUpgradeContent(params tokentypes.NetworkParams, op *types.ContractUpgrade) (*operation.ContractUpgradeContent, error) {
	if params.Name != op.Chain {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "network names missmatch")
	}

	switch params.Type {
	case tokentypes.NetworkType_EVM:
		return &operation.ContractUpgradeContent{
			ChainName:                 op.Chain,
			Contract:                  op.TargetContract,
			NewImplementationContract: op.NewImplementationContract,
			Nonce:                     op.Nonce,
			Type:                      op.Type,
		}, nil
	case tokentypes.NetworkType_Solana:
		return &operation.ContractUpgradeContent{
			ChainName:     op.Chain,
			Contract:      op.TargetContract,
			ByteCodeHash:  op.Hash,
			BufferAccount: op.BufferAccount,
			Nonce:         op.Nonce,
		}, nil
	case tokentypes.NetworkType_Near:
		return &operation.ContractUpgradeContent{
			ChainName:    op.Chain,
			Contract:     op.TargetContract,
			ByteCodeHash: op.Hash,
			Nonce:        op.Nonce,
		}, nil
	}

	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid network type")
}
