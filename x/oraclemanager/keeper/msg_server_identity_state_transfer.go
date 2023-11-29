package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/oraclemanager/types"
)

func (k msgServer) CreateIdentityStateTransferOperation(c context.Context, msg *types.MsgCreateIdentityStateTransferOp) (*types.MsgCreateIdentityStateTransferOpResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	defer k.disableFee(ctx.GasMeter().GasConsumed(), ctx.GasMeter())

	transfer, err := k.rarimo.GetIdentityStateTransfer(ctx, msg)
	if err != nil {
		return nil, err
	}

	// Checking oracle exists and active.
	oracle, ok := k.GetOracle(ctx, &types.OracleIndex{Chain: msg.Chain, Account: msg.Creator})
	if !ok {
		return nil, types.ErrOracleNotFound
	}

	if oracle.Status != types.OracleStatus_Active {
		return nil, types.ErrInactiveOracle
	}

	// Creating transfer operation
	if err := k.rarimo.CreateIdentityStateTransferOperation(ctx, msg.Creator, transfer); err != nil {
		return nil, err
	}

	oracle.CreateOperationsCount = oracle.CreateOperationsCount + 1
	k.SetOracle(ctx, oracle)

	return &types.MsgCreateIdentityStateTransferOpResponse{}, nil
}
