package origin

import (
	"github.com/ethereum/go-ethereum/crypto"
	xcrypto "gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
)

type DefaultOrigin struct {
	TxHash         string
	OperationId    string
	CurrentNetwork string
}

var _ Origin = &DefaultOrigin{}

func (d DefaultOrigin) GetOrigin() (res xcrypto.OriginData) {
	h := crypto.Keccak256([]byte(d.TxHash), []byte(d.OperationId), []byte(d.CurrentNetwork))
	copy(res[:], h[:])
	return res
}

type DefaultOriginBuilder struct {
	tx      string
	id      string
	network string
}

func NewDefaultOriginBuilder() *DefaultOriginBuilder {
	return &DefaultOriginBuilder{}
}

func (b *DefaultOriginBuilder) Build() *DefaultOrigin {
	return &DefaultOrigin{
		TxHash:         b.tx,
		OperationId:    b.id,
		CurrentNetwork: b.network,
	}
}

func (b *DefaultOriginBuilder) SetTxHash(tx string) *DefaultOriginBuilder {
	b.tx = tx
	return b
}

func (b *DefaultOriginBuilder) SetOpId(id string) *DefaultOriginBuilder {
	b.id = id
	return b
}

func (b *DefaultOriginBuilder) SetCurrentNetwork(network string) *DefaultOriginBuilder {
	b.network = network
	return b
}
