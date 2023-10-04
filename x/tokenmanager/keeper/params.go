package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/tokenmanager/types"
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

func (k Keeper) GetNetwork(ctx sdk.Context, name string) (param types.Network, ok bool) {
	params := k.GetParams(ctx)
	for _, network := range params.Networks {
		if network.Name == name {
			return *network, true
		}
	}

	return
}

func (k Keeper) SetNetwork(ctx sdk.Context, network types.Network) {
	params := k.GetParams(ctx)

	defer func() {
		k.SetParams(ctx, params)
	}()

	for index, n := range params.Networks {
		if n.Name == network.Name {
			params.Networks[index] = &network
			return
		}
	}

	params.Networks = append(params.Networks, &network)
}

func (k Keeper) RemoveNetwork(ctx sdk.Context, chain string) {
	params := k.GetParams(ctx)

	defer func() {
		k.SetParams(ctx, params)
	}()

	networks := make([]*types.Network, 0, len(params.Networks)-1)
	for _, network := range params.Networks {
		if network.Name != chain {
			networks = append(networks, network)
		}
	}

	params.Networks = networks
}
