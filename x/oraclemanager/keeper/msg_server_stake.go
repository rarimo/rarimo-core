package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/oraclemanager/types"
)

func (k msgServer) Stake(goCtx context.Context, msg *types.MsgStake) (*types.MsgStakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// REQUIRES: validated index
	if _, ok := k.GetOracle(ctx, msg.Index); ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "oracle already exists")
	}

	params := k.GetParams(ctx)

	// REQUIRES: validated amount
	amount, _ := sdk.NewIntFromString(msg.Amount)
	// REQUIRES: validated amount in params
	requiredAmount, _ := sdk.NewIntFromString(params.MinOracleStake)
	if amount.LT(requiredAmount) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid amount: less then required")
	}

	moduleAccount := k.ak.GetModuleAccount(ctx, types.ModuleName)
	if moduleAccount == nil {
		panic(sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "module account %s does not exist", types.ModuleName))
	}

	// REQUIRES: validated account
	oracleOwnerAddr, _ := sdk.AccAddressFromBech32(msg.Index.Account)
	if err := k.bank.SendCoins(ctx, oracleOwnerAddr, moduleAccount.GetAddress(), sdk.NewCoins(sdk.NewCoin(params.StakeDenom, amount))); err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error withdrawing coins: %s", err.Error())
	}

	k.SetOracle(ctx, types.Oracle{
		Index:  msg.Index,
		Status: types.OracleStatus_Active,
		Stake:  msg.Amount,
	})

	return &types.MsgStakeResponse{}, nil
}
