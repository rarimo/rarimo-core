package pkg

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gogo/protobuf/proto"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/content"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

func GetChangeKey(operation types.Operation) (*types.ChangeKey, error) {
	if operation.OperationType == types.OpType_CHANGE_KEY {
		change := new(types.ChangeKey)
		return change, proto.Unmarshal(operation.Details.Value, change)
	}
	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "invalid operation type")
}

func GetChangeKeyContent(change *types.ChangeKey) (*content.ChangeKeyContent, error) {
	key, err := hexutil.Decode(change.NewKey)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid key hex string")
	}

	return &content.ChangeKeyContent{Key: key}, nil
}
