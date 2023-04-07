package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// UpdateParams updates params with check that there is at least one Inactive party.
// Caller side should carefully set params.IsUpdateRequired by itself, but also check for Inactive party exist
func (k Keeper) UpdateParams(ctx sdk.Context, params types.Params) {
	for _, party := range params.Parties {
		if party.Status == types.PartyStatus_Inactive {
			params.IsUpdateRequired = true
		}
	}

	k.SetParams(ctx, params)
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
