package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/vestingmint/types"
)

func (k Keeper) EndBlocker(ctx sdk.Context) {
	k.IterateOverActiveVestingsQueue(ctx, uint64(ctx.BlockHeight()), func(vesting types.Vesting) (stop bool) {
		k.RemoveFromActiveVestingsQueue(ctx, vesting.NextDepositBlock, vesting.Index)

		address, err := sdk.AccAddressFromBech32(vesting.Account)
		if err != nil {
			return false
		}

		amount, ok := sdk.NewIntFromString(vesting.Payments[vesting.PaymentsDone])
		if !ok {
			return false
		}

		vesting.PaymentsDone += 1

		if vesting.PaymentsDone < uint64(len(vesting.Payments)) {
			vesting.NextDepositBlock = vesting.NextDepositBlock + vesting.DeltaBlocks
			k.AddToActiveVestingsQueue(ctx, vesting.NextDepositBlock, vesting.Index)
		}

		k.SetVesting(ctx, vesting)

		if err := k.bankKeeper.MintTokens(ctx, address, sdk.NewCoins(sdk.NewCoin(vesting.Denom, amount))); err != nil {
			ctx.Logger().Error("failed to mint vesting tokens: " + err.Error())
			return false
		}

		return false
	})
}
