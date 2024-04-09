package keeper

import (
	"math"
	"math/big"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rarimo/rarimo-core/x/cscalist/types"
)

// Treap implements dynamic Merkle tree using treap data structure.
// Proof of concept: https://github.com/olegfomenko/crypto/tree/master/go/dynamic-merkle
type Treap struct {
	Keeper
}

func (t Treap) Split(ctx sdk.Context, root, key string) (string, string) {
	if isEmptyHex(root) {
		return emptyHex, emptyHex
	}

	node, ok := t.GetNode(ctx, root)
	if !ok {
		return emptyHex, emptyHex
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
	if isEmptyHex(r1) {
		return r2
	}

	if isEmptyHex(r2) {
		return r1
	}

	node1, ok := t.GetNode(ctx, r1)
	if !ok {
		return emptyHex
	}

	node2, ok := t.GetNode(ctx, r2)
	if !ok {
		return emptyHex
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
	if isEmptyHex(root) {
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
			node.Left = emptyHex
			t.RemoveNode(ctx, key)
			t.updateNode(ctx, &node)
			t.SetNode(ctx, node)
			break
		}

		root = node.Left
	}

	t.SetRootKey(ctx, t.Merge(ctx, r1, r2))
}

func (t Treap) Insert(ctx sdk.Context, key string) {
	node := types.Node{
		Key:          key,
		Priority:     derivePriority(key),
		Left:         emptyHex,
		Right:        emptyHex,
		ChildrenHash: emptyHex,
		Hash:         key,
	}
	t.SetNode(ctx, node)

	root := t.GetRootKey(ctx)
	if isEmptyHex(root) {
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

	for !isEmptyHex(current) {
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
	if isEmptyHex(node.ChildrenHash) {
		node.Hash = node.Key
		return
	}

	node.Hash = hash(node.ChildrenHash, node.Key)
}

func (t Treap) merkleHashNodes(ctx sdk.Context, left, right string) string {
	l, okl := t.GetNode(ctx, left)
	r, okr := t.GetNode(ctx, right)

	if !okl && !okr {
		return emptyHex
	}

	if !okr {
		return l.Hash
	}

	if !okl {
		return r.Hash
	}

	return hash(r.Hash, l.Hash)
}

// priority = hash(key) % (2^64-1)
// function panics if key is not hex-encoded
func derivePriority(key string) uint64 {
	if !strings.HasPrefix(key, "0x") && !strings.HasPrefix(key, "0X") {
		key = "0x" + key
	}

	var (
		bytes  = mustDecodeHex(key)
		keccak = new(big.Int).SetBytes(crypto.Keccak256(bytes))
		u64    = new(big.Int).SetUint64(math.MaxUint64)
	)

	return keccak.Mod(keccak, u64).Uint64()
}
