package data

import (
	"bytes"

	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/operation"
)

// TransferData defines the token transfer operation - from one network to another with full token metadata
type TransferData struct {
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

var _ Data = &TransferData{}

func (t TransferData) GetContent() operation.ContentData {
	return bytes.Join([][]byte{
		t.TargetAddress,
		[]byte(t.TargetName),
		t.TargetId,
		[]byte(t.TargetURI),
		t.Amount,
		[]byte(t.ImageURI),
		t.ImageHash,
		[]byte(t.TargetSymbol),
		t.TargetDecimals,
	}, []byte{})
}

type TransferDataBuilder struct {
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

func NewTransferDataBuilder() *TransferDataBuilder {
	return &TransferDataBuilder{}
}

func (b *TransferDataBuilder) Build() *TransferData {
	return &TransferData{
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

func (b *TransferDataBuilder) SetAddress(addr string) *TransferDataBuilder {
	b.address = addr
	return b
}

func (b *TransferDataBuilder) SetId(id string) *TransferDataBuilder {
	b.id = id
	return b
}

func (b *TransferDataBuilder) SetAmount(amount string) *TransferDataBuilder {
	b.amount = amount
	return b
}

func (b *TransferDataBuilder) SetName(name string) *TransferDataBuilder {
	b.name = name
	return b
}

func (b *TransferDataBuilder) SetURI(uri string) *TransferDataBuilder {
	b.uri = uri
	return b
}

func (b *TransferDataBuilder) SetSymbol(symbol string) *TransferDataBuilder {
	b.symbol = symbol
	return b
}

func (b *TransferDataBuilder) SetImageURI(uri string) *TransferDataBuilder {
	b.imageUri = uri
	return b
}

func (b *TransferDataBuilder) SetImageHash(hash string) *TransferDataBuilder {
	b.imageHash = hash
	return b
}

func (b *TransferDataBuilder) SetDecimals(d uint8) *TransferDataBuilder {
	b.decimals = d
	return b
}
