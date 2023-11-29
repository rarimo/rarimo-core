package pkg

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gogo/protobuf/proto"
	"github.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
	"github.com/rarimo/rarimo-core/x/rarimocore/types"
)

func GetIdentityGISTTransfer(operation types.Operation) (*types.IdentityGISTTransfer, error) {
	if operation.OperationType == types.OpType_IDENTITY_GIST_TRANSFER {
		change := new(types.IdentityGISTTransfer)
		return change, proto.Unmarshal(operation.Details.Value, change)
	}

	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "invalid operation type")
}

func GetIdentityGISTTransferContent(op *types.IdentityGISTTransfer) (*operation.IdentityGISTTransferContent, error) {
	return &operation.IdentityGISTTransferContent{
		Contract:               hexutil.MustDecode(op.Contract),
		GISTHash:               operation.To32Bytes(hexutil.MustDecode(op.GISTHash)),
		GISTCreatedAtTimestamp: operation.To32Bytes(operation.AmountBytes(op.GISTCreatedAtTimestamp)),
		GISTCreatedAtBlock:     operation.To32Bytes(operation.AmountBytes(op.GISTCreatedAtBlock)),
		ReplacedGISTHash:       operation.To32Bytes(hexutil.MustDecode(op.ReplacedGISTHash)),
	}, nil
}
