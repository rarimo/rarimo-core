package keeper

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rarimo/ldif-sdk/mt"
)

const emptyHex = "0x"

func less(a, b string) bool {
	aa := hexutil.MustDecode(a)
	bb := hexutil.MustDecode(b)
	return bytes.Compare(aa, bb) < 0
}

func hash(a, b string) string {
	if a == emptyHex {
		return b
	}

	if b == emptyHex {
		return a
	}

	aa := hexutil.MustDecode(a)
	bb := hexutil.MustDecode(b)

	if bytes.Compare(aa, bb) < 0 {
		return hexutil.Encode(mt.MustPoseidon(aa, bb))
	}

	return hexutil.Encode(mt.MustPoseidon(bb, aa))
}
