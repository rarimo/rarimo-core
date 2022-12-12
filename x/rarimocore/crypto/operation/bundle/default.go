package bundle

import (
	"bytes"

	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
)

type DefaultBundle struct {
	Salt   []byte
	Bundle []byte
}

var _ Bundle = &DefaultBundle{}

func (d DefaultBundle) GetBundle() (res operation.BundleData) {
	return bytes.Join([][]byte{d.Salt, d.Bundle}, []byte{})
}

type DefaultBundleBuilder struct {
	salt   string
	bundle string
}

func NewDefaultBundleBuilder() *DefaultBundleBuilder {
	return &DefaultBundleBuilder{}
}

func (b *DefaultBundleBuilder) Build() *DefaultBundle {
	return &DefaultBundle{
		Salt:   crypto.TryHexDecode(b.salt),
		Bundle: crypto.TryHexDecode(b.bundle),
	}
}

func (b *DefaultBundleBuilder) SetSalt(salt string) *DefaultBundleBuilder {
	b.salt = salt
	return b
}

func (b *DefaultBundleBuilder) SetBundle(bundle string) *DefaultBundleBuilder {
	b.bundle = bundle
	return b
}
