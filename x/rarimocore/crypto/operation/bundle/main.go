package bundle

import (
	"github.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
)

// Bundle stores information about bundle
type Bundle interface {
	GetBundle() operation.BundleData
}
