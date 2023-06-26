package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/oraclemanager/types"
)

func (k msgServer) Unstake(goCtx context.Context, msg *types.MsgUnstake) (*types.MsgUnstakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// REQUIRES: validated index
	oracle, ok := k.GetOracle(ctx, msg.Index)
	if !ok {
		return nil, types.ErrOracleNotFound
	}

	if oracle.Status == types.OracleStatus_Slashed || oracle.Status == types.OracleStatus_Freezed {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "oracle is slashed or frozen: can not unstake")
	}

	params := k.GetParams(ctx)
	amount := sdk.ZeroInt()

	for i, d := range oracle.Delegations {
		if d.Delegator == msg.Sender {
			amount, _ = sdk.NewIntFromString(d.Amount)
			oracle.Delegations[i].Amount = sdk.ZeroInt().String()
			break
		}
	}

	moduleAccount := k.ak.GetModuleAccount(ctx, types.ModuleName)
	if moduleAccount == nil {
		panic(sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "module account %s does not exist", types.ModuleName))
	}

	senderAddr, _ := sdk.AccAddressFromBech32(msg.Sender)
	if err := k.bank.SendCoins(ctx, moduleAccount.GetAddress(), senderAddr, sdk.NewCoins(sdk.NewCoin(params.StakeDenom, amount))); err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error withdrawing coins: %s", err.Error())
	}

	totalStake, _ := sdk.NewIntFromString(oracle.Stake)
	totalStake = totalStake.Sub(amount)
	oracle.Stake = totalStake.String()

	requiredAmount, _ := sdk.NewIntFromString(params.MinOracleStake)
	if totalStake.LT(requiredAmount) {
		oracle.Status = types.OracleStatus_Inactive
		ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeOracleDeactivated,
			sdk.NewAttribute(types.AttributeKeyChain, msg.Index.Chain),
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Index.Account),
		))
	}

	k.SetOracle(ctx, oracle)
	return &types.MsgUnstakeResponse{}, nil
}
