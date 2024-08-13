package pkg

import (
	"cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gogo/protobuf/proto"
	"github.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
	"github.com/rarimo/rarimo-core/x/rarimocore/types"
)

func GetRootUpdate(operation types.Operation) (*types.RootUpdate, error) {
	if operation.OperationType == types.OpType_UPDATE_ROOT {
		op := new(types.RootUpdate)
		return op, proto.Unmarshal(operation.Details.Value, op)
	}
	return nil, errors.Wrap(sdkerrors.ErrInvalidType, "invalid operation type")
}

func GetRootUpdateContent(op *types.RootUpdate) (*operation.RootUpdateContent, error) {
	return &operation.RootUpdateContent{
		ContractAddress: operation.To32Bytes(hexutil.MustDecode(op.ContractAddress)),
		Root:            operation.To32Bytes(hexutil.MustDecode(op.Root)),
	}, nil
}
