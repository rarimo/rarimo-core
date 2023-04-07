package crypto

import (
	"bytes"
	"crypto/ed25519"
	"math/big"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	merkle "gitlab.com/rarimo/go-merkle"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

const ECDSASignatureSize = 64

func VerifyECDSA(hexSignature string, hexHash string, targetPublicKey string) error {
	if targetPublicKey == "" {
		return nil
	}

	rootBytes, err := hexutil.Decode(hexHash)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidMerkleRoot, "invalid Merkle root hash %v", err)
	}

	sigBytes, err := hexutil.Decode(hexSignature)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidSignature, "invalid ECDSA signature format %v", err)
	}

	if len(sigBytes) < ECDSASignatureSize {
		return sdkerrors.Wrap(ErrInvalidSignature, "invalid ECDSA signature format")
	}

	targetKeyBytes, err := hexutil.Decode(targetPublicKey)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidKey, "invalid ECDSA target key format %v", err)
	}

	if !crypto.VerifySignature(targetKeyBytes, rootBytes, sigBytes[:ECDSASignatureSize]) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "signed ECDSA key does not match")
	}

	return nil
}

func VerifyEdDSA(hexSignature string, hexHash string, targetPublicKey string) error {
	if targetPublicKey == "" {
		return nil
	}

	rootBytes, err := hexutil.Decode(hexHash)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidMerkleRoot, "invalid Merkle root hash %v", err)
	}

	sigBytes, err := hexutil.Decode(hexSignature)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidSignature, "invalid ECDSA signature format %v", err)
	}

	targetKeyBytes, err := hexutil.Decode(targetPublicKey)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidKey, "invalid EdDSA target key format %v", err)
	}

	if !ed25519.Verify(targetKeyBytes, rootBytes, sigBytes) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "signed EdDSA key does not match")
	}

	return nil
}

func VerifyMerkleRoot(hashes []merkle.Content, hexRoot string) error {
	t := merkle.NewTree(crypto.Keccak256, hashes...)

	root, err := hexutil.Decode(hexRoot)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidMerkleRoot, "invalid Merkle root hash %v", err)
	}

	if !bytes.Equal(t.Root(), root) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "wrong Merkle root hash - required %s got %s", hexutil.Encode(t.Root()), hexRoot)
	}

	return nil
}

func TryHexDecode(hexStr string) []byte {
	resp, err := hexutil.Decode(hexStr)
	if err != nil {
		return []byte{}
	}

	return resp
}

func GetPartiesHash(set []*types.Party) []byte {
	var data []byte
	for _, p := range set {
		data = append(data, []byte(p.PubKey)...)
		data = append(data, []byte(p.Address)...)
		data = append(data, []byte(p.Account)...)
		data = append(data, big.NewInt(int64(p.Status)).Bytes()...)
	}
	return crypto.Keccak256(data)
}

func GetThreshold(n int) int {
	var res float32 = float32(n) * 2 / 3
	return int(res)
}
