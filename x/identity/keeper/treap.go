package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/identity/types"
)

const EmptyHashStr = "0x"

type Treap struct {
	Keeper
}

func (t Treap) Split(ctx sdk.Context, root, key string) (string, string) {
	if root == EmptyHashStr {
		return EmptyHashStr, EmptyHashStr
	}

	node, ok := t.GetNode(ctx, root)
	if !ok {
		return EmptyHashStr, EmptyHashStr
	}

	if less(root, key) {
		r1, r2 := t.Split(ctx, node.Right, key)
		node.Right = r1
		t.SetNode(ctx, updateNode(node))
		return root, r2
	}

	r1, r2 := t.Split(ctx, node.Left, key)
	node.Left = r2
	t.SetNode(ctx, updateNode(node))
	return r1, root
}

func (t Treap) Merge(ctx sdk.Context, r1, r2 string) string {
	if r1 == EmptyHashStr {
		return r2
	}

	if r2 == EmptyHashStr {
		return r1
	}

	node1, ok := t.GetNode(ctx, r1)
	if !ok {
		return EmptyHashStr
	}

	node2, ok := t.GetNode(ctx, r2)
	if !ok {
		return EmptyHashStr
	}

	if node1.Priority < node2.Priority {
		node2.Left = t.Merge(ctx, r1, node2.Left)
		t.SetNode(ctx, updateNode(node2))
		return r2
	}

	node1.Right = t.Merge(ctx, node1.Right, r2)
	t.SetNode(ctx, updateNode(node1))
	return r1
}

func (t Treap) Remove(ctx sdk.Context, key string) {
	root := t.GetRootKey(ctx)
	r1, r2 := t.Split(ctx, root, key)

	if r2 == key {
		node, _ := t.GetNode(ctx, r2)
		t.RemoveNode(ctx, r2)
		t.SetRootKey(ctx, t.Merge(ctx, r1, node.Right))
	}

	for {
		node, ok := t.GetNode(ctx, root)
		if !ok {
			// TODO
			return
		}

		if node.Left == key {
			node.Left = EmptyHashStr
			t.RemoveNode(ctx, key)
			t.SetNode(ctx, updateNode(node))
			break
		}

		root = node.Left
	}

	t.SetRootKey(ctx, t.Merge(ctx, r1, r2))
}

func (t Treap) Insert(ctx sdk.Context, key string, priority uint64) {
	r1, r2 := t.Split(ctx, t.GetRootKey(ctx), key)

	t.SetNode(ctx, types.Node{
		Key:          key,
		Priority:     priority,
		Left:         EmptyHashStr,
		Right:        EmptyHashStr,
		ChildrenHash: EmptyHashStr,
		Hash:         key,
	})

	r1 = t.Merge(ctx, r1, key)
	t.SetRootKey(ctx, t.Merge(ctx, r1, r2))
}

func (t Treap) MerklePath(ctx sdk.Context, key string) []string {
	current := t.GetRootKey(ctx)
	result := make([]string, 0, 64)

	for current != EmptyHashStr {
		node, ok := t.GetNode(ctx, current)
		if !ok {
			return nil
		}

		if current == key {
			result = append(result, node.ChildrenHash)
			return result
		}

		if less(key, current) {
			result = append(result, current, node.Right)
			current = node.Left
			continue
		}

		result = append(result, current, node.Left)
		current = node.Right
	}

	return nil
}

func updateNode(node types.Node) types.Node {
	node.ChildrenHash = hash(node.Left, node.Right)
	node.Hash = hash(node.ChildrenHash, node.Hash)
	return node
}
