package crypto

import (
	"bytes"
	"crypto/ed25519"

	"github.com/cbergoon/merkletree"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func ValidateECDSAKey(hexKey string) error {
	return nil
}

func ValidateEdDSAKey(hexKey string) error {
	return nil
}

func VerifyECDSA(hexSignature string, hexHash string, targetPublicKey []byte) error {
	rootBytes, err := hexutil.Decode(hexHash)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid Merkle root hash", err)
	}

	sigBytes, err := hexutil.Decode(hexSignature)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid ECDSA signature format", err)
	}

	if !secp256k1.VerifySignature(targetPublicKey, rootBytes, sigBytes) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "signed ECDSA key does not match", err)
	}

	return nil
}

func VerifyEdDSA(hexSignature string, hexHash string, targetPublicKey []byte) error {
	rootBytes, err := hexutil.Decode(hexHash)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid Merkle root hash", err)
	}

	sigBytes, err := hexutil.Decode(hexSignature)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid ECDSA signature format", err)
	}

	if !ed25519.Verify(targetPublicKey, rootBytes, sigBytes) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "signed EdDSA key does not match", err)
	}

	return nil
}

func VerifyMerkleRoot(hashes []string, hexRoot string) error {
	list := make([]merkletree.Content, 0, len(hashes))
	for _, hash := range hashes {
		list = append(list, HashContent{Hash: hash})
	}

	t, err := merkletree.NewTree(list)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error building merkle tree", err)
	}

	root, err := hexutil.Decode(hexRoot)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid Merkle root hash", err)
	}

	if !bytes.Equal(t.MerkleRoot(), root) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "wrong Merkle root hash", err)
	}

	return nil
}
