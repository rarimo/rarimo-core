package pkg

import (
	"math/big"

	"cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gogo/protobuf/proto"
	"github.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
	"github.com/rarimo/rarimo-core/x/rarimocore/types"
)

func GetCSCARootUpdate(operation types.Operation) (*types.CSCARootUpdate, error) {
	if operation.OperationType == types.OpType_CSCA_ROOT_UPDATE {
		op := new(types.CSCARootUpdate)
		return op, proto.Unmarshal(operation.Details.Value, op)
	}
	return nil, errors.Wrap(sdkerrors.ErrInvalidType, "invalid operation type")
}

func GetCSCARootUpdateContent(op *types.CSCARootUpdate) (*operation.CSCARootUpdateContent, error) {
	return &operation.CSCARootUpdateContent{
		Root:      operation.To32Bytes(hexutil.MustDecode(op.Root)),
		Timestamp: operation.To32Bytes(new(big.Int).SetUint64(op.Timestamp).Bytes()),
	}, nil
}
