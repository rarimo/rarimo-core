package pkg

import (
	"cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gogo/protobuf/proto"
	"github.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
	"github.com/rarimo/rarimo-core/x/rarimocore/types"
)

func GetArbitrary(operation types.Operation) (*types.Arbitrary, error) {
	if operation.OperationType == types.OpType_ARBITRARY {
		op := new(types.Arbitrary)
		return op, proto.Unmarshal(operation.Details.Value, op)
	}
	return nil, errors.Wrap(sdkerrors.ErrInvalidType, "invalid operation type")
}

func GetArbitraryContent(op *types.Arbitrary) (*operation.ArbitraryContent, error) {
	return &operation.ArbitraryContent{
		Data: hexutil.MustDecode(op.Data),
	}, nil
}
