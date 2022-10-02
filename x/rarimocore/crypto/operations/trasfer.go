package operations

import (
	"bytes"

	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto"
)

// TransferOperation defines the token transfer operation - from one network to another
type TransferOperation struct {
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

func NewTransferOperation(addHex, idHex, amount, name, symbol, uri string, decimals uint8) *TransferOperation {
	return &TransferOperation{
		TargetAddress:  tryHexDecode(addHex),
		TargetId:       tryHexDecode(idHex),
		Amount:         amountBytes(amount),
		TargetName:     name,
		TargetSymbol:   symbol,
		TargetURI:      uri,
		TargetDecimals: decimalsBytes(decimals),
	}
}

var _ Operation = &TransferOperation{}

func (t TransferOperation) GetContent() crypto.ContentData {
	return bytes.Join([][]byte{t.TargetAddress, t.TargetId, t.Amount, []byte(t.TargetName), []byte(t.TargetSymbol), []byte(t.TargetURI), t.TargetDecimals}, []byte{})
}
