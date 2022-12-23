package data

import (
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
)

// Data defines the certain operation that should be signed and produced by our bridge
type Data interface {
	GetContent() operation.ContentData
}
