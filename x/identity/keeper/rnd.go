package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

type LCG struct {
	Keeper
}

func (l LCG) Next(ctx sdk.Context) uint64 {
	params := l.GetParams(ctx)
	params.LcgValue = (params.LcgA*params.LcgValue + params.LcgB) % params.LcgMod
	l.SetParams(ctx, params)
	return params.LcgValue
}
