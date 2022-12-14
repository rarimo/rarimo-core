// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tokenmanager/item.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Type int32

const (
	Type_NATIVE       Type = 0
	Type_ERC20        Type = 1
	Type_ERC721       Type = 2
	Type_ERC1155      Type = 3
	Type_METAPLEX_NFT Type = 4
	Type_METAPLEX_FT  Type = 5
	Type_NEAR_FT      Type = 6
	Type_NEAR_NFT     Type = 7
)

var Type_name = map[int32]string{
	0: "NATIVE",
	1: "ERC20",
	2: "ERC721",
	3: "ERC1155",
	4: "METAPLEX_NFT",
	5: "METAPLEX_FT",
	6: "NEAR_FT",
	7: "NEAR_NFT",
}

var Type_value = map[string]int32{
	"NATIVE":       0,
	"ERC20":        1,
	"ERC721":       2,
	"ERC1155":      3,
	"METAPLEX_NFT": 4,
	"METAPLEX_FT":  5,
	"NEAR_FT":      6,
	"NEAR_NFT":     7,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e83fe9c71c21d068, []int{0}
}

type Item struct {
	// hex-encoded
	TokenAddress string `protobuf:"bytes,1,opt,name=tokenAddress,proto3" json:"tokenAddress,omitempty"`
	// hex-encoded
	TokenId  string `protobuf:"bytes,2,opt,name=tokenId,proto3" json:"tokenId,omitempty"`
	Chain    string `protobuf:"bytes,3,opt,name=chain,proto3" json:"chain,omitempty"`
	Index    string `protobuf:"bytes,4,opt,name=index,proto3" json:"index,omitempty"`
	Name     string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Symbol   string `protobuf:"bytes,6,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Uri      string `protobuf:"bytes,7,opt,name=uri,proto3" json:"uri,omitempty"`
	Decimals uint32 `protobuf:"varint,8,opt,name=decimals,proto3" json:"decimals,omitempty"`
	// Seed for deriving address for Solana networks. Encoded into hex string. (optional)
	Seed     string `protobuf:"bytes,9,opt,name=seed,proto3" json:"seed,omitempty"`
	ImageUri string `protobuf:"bytes,10,opt,name=imageUri,proto3" json:"imageUri,omitempty"`
	// Hash of the token image. Encoded into hex string. (optional)
	ImageHash string `protobuf:"bytes,11,opt,name=imageHash,proto3" json:"imageHash,omitempty"`
	TokenType Type   `protobuf:"varint,12,opt,name=tokenType,proto3,enum=rarimo.rarimocore.tokenmanager.Type" json:"tokenType,omitempty"`
	Wrapped   bool   `protobuf:"varint,13,opt,name=wrapped,proto3" json:"wrapped,omitempty"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_e83fe9c71c21d068, []int{0}
}
func (m *Item) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Item.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return m.Size()
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetTokenAddress() string {
	if m != nil {
		return m.TokenAddress
	}
	return ""
}

func (m *Item) GetTokenId() string {
	if m != nil {
		return m.TokenId
	}
	return ""
}

func (m *Item) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *Item) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Item) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *Item) GetDecimals() uint32 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func (m *Item) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

func (m *Item) GetImageUri() string {
	if m != nil {
		return m.ImageUri
	}
	return ""
}

func (m *Item) GetImageHash() string {
	if m != nil {
		return m.ImageHash
	}
	return ""
}

func (m *Item) GetTokenType() Type {
	if m != nil {
		return m.TokenType
	}
	return Type_NATIVE
}

func (m *Item) GetWrapped() bool {
	if m != nil {
		return m.Wrapped
	}
	return false
}

func init() {
	proto.RegisterEnum("rarimo.rarimocore.tokenmanager.Type", Type_name, Type_value)
	proto.RegisterType((*Item)(nil), "rarimo.rarimocore.tokenmanager.Item")
}

func init() { proto.RegisterFile("tokenmanager/item.proto", fileDescriptor_e83fe9c71c21d068) }

