package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/oraclemanager/types"
)

func (k msgServer) CreateTransferOperation(goCtx context.Context, msg *types.MsgCreateTransferOp) (*types.MsgCreateTransferOpResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	defer k.disableFee(ctx.GasMeter().GasConsumed(), ctx.GasMeter())

	// REQUIRES: validated from/to
	transfer, err := k.rarimo.GetTransfer(ctx, msg)
	if err != nil {
		return nil, err
	}

	// Checking oracle exists and active.
	oracle, ok := k.GetOracle(ctx, &types.OracleIndex{Chain: msg.From.Chain, Account: msg.Creator})
	if !ok {
		return nil, types.ErrOracleNotFound
	}

	if oracle.Status != types.OracleStatus_Active {
		return nil, types.ErrInactiveOracle
	}

	// Creating transfer operation
	if err := k.rarimo.CreateTransferOperation(ctx, msg.Creator, transfer, false); err != nil {
		return nil, err
	}

	oracle.CreateOperationsCount = oracle.CreateOperationsCount + 1
	k.SetOracle(ctx, oracle)

	return &types.MsgCreateTransferOpResponse{}, nil
}
