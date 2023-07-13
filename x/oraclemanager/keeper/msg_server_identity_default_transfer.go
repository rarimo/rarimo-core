package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/oraclemanager/types"
)

func (k msgServer) CreateIdentityDefaultTransferOperation(c context.Context, msg *types.MsgCreateIdentityDefaultTransferOp) (*types.MsgCreateIdentityDefaultTransferOpResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	defer k.disableFee(ctx.GasMeter().GasConsumed(), ctx.GasMeter())

	transfer, err := k.rarimo.GetIdentityDefaultTransfer(ctx, msg)
	if err != nil {
		return nil, err
	}

	oracle, ok := k.GetOracle(ctx, &types.OracleIndex{Chain: msg.Chain, Account: msg.Creator})
	if !ok {
		return nil, types.ErrOracleNotFound
	}

	if oracle.Status != types.OracleStatus_Active {
		return nil, types.ErrInactiveOracle
	}

	if err := k.rarimo.CreateIdentityDefaultTransferOperation(ctx, msg.Creator, transfer); err != nil {
		return nil, err
	}

	oracle.CreateOperationsCount = oracle.CreateOperationsCount + 1
	k.SetOracle(ctx, oracle)

	return &types.MsgCreateIdentityDefaultTransferOpResponse{}, nil
}
