package keeper

import (
	"math"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rarimo/ldif-sdk/mt"
	"github.com/rarimo/rarimo-core/x/cscalist/types"
)

// Treap implements dynamic Merkle tree using treap data structure.
// Proof of concept: https://github.com/olegfomenko/crypto/tree/master/go/dynamic-merkle
type Treap struct {
	k Keeper
}

// Insert adds key to the treap. Key must be hex string with 0x prefix.
func (t Treap) Insert(ctx sdk.Context, key string) {
	node := types.Node{
		Key:          key,
		Priority:     derivePriority(key),
		Left:         emptyHex,
		Right:        emptyHex,
		Hash:         key,
		ChildrenHash: emptyHex,
	}
	t.k.SetNode(ctx, node)

	root := t.k.GetRootKey(ctx)
	if root != emptyHex {
		t.k.SetRootKey(ctx, node.Key)
		return
	}

	r1, r2 := t.split(ctx, root, key)
	r1 = t.merge(ctx, r1, key)
	t.k.SetRootKey(ctx, t.merge(ctx, r1, r2))
}

// Remove removes key from the treap. Key must be hex string with 0x prefix.
func (t Treap) Remove(ctx sdk.Context, key string) {
	root := t.k.GetRootKey(ctx)
	if root == emptyHex {
		return
	}

	r1, r2 := t.split(ctx, root, key)

	if r2 == key {
		node, _ := t.k.GetNode(ctx, r2)
		t.k.RemoveNode(ctx, r2)
		t.k.SetRootKey(ctx, t.merge(ctx, r1, node.Right))
		return
	}

	root = r2
	for {
		node, ok := t.k.GetNode(ctx, root)
		if !ok {
			return
		}

		if node.Left == key {
			node.Left = emptyHex
			t.k.RemoveNode(ctx, key)
			t.updateNode(ctx, &node)
			t.k.SetNode(ctx, node)
			break
		}

		root = node.Left
	}

	t.k.SetRootKey(ctx, t.merge(ctx, r1, r2))
}

// MerklePath provides a Merkle path with Node.MerkleHash sibling values.
// Key must be hex string with 0x prefix.
func (t Treap) MerklePath(ctx sdk.Context, key string) []string {
	current := t.k.GetRootKey(ctx)
	result := make([]string, 0, 64)

	for current != emptyHex {
		node, ok := t.k.GetNode(ctx, current)
		if !ok {
			return nil
		}

		if current == key {
			result = append(result, node.ChildrenHash)
			return result
		}

		if less(current, key) {
			result = append(result, current)
			left, ok := t.k.GetNode(ctx, node.Left)
			if ok {
				result = append(result, left.Hash)
			}

			current = node.Right
			continue
		}

		result = append(result, current)
		right, ok := t.k.GetNode(ctx, node.Right)
		if ok {
			result = append(result, right.Hash)
		}

		current = node.Left
	}

	return nil
}

func (t Treap) split(ctx sdk.Context, root, key string) (string, string) {
	if root == emptyHex {
		return emptyHex, emptyHex
	}

	node, ok := t.k.GetNode(ctx, root)
	if !ok {
		return emptyHex, emptyHex
	}

	if less(root, key) {
		r1, r2 := t.split(ctx, node.Right, key)
		node.Right = r1
		t.updateNode(ctx, &node)
		t.k.SetNode(ctx, node)
		return root, r2
	}

	r1, r2 := t.split(ctx, node.Left, key)
	node.Left = r2
	t.updateNode(ctx, &node)
	t.k.SetNode(ctx, node)
	return r1, root
}

func (t Treap) merge(ctx sdk.Context, r1, r2 string) string {
	if r1 == emptyHex {
		return r2
	}

	if r2 == emptyHex {
		return r1
	}

	node1, ok := t.k.GetNode(ctx, r1)
	if !ok {
		return emptyHex
	}

	node2, ok := t.k.GetNode(ctx, r2)
	if !ok {
		return emptyHex
	}

	if node1.Priority > node2.Priority {
		node1.Right = t.merge(ctx, node1.Right, r2)
		t.updateNode(ctx, &node1)
		t.k.SetNode(ctx, node1)
		return r1
	}

	node2.Left = t.merge(ctx, r1, node2.Left)
	t.updateNode(ctx, &node2)
	t.k.SetNode(ctx, node2)
	return r2
}

func (t Treap) updateNode(ctx sdk.Context, node *types.Node) {
	node.ChildrenHash = t.merkleHashNodes(ctx, node.Left, node.Right)
	if node.ChildrenHash == emptyHex {
		node.Hash = node.Key
		return
	}

	node.Hash = hash(node.ChildrenHash, node.Key)
}

func (t Treap) merkleHashNodes(ctx sdk.Context, left, right string) string {
	l, okl := t.k.GetNode(ctx, left)
	r, okr := t.k.GetNode(ctx, right)

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

// priority = poseidon(key) % (2^64-1)
func derivePriority(key string) uint64 {
	var (
		hex     = hexutil.MustDecode(key)
		keyHash = new(big.Int).SetBytes(mt.MustPoseidon(hex).Bytes())
		u64     = new(big.Int).SetUint64(math.MaxUint64)
	)

	return keyHash.Mod(keyHash, u64).Uint64()
}
