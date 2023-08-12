package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/identity/types"
)

const EmptyHashStr = "0x"

// Treap implements dynamic Merkle tree.
// Proof of concept: https://github.com/olegfomenko/go-treap-merkle
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
		t.updateNode(ctx, &node)
		t.SetNode(ctx, node)
		return root, r2
	}

	r1, r2 := t.Split(ctx, node.Left, key)
	node.Left = r2
	t.updateNode(ctx, &node)
	t.SetNode(ctx, node)
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

	if node1.Priority > node2.Priority {
		node1.Right = t.Merge(ctx, node1.Right, r2)
		t.updateNode(ctx, &node1)
		t.SetNode(ctx, node1)
		return r1
	}

	node2.Left = t.Merge(ctx, r1, node2.Left)
	t.updateNode(ctx, &node2)
	t.SetNode(ctx, node2)
	return r2
}

func (t Treap) Remove(ctx sdk.Context, key string) {
	root := t.GetRootKey(ctx)
	if root == EmptyHashStr {
		return
	}

	r1, r2 := t.Split(ctx, root, key)

	if r2 == key {
		node, _ := t.GetNode(ctx, r2)
		t.RemoveNode(ctx, r2)
		t.SetRootKey(ctx, t.Merge(ctx, r1, node.Right))
		return
	}

	root = r2
	for {
		node, ok := t.GetNode(ctx, root)
		if !ok {
			return
		}

		if node.Left == key {
			node.Left = EmptyHashStr
			t.RemoveNode(ctx, key)
			t.updateNode(ctx, &node)
			t.SetNode(ctx, node)
			break
		}

		root = node.Left
	}

	t.SetRootKey(ctx, t.Merge(ctx, r1, r2))
}

func (t Treap) Insert(ctx sdk.Context, key string, priority uint64) {
	node := types.Node{
		Key:          key,
		Priority:     priority,
		Left:         EmptyHashStr,
		Right:        EmptyHashStr,
		ChildrenHash: EmptyHashStr,
		Hash:         key,
	}
	t.SetNode(ctx, node)

	root := t.GetRootKey(ctx)
	if root == EmptyHashStr {
		t.SetRootKey(ctx, node.Key)
		return
	}

	r1, r2 := t.Split(ctx, root, key)
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

		if less(current, key) {
			result = append(result, current)
			left, ok := t.GetNode(ctx, node.Left)
			if ok {
				result = append(result, left.Hash)
			}

			current = node.Right
			continue
		}

		result = append(result, current)
		right, ok := t.GetNode(ctx, node.Right)
		if ok {
			result = append(result, right.Hash)
		}

		current = node.Left
	}

	return nil
}

func (t Treap) updateNode(ctx sdk.Context, node *types.Node) {
	node.ChildrenHash = t.merkleHashNodes(ctx, node.Left, node.Right)
	if node.ChildrenHash == EmptyHashStr {
		node.Hash = node.Key
		return
	}

	node.Hash = hash(node.ChildrenHash, node.Key)
}

func (t Treap) merkleHashNodes(ctx sdk.Context, left, right string) string {
	l, okl := t.GetNode(ctx, left)
	r, okr := t.GetNode(ctx, right)

	if !okl && !okr {
		return EmptyHashStr
	}

	if !okr {
		return l.Hash
	}

	if !okl {
		return r.Hash
	}

	return hash(r.Hash, l.Hash)
}
