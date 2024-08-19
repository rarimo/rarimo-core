package pkg

import (
	"cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gogo/protobuf/proto"
	"github.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
	"github.com/rarimo/rarimo-core/x/rarimocore/types"
	"time"
)

func GetPassportRootUpdate(operation types.Operation) (*types.PassportRootUpdate, error) {
	if operation.OperationType == types.OpType_PASSPORT_ROOT_UPDATE {
		op := new(types.PassportRootUpdate)
		return op, proto.Unmarshal(operation.Details.Value, op)
	}
	return nil, errors.Wrap(sdkerrors.ErrInvalidType, "invalid operation type")
}

func GetPassportRootUpdateContent(op *types.PassportRootUpdate) (*operation.PassportRootUpdateContent, error) {
	timestamp := time.Unix(int64(op.Timestamp), 0)
	timestampBytes, err := timestamp.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	
	if op.Root[:2] != "0x" {
		op.Root = "0x" + op.Root
	}

	root, err := hexutil.Decode(op.Root)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	contractAddress, err := hexutil.Decode(op.ContractAddress)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	return &operation.PassportRootUpdateContent{
		ContractAddress: operation.To32Bytes(contractAddress),
		Root:            operation.To32Bytes(root),
		RootTimestamp:   operation.To32Bytes(timestampBytes),
	}, nil
}
