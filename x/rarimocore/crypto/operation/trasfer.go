package operation

import (
	"bytes"

	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto"
)

// TransferOperation defines the token transfer operation - from one network to another.
// Use for EVM networks where token name, symbol and decimals are already defined by parent contract
type TransferOperation struct {
	// Collection address on target chain
	TargetAddress []byte
	// TokenId on target chain
	TargetId []byte
	// Memory representation of amount integer as a byte array in big-endian (with leading zeros if needed)
	// Use binary.BigEndian.PutUint64(amount, c.Amount)
	Amount []byte
	// Target metadata information
	TargetURI string
}

func NewTransferOperation(addHex, idHex, amount, uri string) *TransferOperation {
	return &TransferOperation{
		TargetAddress: crypto.TryHexDecode(addHex),
		TargetId:      crypto.TryHexDecode(idHex),
		Amount:        amountBytes(amount),
		TargetURI:     uri,
	}
}

var _ Operation = &TransferOperation{}

func (t TransferOperation) GetContent() crypto.ContentData {
	return bytes.Join([][]byte{t.TargetAddress, t.TargetId, []byte(t.TargetURI), t.Amount}, []byte{})
}

// TransferFullMetaOperation defines the token transfer operation - from one network to another with full token metadata
type TransferFullMetaOperation struct {
	// Collection address on target chain
	TargetAddress []byte
	// TokenId on target chain
	TargetId []byte
	// Memory representation of amount integer as a byte array in big-endian (with leading zeros if needed)
	// Use binary.BigEndian.PutUint64(amount, c.Amount)
	Amount []byte
	// Target metadata information
	TargetName     string
	TargetSymbol   string
	TargetURI      string
	TargetDecimals []byte
}

func NewTransferFullMetaOperation(addHex, idHex, amount, name, symbol, uri string, decimals uint8) *TransferFullMetaOperation {
	return &TransferFullMetaOperation{
		TargetAddress:  crypto.TryHexDecode(addHex),
		TargetId:       crypto.TryHexDecode(idHex),
		Amount:         amountBytes(amount),
		TargetName:     name,
		TargetSymbol:   symbol,
		TargetURI:      uri,
		TargetDecimals: decimalsBytes(decimals),
	}
}

var _ Operation = &TransferFullMetaOperation{}

func (t TransferFullMetaOperation) GetContent() crypto.ContentData {
	return bytes.Join([][]byte{t.TargetAddress, []byte(t.TargetName), t.TargetId, []byte(t.TargetSymbol), t.Amount, []byte(t.TargetURI), t.TargetDecimals}, []byte{})
}
