package pkg

import (
	"cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gogo/protobuf/proto"
	"github.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
	"github.com/rarimo/rarimo-core/x/rarimocore/types"
	"math/big"
)

func GetPassportRootUpdate(operation types.Operation) (*types.PassportRootUpdate, error) {
	if operation.OperationType == types.OpType_PASSPORT_ROOT_UPDATE {
		op := new(types.PassportRootUpdate)
		return op, proto.Unmarshal(operation.Details.Value, op)
	}
	return nil, errors.Wrap(sdkerrors.ErrInvalidType, "invalid operation type")
}

func GetPassportRootUpdateContent(op *types.PassportRootUpdate) (*operation.PassportRootUpdateContent, error) {
	return &operation.PassportRootUpdateContent{
		SourceContractAddress:      hexutil.MustDecode(op.SourceContractAddress),
		DestinationContractAddress: hexutil.MustDecode(op.DestinationContractAddress),
		Root:                       operation.To32Bytes(hexutil.MustDecode(op.Root)),
		RootTimestamp:              operation.To32Bytes(new(big.Int).SetInt64(op.Timestamp).Bytes()),
	}, nil
}
