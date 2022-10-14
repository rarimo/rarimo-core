package bundle

import (
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/content"
)

// Bundle stores information about bundle
type Bundle interface {
	GetBundle() content.BundleData
}
