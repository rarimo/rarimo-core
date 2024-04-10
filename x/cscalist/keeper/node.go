package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/cscalist/types"
)

func (k Keeper) SetNode(ctx sdk.Context, n types.Node) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeKeyPrefix))
	b := k.cdc.MustMarshal(&n)
	store.Set(types.NodeKey(n.Key), b)
}

func (k Keeper) GetNode(
	ctx sdk.Context,
	key string,
) (val types.Node, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeKeyPrefix))

	b := store.Get(types.NodeKey(key))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) RemoveNode(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeKeyPrefix))
	store.Delete(types.NodeKey(key))
}

// RemoveTree removes all nodes along with NodeKeyPrefix from the store
func (k Keeper) RemoveTree(ctx sdk.Context) {
	ctx.KVStore(k.storeKey).Delete(types.NodeKey(types.NodeKeyPrefix))
}

func (k Keeper) GetAllNode(ctx sdk.Context) []types.Node {
	const optCap = 64 // we definitely have more than 50 certificates now
	var (
		tree  = make([]types.Node, 0, optCap)
		store = prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeKeyPrefix))
		it    = sdk.KVStorePrefixIterator(store, []byte{})
	)

	defer func() { _ = it.Close() }()

	for ; it.Valid(); it.Next() {
		var val types.Node
		k.cdc.MustUnmarshal(it.Value(), &val)
		tree = append(tree, val)
	}

	return tree
}
