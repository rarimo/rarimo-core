package pkg

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gogo/protobuf/proto"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func GetIdentityDefaultTransfer(operation types.Operation) (*types.IdentityDefaultTransfer, error) {
	if operation.OperationType == types.OpType_IDENTITY_DEFAULT_TRANSFER {
		change := new(types.IdentityDefaultTransfer)
		return change, proto.Unmarshal(operation.Details.Value, change)
	}
	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "invalid operation type")
}

func GetIdentityDefaultTransferContent(op *types.IdentityDefaultTransfer) (*operation.IdentityDefaultTransferContent, error) {
	return &operation.IdentityDefaultTransferContent{
		Contract:                 hexutil.MustDecode(op.Contract),
		GISTHash:                 hexutil.MustDecode(op.GISTHash),
		Id:                       operation.To32Bytes(hexutil.MustDecode(op.Id)),
		StateHash:                hexutil.MustDecode(op.StateHash),
		StateCreatedAtTimestamp:  operation.To32Bytes(operation.AmountBytes(op.StateCreatedAtTimestamp)),
		StateCreatedAtBlock:      operation.To32Bytes(operation.AmountBytes(op.StateCreatedAtBlock)),
		StateReplacedAtTimestamp: operation.To32Bytes(operation.AmountBytes(op.StateReplacedAtTimestamp)),
		StateReplacedAtBlock:     operation.To32Bytes(operation.AmountBytes(op.StateReplacedAtBlock)),
		StateReplacedBy:          hexutil.MustDecode(op.StateReplacedBy),
		GISTReplacedBy:           hexutil.MustDecode(op.GISTReplacedBy),
		GISTCreatedAtTimestamp:   operation.To32Bytes(operation.AmountBytes(op.GISTCreatedAtTimestamp)),
		GISTCreatedAtBlock:       operation.To32Bytes(operation.AmountBytes(op.GISTCreatedAtBlock)),
		GISTReplacedAtTimestamp:  operation.To32Bytes(operation.AmountBytes(op.GISTReplacedAtTimestamp)),
		GISTReplacedAtBlock:      operation.To32Bytes(operation.AmountBytes(op.GISTReplacedAtBlock)),
		ReplacedStateHash:        hexutil.MustDecode(op.ReplacedStateHash),
	}, nil
}
