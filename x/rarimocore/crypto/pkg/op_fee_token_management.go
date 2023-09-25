package pkg

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gogo/protobuf/proto"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation/data"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
	tokentypes "gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func GetFeeTokenManagement(operation types.Operation) (*types.FeeTokenManagement, error) {
	if operation.OperationType == types.OpType_FEE_TOKEN_MANAGEMENT {
		op := new(types.FeeTokenManagement)
		return op, proto.Unmarshal(operation.Details.Value, op)
	}
	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "invalid operation type")
}

func GetFeeTokenManagementContent(params *tokentypes.FeeNetworkParams, op *types.FeeTokenManagement) (*operation.FeeTokenManagementContent, error) {
	var receiver []byte
	if op.Receiver != "" {
		receiver = hexutil.MustDecode(op.Receiver)
	}

	return &operation.FeeTokenManagementContent{
		Nonce:          operation.To32Bytes(operation.AmountBytes(op.Nonce)),
		TargetNetwork:  op.Chain,
		Receiver:       receiver,
		TargetContract: hexutil.MustDecode(params.Contract),
		Data:           data.NewFeeTokenDataBuilder().SetOpType(op.OpType).SetAddress(op.Token.Contract).SetAmount(op.Token.Amount).Build().GetContent(),
	}, nil
}
