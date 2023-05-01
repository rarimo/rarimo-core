package operation

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common/hexutil"
	eth "github.com/ethereum/go-ethereum/crypto"
	merkle "gitlab.com/rarimo/go-merkle"
)

// ContractUpgradeContent implements the Content interface provided by go-merkle and represents the content stored in the tree.
type ContractUpgradeContent struct {
	ChainName                 string
	Contract                  string
	NewImplementationContract string
	ByteCodeHash              string
	BufferAccount             string
	Nonce                     string
}

var _ merkle.Content = ContractUpgradeContent{}

func (c ContractUpgradeContent) CalculateHash() []byte {
	// Solana: HASH(contract hash, chain name, nonce, buffer address, contract address)
	// Near: HASH(contract hash, chain name, nonce, contract address)
	// EVM: HASH(new implementation address, chain name, nonce, contract address)
	return eth.Keccak256(decodeOmitemptyHex(c.ByteCodeHash), decodeOmitemptyHex(c.NewImplementationContract), []byte(c.ChainName), hexutil.MustDecode(c.Nonce), decodeOmitemptyHex(c.Contract), decodeOmitemptyHex(c.BufferAccount))
}

// Equals tests for equality of two Contents
func (c ContractUpgradeContent) Equals(other merkle.Content) bool {
	if oc, ok := other.(ContractUpgradeContent); ok {
		return bytes.Equal(oc.CalculateHash(), c.CalculateHash())
	}

	return false
}

func decodeOmitemptyHex(hex string) []byte {
	if hex == "" {
		return []byte{}
	}

	return hexutil.MustDecode(hex)
}
