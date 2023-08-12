package keeper

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func less(a, b string) bool {
	aa := hexutil.MustDecode(a)
	bb := hexutil.MustDecode(b)
	return bytes.Compare(aa, bb) < 0
}

func hash(a, b string) string {
	if a == EmptyHashStr {
		return b
	}

	if b == EmptyHashStr {
		return a
	}

	aa := hexutil.MustDecode(a)
	bb := hexutil.MustDecode(b)

	if bytes.Compare(aa, bb) < 0 {
		return hexutil.Encode(crypto.Keccak256(aa, bb))
	}

	return hexutil.Encode(crypto.Keccak256(bb, aa))
}
