package origin

import (
	xcrypto "gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
)

// Origin stores information about deposit (lock) transaction
type Origin interface {
	GetOrigin() xcrypto.OriginData
}
