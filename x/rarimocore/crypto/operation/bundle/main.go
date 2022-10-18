package bundle

import (
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/operation"
)

// Bundle stores information about bundle
type Bundle interface {
	GetBundle() operation.BundleData
}
