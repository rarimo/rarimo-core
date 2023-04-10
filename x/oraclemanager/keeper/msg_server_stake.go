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
	oracle, ok := k.GetOracle(ctx, msg.Index)
	if ok && oracle.Status != types.OracleStatus_Inactive && oracle.Status != types.OracleStatus_Active {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "can not stake for oracle: invalid status in existing entry")
	}

	existingStake := sdk.ZeroInt()
	if ok {
		existingStake, _ = sdk.NewIntFromString(oracle.Stake)
	}

	params := k.GetParams(ctx)

	// REQUIRES: validated amount
	amount, _ := sdk.NewIntFromString(msg.Amount)

	moduleAccount := k.ak.GetModuleAccount(ctx, types.ModuleName)
	if moduleAccount == nil {
		panic(sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "module account %s does not exist", types.ModuleName))
	}

	// REQUIRES: validated account
	oracleOwnerAddr, _ := sdk.AccAddressFromBech32(msg.Index.Account)
	if err := k.bank.SendCoins(ctx, oracleOwnerAddr, moduleAccount.GetAddress(), sdk.NewCoins(sdk.NewCoin(params.StakeDenom, amount))); err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error withdrawing coins: %s", err.Error())
	}

	totalStake := existingStake.Add(amount)
	resultStatus := types.OracleStatus_Active

	// REQUIRES: validated amount in params
	requiredAmount, _ := sdk.NewIntFromString(params.MinOracleStake)
	if totalStake.LT(requiredAmount) {
		resultStatus = types.OracleStatus_Inactive
	}

	k.SetOracle(ctx, types.Oracle{
		Index:  msg.Index,
		Status: resultStatus,
		Stake:  totalStake.String(),
	})

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeOracleActivated,
		sdk.NewAttribute(types.AttributeKeyChain, msg.Index.Chain),
		sdk.NewAttribute(types.AttributeKeyAccount, msg.Index.Account),
	))

	return &types.MsgStakeResponse{}, nil
}
