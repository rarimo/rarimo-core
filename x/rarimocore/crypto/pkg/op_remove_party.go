package pkg

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/operation"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

func GetRemoveParty(operation types.Operation) (*types.RemoveParty, error) {
	if operation.OperationType == types.OpType_REMOVE_PARTY {
		change := new(types.RemoveParty)
		return change, proto.Unmarshal(operation.Details.Value, change)
	}
	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "invalid operation type")
}

func GetRemovePartyContent(op *types.RemoveParty) (*operation.RemovePartyContent, error) {
	return &operation.RemovePartyContent{
		Set:   op.CurrentSet,
		Index: op.PartyIndex,
	}, nil
}
