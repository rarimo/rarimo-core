package operations

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func tryHexDecode(hexStr string) []byte {
	resp, err := hexutil.Decode(hexStr)
	if err != nil {
		return []byte{}
	}

	return resp
}

func amountBytes(amount string) []byte {
	am, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		// it is NFT
		am = new(big.Int).SetInt64(1)
	}

	amBytes := am.Bytes()
	result := make([]byte, 32)

	for i := range amBytes {
		result[31-i] = amBytes[len(amBytes)-1-i]
	}
	return result
}
