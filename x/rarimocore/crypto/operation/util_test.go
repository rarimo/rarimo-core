package operation

import (
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
)

func TestIntTo32Bytes(t *testing.T) {
	assert.Equal(t, "0x0000000000000000000000000000000000000000000000000000000000000013", hexutil.Encode(IntTo32Bytes(19)))
	assert.Equal(t, "0x0000000000000000000000000000000000000000000000000000000000000000", hexutil.Encode(IntTo32Bytes(0)))
}

func TestAmountBytes(t *testing.T) {
	assert.Equal(t, "0x13", hexutil.Encode(AmountBytes("19")))
	assert.Equal(t, "0x", hexutil.Encode(AmountBytes("0")))
}

func TestTo32Bytes(t *testing.T) {
	assert.Equal(t, "0x0000000000000000000000000000000000000000000000000000000000000013", hexutil.Encode(To32Bytes([]byte{0x13})))
	assert.Equal(t, "0x0000000000000000000000000000000000000000000000000000000000000000", hexutil.Encode(To32Bytes([]byte{})))
}
