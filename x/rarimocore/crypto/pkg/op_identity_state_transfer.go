package pkg

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gogo/protobuf/proto"
	"github.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
	"github.com/rarimo/rarimo-core/x/rarimocore/types"
)

func GetIdentityStateTransfer(operation types.Operation) (*types.IdentityStateTransfer, error) {
	if operation.OperationType == types.OpType_IDENTITY_STATE_TRANSFER {
		change := new(types.IdentityStateTransfer)
		return change, proto.Unmarshal(operation.Details.Value, change)
	}
	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "invalid operation type")
}

func GetIdentityStateTransferContent(op *types.IdentityStateTransfer) (*operation.IdentityStateTransferContent, error) {
	return &operation.IdentityStateTransferContent{
		Contract:                hexutil.MustDecode(op.Contract),
		Id:                      operation.To32Bytes(hexutil.MustDecode(op.Id)),
		StateHash:               operation.To32Bytes(hexutil.MustDecode(op.StateHash)),
		StateCreatedAtTimestamp: operation.To32Bytes(operation.AmountBytes(op.StateCreatedAtTimestamp)),
		StateCreatedAtBlock:     operation.To32Bytes(operation.AmountBytes(op.StateCreatedAtBlock)),
		ReplacedStateHash:       operation.To32Bytes(hexutil.MustDecode(op.ReplacedStateHash)),
	}, nil
}
