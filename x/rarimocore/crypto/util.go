package crypto

import (
	"bytes"
	"crypto/ed25519"

	"github.com/cbergoon/merkletree"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

const ECDSAPublicKeySize = 33

func ValidateECDSAKey(hexKey string) error {
	if hexKey == "" {
		// Key has not been initialized
		return nil
	}

	keyBytes, err := hexutil.Decode(hexKey)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidKey, "invalid ECDSA key format", err)
	}

	if len(keyBytes) != ECDSAPublicKeySize {
		return ErrInvalidKey
	}

	return nil
}

func ValidateEdDSAKey(hexKey string) error {
	if hexKey == "" {
		// Key has not been initialized
		return nil
	}

	keyBytes, err := hexutil.Decode(hexKey)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidKey, "invalid ECDSA key format", err)
	}

	if len(keyBytes) != ed25519.PublicKeySize {
		return ErrInvalidKey
	}

	return nil
}

func VerifyECDSA(hexSignature string, hexHash string, targetPublicKey string) error {
	if targetPublicKey == "" {
		return nil
	}

	rootBytes, err := hexutil.Decode(hexHash)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidMerkleRoot, "invalid Merkle root hash", err)
	}

	sigBytes, err := hexutil.Decode(hexSignature)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidSignature, "invalid ECDSA signature format", err)
	}

	targetKeyBytes, err := hexutil.Decode(targetPublicKey)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidKey, "invalid ECDSA target key format", err)
	}

	if !secp256k1.VerifySignature(targetKeyBytes, rootBytes, sigBytes) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "signed ECDSA key does not match", err)
	}

	return nil
}

func VerifyEdDSA(hexSignature string, hexHash string, targetPublicKey string) error {
	if targetPublicKey == "" {
		return nil
	}

	rootBytes, err := hexutil.Decode(hexHash)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidMerkleRoot, "invalid Merkle root hash", err)
	}

	sigBytes, err := hexutil.Decode(hexSignature)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidSignature, "invalid ECDSA signature format", err)
	}

	targetKeyBytes, err := hexutil.Decode(targetPublicKey)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidKey, "invalid EdDSA target key format", err)
	}

	if !ed25519.Verify(targetKeyBytes, rootBytes, sigBytes) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "signed EdDSA key does not match", err)
	}

	return nil
}

func VerifyMerkleRoot(hashes []merkletree.Content, hexRoot string) error {
	t, err := merkletree.NewTree(hashes)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidMerkleRoot, "error building merkle tree", err)
	}

	root, err := hexutil.Decode(hexRoot)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidMerkleRoot, "invalid Merkle root hash", err)
	}

	if !bytes.Equal(t.MerkleRoot(), root) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "wrong Merkle root hash", err)
	}

	return nil
}
