package crypto

import (
	"crypto/sha256"

	"github.com/cbergoon/merkletree"
)

// HashContent implements the Content interface provided by merkletree and represents the content stored in the tree.
type HashContent struct {
	Hash string
}

func (c HashContent) CalculateHash() ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(c.Hash)); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

//Equals tests for equality of two Contents
func (c HashContent) Equals(other merkletree.Content) (bool, error) {
	if oc, ok := other.(HashContent); ok {
		return c.Hash == oc.Hash, nil
	}
	return false, nil
}