var fileDescriptor_e83fe9c71c21d068 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x3b, 0x6d, 0x9a, 0x36, 0xaf, 0x5d, 0x0d, 0x0f, 0xd1, 0x41, 0x24, 0x94, 0x3d, 0x48,
	0x11, 0x4d, 0x6d, 0x65, 0xf1, 0x5c, 0x4b, 0x16, 0x2b, 0x6b, 0x91, 0x10, 0x45, 0xbc, 0xc8, 0x34,
	0x19, 0xbb, 0x83, 0x9d, 0x24, 0x4c, 0x22, 0x6e, 0xbf, 0x85, 0x1f, 0xcb, 0xe3, 0x1e, 0x3d, 0x4a,
	0xfb, 0x3d, 0x44, 0xe6, 0x75, 0xb7, 0xeb, 0xde, 0x3c, 0xe5, 0xfd, 0x7e, 0xef, 0x3f, 0x43, 0xde,
	0x63, 0xe0, 0x41, 0x5d, 0x7c, 0x95, 0xb9, 0x16, 0xb9, 0x58, 0x49, 0x33, 0x52, 0xb5, 0xd4, 0x61,
	0x69, 0x8a, 0xba, 0xc0, 0xc7, 0x46, 0x18, 0xf5, 0x65, 0x43, 0x90, 0x16, 0xeb, 0xd0, 0xa2, 0x2e,
	0xd2, 0xc2, 0xc8, 0xf0, 0xdf, 0x23, 0xc7, 0x7f, 0x9a, 0xe0, 0xcc, 0x6b, 0xa9, 0xf1, 0x18, 0xfa,
	0xd4, 0x98, 0x66, 0x99, 0x91, 0x55, 0xc5, 0xd9, 0x80, 0x0d, 0xbd, 0xf8, 0x96, 0x43, 0x0e, 0x1d,
	0xe2, 0x79, 0xc6, 0x9b, 0xd4, 0xbe, 0x46, 0xbc, 0x07, 0xed, 0xf4, 0x5c, 0xa8, 0x9c, 0xb7, 0xc8,
	0xef, 0xc1, 0x5a, 0x95, 0x67, 0xf2, 0x82, 0x3b, 0x7b, 0x4b, 0x80, 0x08, 0x4e, 0x2e, 0xb4, 0xe4,
	0x6d, 0x92, 0x54, 0xe3, 0x7d, 0x70, 0xab, 0x8d, 0x5e, 0x16, 0x6b, 0xee, 0x92, 0xbd, 0x22, 0xf4,
	0xa1, 0xf5, 0xcd, 0x28, 0xde, 0x21, 0x69, 0x4b, 0x7c, 0x08, 0xdd, 0x4c, 0xa6, 0x4a, 0x8b, 0x75,
	0xc5, 0xbb, 0x03, 0x36, 0x3c, 0x8a, 0x0f, 0x6c, 0x6f, 0xae, 0xa4, 0xcc, 0xb8, 0xb7, 0xbf, 0xd9,
	0xd6, 0x36, 0xaf, 0xb4, 0x58, 0xc9, 0xf7, 0x46, 0x71, 0x20, 0x7f, 0x60, 0x7c, 0x04, 0x1e, 0xd5,
	0xaf, 0x45, 0x75, 0xce, 0x7b, 0xd4, 0xbc, 0x11, 0xf8, 0x06, 0x3c, 0x1a, 0x2f, 0xd9, 0x94, 0x92,
	0xf7, 0x07, 0x6c, 0x78, 0x67, 0xf2, 0x34, 0xfc, 0xbf, 0xb5, 0x86, 0xf5, 0xa6, 0x94, 0xf1, 0xcd,
	0x71, 0xbb, 0xb9, 0xef, 0x46, 0x94, 0xa5, 0xcc, 0xf8, 0xd1, 0x80, 0x0d, 0xbb, 0xf1, 0x35, 0x3e,
	0xa9, 0xc1, 0xb1, 0x61, 0x04, 0x70, 0x17, 0xd3, 0x64, 0xfe, 0x21, 0xf2, 0x1b, 0xe8, 0x41, 0x3b,
	0x8a, 0x67, 0x93, 0xe7, 0x3e, 0xb3, 0x3a, 0x8a, 0x67, 0x2f, 0x27, 0x63, 0xbf, 0x89, 0x3d, 0xe8,
	0x44, 0xf1, 0x6c, 0x3c, 0x3e, 0x39, 0xf1, 0x5b, 0xe8, 0x43, 0xff, 0x6d, 0x94, 0x4c, 0xdf, 0x9d,
	0x45, 0x1f, 0x3f, 0x2f, 0x4e, 0x13, 0xdf, 0xc1, 0xbb, 0xd0, 0x3b, 0x98, 0xd3, 0xc4, 0x6f, 0xdb,
	0xfc, 0x22, 0x9a, 0xc6, 0x16, 0x5c, 0xec, 0x43, 0x97, 0xc0, 0x66, 0x3b, 0xaf, 0xce, 0x7e, 0x6e,
	0x03, 0x76, 0xb9, 0x0d, 0xd8, 0xef, 0x6d, 0xc0, 0x7e, 0xec, 0x82, 0xc6, 0xe5, 0x2e, 0x68, 0xfc,
	0xda, 0x05, 0x8d, 0x4f, 0x93, 0x95, 0xaa, 0xd7, 0x62, 0x19, 0xa6, 0x85, 0x1e, 0xed, 0xa7, 0xbb,
	0xfa, 0x3c, 0xb3, 0x53, 0x8e, 0x2e, 0x46, 0xb7, 0x5e, 0x9c, 0xfd, 0xf5, 0x6a, 0xe9, 0xd2, 0x3e,
	0x5e, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x04, 0x85, 0xfa, 0x8e, 0x02, 0x00, 0x00,
}

