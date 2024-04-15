package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/cscalist/types"
)

func (k Keeper) SetNode(ctx sdk.Context, n types.Node) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.NodeKeyPrefix))
	b := k.cdc.MustMarshal(&n)
	store.Set([]byte(n.Key), b)
}

func (k Keeper) GetNode(ctx sdk.Context, key string) (val types.Node, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.NodeKeyPrefix))

	b := store.Get([]byte(key))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) RemoveNode(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.NodeKeyPrefix))
	store.Delete([]byte(key))
}

// RemoveTree removes all nodes iteratively from the store. It should be done
// with dropping just NodeKeyPrefix content, but requires deep research of Cosmos
// SDK store.
func (k Keeper) RemoveTree(ctx sdk.Context) {
	var (
		store = prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.NodeKeyPrefix))
		it    = sdk.KVStorePrefixIterator(store, nil)
	)

	defer func() { _ = it.Close() }()

	for ; it.Valid(); it.Next() {
		store.Delete(it.Key())
	}

	k.SetRootKey(ctx, emptyHex)
}

func (k Keeper) GetAllNode(ctx sdk.Context) []types.Node {
	const optCap = 382 // current number of certificates
	var (
		tree  = make([]types.Node, 0, optCap)
		store = prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.NodeKeyPrefix))
		it    = sdk.KVStorePrefixIterator(store, nil)
	)

	defer func() { _ = it.Close() }()

	for ; it.Valid(); it.Next() {
		var val types.Node
		k.cdc.MustUnmarshal(it.Value(), &val)
		tree = append(tree, val)
	}

	return tree
}
