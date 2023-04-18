package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	store := ctx.KVStore(k.storeKey)

	b := store.Get(types.KeyPrefix(types.ParamsKey))
	if b == nil {
		return types.DefaultParams()
	}

	k.cdc.MustUnmarshal(b, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	if err := params.Validate(); err != nil {
		panic("failed to set params: " + err.Error())
	}

	b := k.cdc.MustMarshal(&params)
	ctx.KVStore(k.storeKey).Set(types.KeyPrefix(types.ParamsKey), b)
}

func (k Keeper) GetKeyECDSA(ctx sdk.Context) string {
	params := k.GetParams(ctx)
	return params.GetKeyECDSA()
}

func (k Keeper) UpdateLastSignature(ctx sdk.Context, sig string) {
	params := k.GetParams(ctx)
	params.LastSignature = sig
	k.SetParams(ctx, params)
}

func (k Keeper) getActivePartiesAmount(ctx sdk.Context) int {
	params := k.GetParams(ctx)
	return getActivePartiesAmount(params.Parties)
}

func getActivePartiesAmount(parties []*types.Party) int {
	activePartyCount := 0

	for _, party := range parties {
		if party.Status == types.PartyStatus_Active {
			activePartyCount++
		}
	}

	return activePartyCount
}
