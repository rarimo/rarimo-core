package pkg

import (
	"math/big"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gogo/protobuf/proto"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func GetIdentityAggregatedTransfer(operation types.Operation) (*types.IdentityAggregatedTransfer, error) {
	if operation.OperationType == types.OpType_IDENTITY_AGGREGATED_TRANSFER {
		change := new(types.IdentityAggregatedTransfer)
		return change, proto.Unmarshal(operation.Details.Value, change)
	}
	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "invalid operation type")
}

func GetIdentityAggregatedTransferContent(op *types.IdentityAggregatedTransfer) (*operation.IdentityAggregatedTransferContent, error) {
	return &operation.IdentityAggregatedTransferContent{
		Contract:      hexutil.MustDecode(op.Contract),
		GISTHash:      operation.To32Bytes(hexutil.MustDecode(op.GISTHash)),
		StateRootHash: operation.To32Bytes(hexutil.MustDecode(op.GISTHash)),
		Timestamp:     operation.To32Bytes(new(big.Int).SetUint64(op.Timestamp).Bytes()),
		Chain:         op.Chain,
	}, nil
}
