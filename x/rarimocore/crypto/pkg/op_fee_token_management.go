package pkg

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gogo/protobuf/proto"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation/data"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation/origin"
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

func GetFeeTokenManagementContent(strOrigin string, params tokentypes.NetworkParams, op *types.FeeTokenManagement) (*operation.FeeTokenManagementContent, error) {
	return &operation.FeeTokenManagementContent{
		Origin:         origin.NewStringOriginBuilder().SetString(strOrigin).Build().GetOrigin(),
		TargetNetwork:  params.Name,
		Receiver:       hexutil.MustDecode(op.Receiver),
		TargetContract: hexutil.MustDecode(params.Contract),
		Data:           data.NewFeeTokenDataBuilder().SetOpType(op.OpType).SetAddress(op.Token.GetContract()).SetAmount(op.Token.Amount).Build().GetContent(),
	}, nil
}
