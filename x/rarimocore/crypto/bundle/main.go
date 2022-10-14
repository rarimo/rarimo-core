package bundle

import (
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto"
)

// Bundle stores information about bundle
type Bundle interface {
	GetBundle() crypto.BundleData
}