func (m *Item) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Item) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Item) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Wrapped {
		i--
		if m.Wrapped {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x68
	}
	if m.TokenType != 0 {
		i = encodeVarintItem(dAtA, i, uint64(m.TokenType))
		i--
		dAtA[i] = 0x60
	}
	if len(m.ImageHash) > 0 {
		i -= len(m.ImageHash)
		copy(dAtA[i:], m.ImageHash)
		i = encodeVarintItem(dAtA, i, uint64(len(m.ImageHash)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.ImageUri) > 0 {
		i -= len(m.ImageUri)
		copy(dAtA[i:], m.ImageUri)
		i = encodeVarintItem(dAtA, i, uint64(len(m.ImageUri)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Seed) > 0 {
		i -= len(m.Seed)
		copy(dAtA[i:], m.Seed)
		i = encodeVarintItem(dAtA, i, uint64(len(m.Seed)))
		i--
		dAtA[i] = 0x4a
	}
	if m.Decimals != 0 {
		i = encodeVarintItem(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x40
	}
	if len(m.Uri) > 0 {
		i -= len(m.Uri)
		copy(dAtA[i:], m.Uri)
		i = encodeVarintItem(dAtA, i, uint64(len(m.Uri)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintItem(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintItem(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintItem(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintItem(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TokenId) > 0 {
		i -= len(m.TokenId)
		copy(dAtA[i:], m.TokenId)
		i = encodeVarintItem(dAtA, i, uint64(len(m.TokenId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TokenAddress) > 0 {
		i -= len(m.TokenAddress)
		copy(dAtA[i:], m.TokenAddress)
		i = encodeVarintItem(dAtA, i, uint64(len(m.TokenAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintItem(dAtA []byte, offset int, v uint64) int {
	offset -= sovItem(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Item) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TokenAddress)
	if l > 0 {
		n += 1 + l + sovItem(uint64(l))
	}
	l = len(m.TokenId)
	if l > 0 {
		n += 1 + l + sovItem(uint64(l))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovItem(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovItem(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovItem(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovItem(uint64(l))
	}
	l = len(m.Uri)
	if l > 0 {
		n += 1 + l + sovItem(uint64(l))
	}
	if m.Decimals != 0 {
		n += 1 + sovItem(uint64(m.Decimals))
	}
	l = len(m.Seed)
	if l > 0 {
		n += 1 + l + sovItem(uint64(l))
	}
	l = len(m.ImageUri)
	if l > 0 {
		n += 1 + l + sovItem(uint64(l))
	}
	l = len(m.ImageHash)
	if l > 0 {
		n += 1 + l + sovItem(uint64(l))
	}
	if m.TokenType != 0 {
		n += 1 + sovItem(uint64(m.TokenType))
	}
	if m.Wrapped {
		n += 2
	}
	return n
}

func sovItem(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozItem(x uint64) (n int) {
	return sovItem(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Item) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowItem
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Item: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Item: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimals |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seed = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImageUri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImageUri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImageHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImageHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenType", wireType)
			}
			m.TokenType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TokenType |= Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Wrapped", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Wrapped = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipItem(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthItem
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipItem(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowItem
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowItem
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowItem
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthItem
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupItem
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthItem
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthItem        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowItem          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupItem = fmt.Errorf("proto: unexpected end of group")
)
