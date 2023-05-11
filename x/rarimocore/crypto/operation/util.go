package operation

import (
	"math/big"
)

func AmountBytes(amount string) []byte {
	am, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		return []byte{}
	}

	return am.Bytes()
}

func To32Bytes(arr []byte) []byte {
	if len(arr) >= 32 || len(arr) == 0 {
		return arr
	}

	res := make([]byte, 32-len(arr))
	return append(res, arr...)
}
