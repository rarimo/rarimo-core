package data

import (
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/operation"
)

// Data defines the certain operation that should be signed and produced by our bridge
type Data interface {
	GetContent() operation.ContentData
}
