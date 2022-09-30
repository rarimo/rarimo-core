package origin

import (
	"github.com/ethereum/go-ethereum/crypto"
	xcrypto "gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto"
)

// Origin stores information about deposit (lock) transaction
type Origin interface {
	GetOrigin() xcrypto.OriginData
}

type DefaultOrigin struct {
	TxHash         string
	OperationId    string
	CurrentNetwork string
}

func NewDefaultOrigin(tx, op, network string) *DefaultOrigin {
	return &DefaultOrigin{
		TxHash:         tx,
		OperationId:    op,
		CurrentNetwork: network,
	}
}

var _ Origin = &DefaultOrigin{}

func (d DefaultOrigin) GetOrigin() (res xcrypto.OriginData) {
	h := crypto.Keccak256([]byte(d.TxHash), []byte(d.OperationId), []byte(d.CurrentNetwork))
	copy(res[:], h[:])
	return res
}
