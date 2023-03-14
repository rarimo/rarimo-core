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

	if oracle.Status == types.OracleStatus_Slashed {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "oracle is slashed: can not unstake")
	}

	params := k.GetParams(ctx)
	amount, _ := sdk.NewIntFromString(oracle.Stake)

	moduleAccount := k.ak.GetModuleAccount(ctx, types.ModuleName)
	if moduleAccount == nil {
		panic(sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "module account %s does not exist", types.ModuleName))
	}

	oracleOwnerAddr, _ := sdk.AccAddressFromBech32(oracle.Index.Account)
	if err := k.bank.SendCoins(ctx, moduleAccount.GetAddress(), oracleOwnerAddr, sdk.NewCoins(sdk.NewCoin(params.StakeDenom, amount))); err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error withdrawing coins: %s", err.Error())
	}

	oracle.Status = types.OracleStatus_Inactive
	oracle.Stake = sdk.ZeroInt().String()
	k.SetOracle(ctx, oracle)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeOracleDeactivated,
		sdk.NewAttribute(types.AttributeKeyChain, msg.Index.Chain),
		sdk.NewAttribute(types.AttributeKeyAccount, msg.Index.Account),
	))

	return &types.MsgUnstakeResponse{}, nil
}
