package data

import (
	"bytes"

	tokentypes "github.com/rarimo/rarimo-core/x/tokenmanager/types"

	"github.com/rarimo/rarimo-core/x/rarimocore/crypto"
	"github.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
)

// TransferData defines the token transfer operation - from one network to another with full token metadata
type TransferData struct {
	// Network type to define content hash function
	networkType tokentypes.NetworkType

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
	switch t.networkType {
	case tokentypes.NetworkType_Near:
		return t.getNearContent()
	default:
		return t.getContent()
	}
}

func (t TransferData) getContent() operation.ContentData {
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

func (t TransferData) getNearContent() operation.ContentData {
	return bytes.Join([][]byte{
		operation.IntTo32Bytes(len(t.TargetAddress)),
		t.TargetAddress,
		operation.IntTo32Bytes(len([]byte(t.TargetName))),
		[]byte(t.TargetName),
		operation.IntTo32Bytes(len(t.TargetId)),
		t.TargetId,
		operation.IntTo32Bytes(len(t.Amount)),
		t.Amount,
		operation.IntTo32Bytes(len([]byte(t.ImageURI))),
		[]byte(t.ImageURI),
		operation.IntTo32Bytes(len(t.ImageHash)),
		t.ImageHash,
	}, []byte{})
}

type TransferDataBuilder struct {
	networkType tokentypes.NetworkType

	address   []byte
	id        []byte
	amount    []byte
	name      string
	symbol    string
	uri       string
	imageUri  string
	imageHash []byte
	decimals  []byte
}

func NewTransferDataBuilder() *TransferDataBuilder {
	return &TransferDataBuilder{}
}

func (b *TransferDataBuilder) Build() *TransferData {
	return &TransferData{
		networkType: b.networkType,

		TargetAddress:  b.address,
		TargetId:       b.id,
		Amount:         b.amount,
		TargetName:     b.name,
		TargetSymbol:   b.symbol,
		TargetURI:      b.uri,
		ImageURI:       b.imageUri,
		ImageHash:      b.imageHash,
		TargetDecimals: b.decimals,
	}
}

func (b *TransferDataBuilder) SetNetworkType(networkType tokentypes.NetworkType) *TransferDataBuilder {
	b.networkType = networkType
	return b
}

func (b *TransferDataBuilder) SetAddress(addr string) *TransferDataBuilder {
	b.address = crypto.TryHexDecode(addr)
	return b
}

func (b *TransferDataBuilder) SetId(id string) *TransferDataBuilder {
	b.id = operation.To32Bytes(crypto.TryHexDecode(id))
	return b
}

func (b *TransferDataBuilder) SetAmount(amount string) *TransferDataBuilder {
	b.amount = operation.To32Bytes(operation.AmountBytes(amount))
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
	b.imageHash = crypto.TryHexDecode(hash)
	return b
}

func (b *TransferDataBuilder) SetDecimals(d uint8) *TransferDataBuilder {
	b.decimals = []byte{d}
	return b
}
