package data

import (
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
)

func TestIntTo32Bytes(t *testing.T) {
	assert.Equal(t, "0x0000000000000000000000000000000000000000000000000000000000000013", hexutil.Encode(intTo32Bytes(19)))
	assert.Equal(t, "0x0000000000000000000000000000000000000000000000000000000000000000", hexutil.Encode(intTo32Bytes(0)))
}

func TestAmountBytes(t *testing.T) {
	assert.Equal(t, "0x13", hexutil.Encode(amountBytes("19")))
	assert.Equal(t, "0x", hexutil.Encode(amountBytes("0")))
}

func TestTo32Bytes(t *testing.T) {
	assert.Equal(t, "0x0000000000000000000000000000000000000000000000000000000000000013", hexutil.Encode(to32Bytes([]byte{0x13})))
	assert.Equal(t, "0x0000000000000000000000000000000000000000000000000000000000000000", hexutil.Encode(to32Bytes([]byte{})))
}
