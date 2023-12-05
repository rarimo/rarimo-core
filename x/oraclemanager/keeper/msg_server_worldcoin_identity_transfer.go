package keeper

import (
	"context"

	"github.com/rarimo/rarimo-core/x/oraclemanager/types"
)

func (k msgServer) CreateWorldCoinIdentityTransferOperation(c context.Context, msg *types.MsgCreateWorldCoinIdentityTransferOp) (*types.MsgCreateWorldCoinIdentityTransferOpResponse, error) {
	// TODO: Implement me
	return (&types.UnimplementedMsgServer{}).CreateWorldCoinIdentityTransferOperation(c, msg)
}
