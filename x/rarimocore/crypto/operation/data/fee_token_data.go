package data

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
	rarimotypes "gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
	tokenmanagertypes "gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

var EmptyEVMAddress = common.Address{}

// FeeTokenData defines the fee token management operation
type FeeTokenData struct {
	OpType rarimotypes.FeeTokenManagementType
	// Fee token address
	Address []byte
	// Memory representation of amount integer as a byte array in big-endian (with leading zeros if needed)
	// Use binary.BigEndian.PutUint64(amount, c.Amount)
	Amount []byte
}

var _ Data = &FeeTokenData{}

func (f FeeTokenData) GetContent() operation.ContentData {
	return bytes.Join([][]byte{
		[]byte{byte(f.OpType)},
		f.Address,
		f.Amount,
	}, []byte{})
}

type FeeTokenDataBuilder struct {
	network tokenmanagertypes.NetworkType
	opType  rarimotypes.FeeTokenManagementType
	address []byte
	amount  []byte
}

func NewFeeTokenDataBuilder() *FeeTokenDataBuilder {
	return &FeeTokenDataBuilder{}
}

func (b *FeeTokenDataBuilder) Build() *FeeTokenData {
	return &FeeTokenData{
		OpType:  b.opType,
		Address: b.address,
		Amount:  b.amount,
	}
}

func (b *FeeTokenDataBuilder) SetNetworkType(network tokenmanagertypes.NetworkType) *FeeTokenDataBuilder {
	b.network = network
	return b
}

func (b *FeeTokenDataBuilder) SetAddress(addr string) *FeeTokenDataBuilder {
	if addr == "" {
		if b.network == tokenmanagertypes.NetworkType_EVM {
			b.address = EmptyEVMAddress[:]
		}

		return b
	}

	b.address = crypto.TryHexDecode(addr)
	return b
}

func (b *FeeTokenDataBuilder) SetAmount(amount string) *FeeTokenDataBuilder {
	b.amount = operation.To32Bytes(operation.AmountBytes(amount))
	return b
}

func (b *FeeTokenDataBuilder) SetOpType(tp rarimotypes.FeeTokenManagementType) *FeeTokenDataBuilder {
	b.opType = tp
	return b
}
