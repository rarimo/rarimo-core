package keeper

import (
	"math"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rarimo/ldif-sdk/mt"
	"github.com/rarimo/rarimo-core/x/cscalist/types"
	"github.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
)

// Treap implements dynamic Merkle tree using treap data structure.
// Proof of concept: https://github.com/olegfomenko/crypto/tree/master/go/dynamic-merkle
type Treap struct {
	store treapStore
}

type treapStore interface {
	GetRootKey(ctx sdk.Context) string
	SetRootKey(ctx sdk.Context, key string)
	GetNode(ctx sdk.Context, node string) (types.Node, bool)
	SetNode(ctx sdk.Context, node types.Node)
	RemoveNode(ctx sdk.Context, node string)
}

// Insert adds key to the treap. Key must be hex string with 0x prefix.
func (t Treap) Insert(ctx sdk.Context, key string) {
	middle := types.Node{
		Key:          key,
		Priority:     derivePriority(key),
		Left:         emptyHex,
		Right:        emptyHex,
		Hash:         key,
		ChildrenHash: emptyHex,
	}
	t.store.SetNode(ctx, middle)

	root := t.store.GetRootKey(ctx)
	if root == emptyHex {
		t.store.SetRootKey(ctx, middle.Key)
		return
	}

	left, right := t.split(ctx, root, key)
	left = t.merge(ctx, left, key)
	t.store.SetRootKey(ctx, t.merge(ctx, left, right))
}

// Remove removes key from the treap. Key must be hex string with 0x prefix.
func (t Treap) Remove(ctx sdk.Context, key string) {
	root := t.store.GetRootKey(ctx)
	if root == emptyHex {
		return
	}

	var (
		keyBig       = new(big.Int).SetBytes(hexutil.MustDecode(key))
		keySub1Big   = new(big.Int).Sub(keyBig, big.NewInt(1))
		keySub1Bytes = operation.To32Bytes(keySub1Big.Bytes()) // padding is lost on conversion to big
		keySub1      = hexutil.Encode(keySub1Bytes)
	)

	// Split the tree by key-1 => target key in the right subtree
	left, right := t.split(ctx, root, keySub1)
	if right == emptyHex {
		return
	}

	// Split the subtree by key => target key is one left node (middle)
	middle, right := t.split(ctx, right, key)
	t.store.RemoveNode(ctx, middle)
	t.store.SetRootKey(ctx, t.merge(ctx, left, right))
}

func (t Treap) split(ctx sdk.Context, root, key string) (string, string) {
	if root == emptyHex {
		return emptyHex, emptyHex
	}

	node, ok := t.store.GetNode(ctx, root)
	if !ok {
		return emptyHex, emptyHex
	}

	// Removal implementation relies on 'root <= key'
	if less(root, key) || root == key {
		left, right := t.split(ctx, node.Right, key)
		node.Right = left
		t.updateNode(ctx, &node)
		return root, right
	}

	left, right := t.split(ctx, node.Left, key)
	node.Left = right
	t.updateNode(ctx, &node)
	return left, root
}

func (t Treap) merge(ctx sdk.Context, left, right string) string {
	if left == emptyHex {
		return right
	}

	if right == emptyHex {
		return left
	}

	node1, ok := t.store.GetNode(ctx, left)
	if !ok {
		return emptyHex
	}

	node2, ok := t.store.GetNode(ctx, right)
	if !ok {
		return emptyHex
	}

	if node1.Priority > node2.Priority {
		node1.Right = t.merge(ctx, node1.Right, right)
		t.updateNode(ctx, &node1)
		return left
	}

	node2.Left = t.merge(ctx, left, node2.Left)
	t.updateNode(ctx, &node2)
	return right
}

func (t Treap) updateNode(ctx sdk.Context, node *types.Node) {
	node.ChildrenHash = t.merkleHashNodes(ctx, node.Left, node.Right)
	node.Hash = node.Key
	if node.ChildrenHash != emptyHex {
		node.Hash = hash(node.ChildrenHash, node.Key)
	}

	t.store.SetNode(ctx, *node)
}

func (t Treap) merkleHashNodes(ctx sdk.Context, left, right string) string {
	l, okl := t.store.GetNode(ctx, left)
	r, okr := t.store.GetNode(ctx, right)

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
