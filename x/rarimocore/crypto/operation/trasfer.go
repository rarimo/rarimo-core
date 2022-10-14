package operation

import (
	"bytes"

	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/content"
)

// TransferOperation defines the token transfer operation - from one network to another with full token metadata
type TransferOperation struct {
	// Collection address on target chain
	TargetAddress []byte
	// TokenId on target chain
	TargetId []byte
	// Memory representation of amount integer as a byte array in big-endian (with leading zeros if needed)
	// Use binary.BigEndian.PutUint64(amount, c.Amount)
	Amount []byte
	// Target metadata information !!OPTIONAL!!
	TargetName     string
	TargetSymbol   string
	TargetURI      string
	ImageURI       string
	ImageHash      []byte
	TargetDecimals []byte
}

var _ Operation = &TransferOperation{}

func (t TransferOperation) GetContent() content.ContentData {
	return bytes.Join([][]byte{
		t.TargetAddress,
		[]byte(t.TargetName),
		t.TargetId,
		[]byte(t.TargetSymbol),
		t.Amount,
		[]byte(t.TargetURI),
		[]byte(t.ImageURI),
		t.ImageHash,
		t.TargetDecimals,
	}, []byte{})
}

type TransferOperationBuilder struct {
	address   string
	id        string
	amount    string
	name      string
	symbol    string
	uri       string
	imageUri  string
	imageHash string
	decimals  uint8
}

func NewTransferOperationBuilder() *TransferOperationBuilder {
	return &TransferOperationBuilder{}
}

func (b *TransferOperationBuilder) Build() *TransferOperation {
	return &TransferOperation{
		TargetAddress:  crypto.TryHexDecode(b.address),
		TargetId:       crypto.TryHexDecode(b.id),
		Amount:         amountBytes(b.amount),
		TargetName:     b.name,
		TargetSymbol:   b.symbol,
		TargetURI:      b.uri,
		ImageURI:       b.imageUri,
		ImageHash:      crypto.TryHexDecode(b.imageHash),
		TargetDecimals: decimalsBytes(b.decimals),
	}
}

func (b *TransferOperationBuilder) SetAddress(addr string) *TransferOperationBuilder {
	b.address = addr
	return b
}

func (b *TransferOperationBuilder) SetId(id string) *TransferOperationBuilder {
	b.id = id
	return b
}

func (b *TransferOperationBuilder) SetAmount(amount string) *TransferOperationBuilder {
	b.amount = amount
	return b
}

func (b *TransferOperationBuilder) SetName(name string) *TransferOperationBuilder {
	b.name = name
	return b
}

func (b *TransferOperationBuilder) SetURI(uri string) *TransferOperationBuilder {
	b.uri = uri
	return b
}

func (b *TransferOperationBuilder) SetSymbol(symbol string) *TransferOperationBuilder {
	b.symbol = symbol
	return b
}

func (b *TransferOperationBuilder) SetImageURI(uri string) *TransferOperationBuilder {
	b.imageUri = uri
	return b
}

func (b *TransferOperationBuilder) SetImageHash(hash string) *TransferOperationBuilder {
	b.imageHash = hash
	return b
}

func (b *TransferOperationBuilder) SetDecimals(d uint8) *TransferOperationBuilder {
	b.decimals = d
	return b
}
