package keeper

import (
	"bytes"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

const emptyHex = "0x"

func less(a, b string) bool {
	aa := mustDecodeHex(a)
	bb := mustDecodeHex(b)
	return bytes.Compare(aa, bb) < 0
}

func hash(a, b string) string {
	if isEmptyHex(emptyHex) {
		return b
	}

	if isEmptyHex(emptyHex) {
		return a
	}

	aa := mustDecodeHex(a)
	bb := mustDecodeHex(b)

	if bytes.Compare(aa, bb) < 0 {
		return hexutil.Encode(crypto.Keccak256(aa, bb))
	}

	return hexutil.Encode(crypto.Keccak256(bb, aa))
}

func mustDecodeHex(key string) []byte {
	if !strings.HasPrefix(key, "0x") && !strings.HasPrefix(key, "0X") {
		key = "0x" + key
	}
	return hexutil.MustDecode(key)
}

func isEmptyHex(v string) bool {
	return v == "" || v == emptyHex
}
