package operation

import (
	"math/big"
)

func amountBytes(amount string) []byte {
	am, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		// it is NFT
		am = new(big.Int).SetInt64(1)
	}

	return bigInt256Bytes(am)
}

func decimalsBytes(decimals uint8) []byte {
	dm := new(big.Int).SetUint64(uint64(decimals))
	return bigInt256Bytes(dm)
}

func bigInt256Bytes(i *big.Int) []byte {
	iBytes := i.Bytes()
	result := make([]byte, 32)

	for i := range iBytes {
		result[31-i] = iBytes[len(iBytes)-1-i]
	}
	return result
}
