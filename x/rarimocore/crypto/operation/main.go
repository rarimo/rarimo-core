package operation

import (
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/content"
)

// Operation defines the certain operation that should be signed and produced by our bridge
type Operation interface {
	GetContent() content.ContentData
}
