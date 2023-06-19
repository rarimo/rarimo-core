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

		amount, ok := sdk.NewIntFromString(vesting.Amount)
		if !ok {
			return false
		}

		summary, ok := sdk.NewIntFromString(vesting.Paid)
		if !ok {
			return false
		}

		summary = summary.Add(amount)
		vesting.Paid = summary.String()

		if summary.LT(amount.MulRaw(int64(vesting.PaymentsCount))) {
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
