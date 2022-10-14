package origin

import (
	xcrypto "gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/operation"
)

// Origin stores information about deposit (lock) transaction
type Origin interface {
	GetOrigin() xcrypto.OriginData
}
