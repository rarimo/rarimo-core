package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto"
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

func (k Keeper) UpdateParams(ctx sdk.Context, params types.Params) {
	params.Threshold = uint64(crypto.GetThreshold(k.getActivePartiesAmount(ctx)))
	params.IsUpdateRequired = true

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

func (k Keeper) SetPartyStatus(ctx sdk.Context, account string, status types.PartyStatus) (changed bool) {
	params := k.GetParams(ctx)

	var party *types.Party

	for _, p := range params.Parties {
		if p.Account == account {
			party = p
			break
		}
	}

	if party == nil {
		return false
	}

	party.Status = status
	k.SetParams(ctx, params)

	return true
}

func (k Keeper) IncrementPartyViolationsCounter(ctx sdk.Context, account string) (ok bool) {
	params := k.GetParams(ctx)

	var party *types.Party

	for _, p := range params.Parties {
		if p.Account == account {
			party = p
			break
		}
	}

	if party == nil {
		return false
	}

	party.ViolationsCount++
	k.SetParams(ctx, params)

	return true
}

func (k Keeper) getActivePartiesAmount(ctx sdk.Context) int {
	activePartyCount := 0
	params := k.GetParams(ctx)

	for _, party := range params.Parties {
		if party.Status == types.PartyStatus_Active {
			activePartyCount++
		}
	}

	return activePartyCount
}
