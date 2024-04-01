package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

type LCG struct {
	Keeper
}

// Next function returns the next pseudo-random value depending on generator state in module params.
// Such simple solution is claimed to be enough for balancing binary tree nodes.
// Linear congruential generator
// More about: https://en.wikipedia.org/wiki/Linear_congruential_generator
func (l LCG) Next(ctx sdk.Context) uint64 {
	params := l.GetParams(ctx)
	params.LcgValue = (params.LcgA*params.LcgValue + params.LcgB) % params.LcgMod
	l.SetParams(ctx, params)
	return params.LcgValue
}
