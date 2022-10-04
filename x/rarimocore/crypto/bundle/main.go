package bundle

import (
	"bytes"

	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto"
)

// Bundle stores information about bundle
type Bundle interface {
	GetBundle() crypto.BundleData
}

type DefaultBundle struct {
	Salt   []byte
	Bundle []byte
}

func NewDefaultBundle(salt, bundle string) *DefaultBundle {
	return &DefaultBundle{
		Salt:   crypto.TryHexDecode(salt),
		Bundle: crypto.TryHexDecode(bundle),
	}
}

var _ Bundle = &DefaultBundle{}

func (d DefaultBundle) GetBundle() (res crypto.BundleData) {
	return bytes.Join([][]byte{d.Salt, d.Bundle}, []byte{})
}
