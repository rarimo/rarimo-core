package pkg

import (
	"cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gogo/protobuf/proto"
	"github.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
	"github.com/rarimo/rarimo-core/x/rarimocore/types"
)

func GetWorldCoinIdentityTransfer(operation types.Operation) (*types.WorldCoinIdentityTransfer, error) {
	if operation.OperationType == types.OpType_WORLDCOIN_IDENTITY_TRANSFER {
		op := new(types.WorldCoinIdentityTransfer)
		return op, proto.Unmarshal(operation.Details.Value, op)
	}
	return nil, errors.Wrap(sdkerrors.ErrInvalidType, "invalid operation type")
}

func GetWorldCoinIdentityTransferContent(op *types.WorldCoinIdentityTransfer) (*operation.WorldCoinIdentityTransferContent, error) {
	return &operation.WorldCoinIdentityTransferContent{
		Contract:  hexutil.MustDecode(op.Contract),
		PrevState: operation.To32Bytes(hexutil.MustDecode(op.PrevState)),
		State:     operation.To32Bytes(hexutil.MustDecode(op.State)),
		Timestamp: operation.To32Bytes(operation.AmountBytes(op.Timestamp)),
	}, nil
}
