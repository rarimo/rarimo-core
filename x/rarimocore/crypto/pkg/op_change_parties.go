package pkg

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/operation"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

func GetChangeParties(operation types.Operation) (*types.ChangeParties, error) {
	if operation.OperationType == types.OpType_CHANGE_PARTIES {
		change := new(types.ChangeParties)
		return change, proto.Unmarshal(operation.Details.Value, change)
	}
	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "invalid operation type")
}

func GetChangePartiesContent(op *types.ChangeParties) (*operation.ChangePartiesContent, error) {
	return &operation.ChangePartiesContent{
		NewSet:    op.Parties,
		Signature: op.Signature,
	}, nil
}
