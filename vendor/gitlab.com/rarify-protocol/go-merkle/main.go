package go_merkle

import (
	"bytes"
	"crypto/sha256"
)

type Content interface {
	CalculateHash() []byte
	Equals(other Content) bool
}

type DefaultContent struct {
	data []byte
}

var _ Content = &DefaultContent{}

func (d *DefaultContent) CalculateHash() []byte {
	return sha256.New().Sum(d.data)
}

func (d *DefaultContent) Equals(other Content) bool {
	return bytes.Equal(d.CalculateHash(), other.CalculateHash())
}

type HashF func(data ...[]byte) []byte

type node struct {
	left, right    *node
	lIndex, rIndex int
	hash           []byte
}

func (n *node) updateHash(hash HashF) {
	if n.left == nil && n.right == nil {
		return
	}

	if n.left == nil {
		n.hash = n.right.hash
		return
	}

	if n.right == nil {
		n.hash = n.left.hash
		return
	}

	var hashData []byte
	if bytes.Compare(n.left.hash, n.right.hash) >= 0 {
		hashData = append(n.left.hash, n.right.hash...)
	} else {
		hashData = append(n.right.hash, n.left.hash...)
	}

	n.hash = hash(hashData)
}

type Tree struct {
	contents     []Content
	hashFunction HashF
	root         node
}

func NewTree(hash HashF, list ...Content) *Tree {
	if len(list) == 0 {
		return nil
	}

	var nodes []node = make([]node, 0, len(list))
	for i, content := range list {
		nodes = append(nodes, node{
			nil,
			nil,
			i,
			i,
			content.CalculateHash()})
	}

	for len(nodes) > 1 {
		var newNodes = make([]node, 0, len(nodes)/2+1)
		for i := 1; i < len(nodes); i += 2 {
			newNode := node{
				left:   &nodes[i-1],
				right:  &nodes[i],
				lIndex: nodes[i-1].lIndex,
				rIndex: nodes[i].rIndex,
			}

			newNode.updateHash(hash)
			newNodes = append(newNodes, newNode)
		}

		if len(nodes)%2 == 1 {
			newNodes = append(newNodes, nodes[len(nodes)-1])
		}
		nodes = newNodes
	}

	return &Tree{
		contents:     list,
		hashFunction: hash,
		root:         nodes[0],
	}
}

func (t *Tree) Root() []byte {
	return t.root.hash
}

func (t *Tree) Path(content Content) ([][]byte, bool) {
	for i, c := range t.contents {
		if c.Equals(content) {
			var path [][]byte = make([][]byte, 0, 64)
			var currentNode = &t.root
			for !bytes.Equal(currentNode.hash, c.CalculateHash()) {
				if currentNode.left == nil {
					currentNode = currentNode.right
					continue
				}

				if currentNode.right == nil {
					currentNode = currentNode.left
					continue
				}

				if i <= currentNode.left.rIndex {
					path = append(path, currentNode.right.hash)
					currentNode = currentNode.left
				} else {
					path = append(path, currentNode.left.hash)
					currentNode = currentNode.right
				}
			}

			var result [][]byte = make([][]byte, 0, len(path))
			for i := len(path) - 1; i >= 0; i-- {
				result = append(result, path[i])
			}

			return result, true
		}
	}

	return nil, false
}
