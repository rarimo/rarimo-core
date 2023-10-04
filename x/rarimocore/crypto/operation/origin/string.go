package origin

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	xcrypto "github.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
)

type StringOrigin struct {
	Str string
}

var _ Origin = &StringOrigin{}

func (s *StringOrigin) GetOrigin() (res xcrypto.OriginData) {
	h := hexutil.MustDecode(s.Str)
	copy(res[:], h[:])
	return res
}

type StringOriginBuilder struct {
	str string
}

func NewStringOriginBuilder() *StringOriginBuilder {
	return &StringOriginBuilder{}
}

func (b *StringOriginBuilder) Build() *StringOrigin {
	return &StringOrigin{
		Str: b.str,
	}
}

func (b *StringOriginBuilder) SetString(str string) *StringOriginBuilder {
	b.str = str
	return b
}
