package data

import (
	"math/big"
)

func amountBytes(amount string) []byte {
	am, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		return []byte{}
	}

	return am.Bytes()
}

func to32Bytes(arr []byte) []byte {
	if len(arr) >= 32 || len(arr) == 0 {
		return arr
	}

	res := make([]byte, 32-len(arr))
	return append(res, arr...)
}

func intTo32Bytes(amount int) []byte {
	return to32Bytes(big.NewInt(int64(amount)).Bytes())
}
