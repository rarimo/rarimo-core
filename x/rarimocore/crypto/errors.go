package crypto

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const ErrorNamespaceCrypto = "crypto"

var (
	ErrInvalidKey        = sdkerrors.Register(ErrorNamespaceCrypto, 1200, "The public key is invalid in some way")
	ErrInvalidMerkleRoot = sdkerrors.Register(ErrorNamespaceCrypto, 1201, "The Merkle root hash is invalid in some way")
	ErrInvalidSignature  = sdkerrors.Register(ErrorNamespaceCrypto, 1202, "The signature is invalid in some way")
)
