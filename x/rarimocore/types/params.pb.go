// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rarimocore/params.proto

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

type ParamsUpdateType int32

const (
	ParamsUpdateType_CHANGE_SET ParamsUpdateType = 0
	ParamsUpdateType_OTHER      ParamsUpdateType = 1
)

var ParamsUpdateType_name = map[int32]string{
	0: "CHANGE_SET",
	1: "OTHER",
}

var ParamsUpdateType_value = map[string]int32{
	"CHANGE_SET": 0,
	"OTHER":      1,
}

func (x ParamsUpdateType) String() string {
	return proto.EnumName(ParamsUpdateType_name, int32(x))
}

func (ParamsUpdateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_997c11d1d0285e72, []int{0}
}

type Party struct {
	// PublicKey in hex format
	PubKey string `protobuf:"bytes,1,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	// Server address connect to
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// Rarimo core account
	Account  string `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
	Verified bool   `protobuf:"varint,4,opt,name=verified,proto3" json:"verified,omitempty"`
}

func (m *Party) Reset()         { *m = Party{} }
func (m *Party) String() string { return proto.CompactTextString(m) }
func (*Party) ProtoMessage()    {}
func (*Party) Descriptor() ([]byte, []int) {
	return fileDescriptor_997c11d1d0285e72, []int{0}
}
func (m *Party) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Party) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Party.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Party) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Party.Merge(m, src)
}
func (m *Party) XXX_Size() int {
	return m.Size()
}
func (m *Party) XXX_DiscardUnknown() {
	xxx_messageInfo_Party.DiscardUnknown(m)
}

var xxx_messageInfo_Party proto.InternalMessageInfo

func (m *Party) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *Party) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Party) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *Party) GetVerified() bool {
	if m != nil {
		return m.Verified
	}
	return false
}

// Params defines the parameters for the module.
type Params struct {
	KeyECDSA                  string   `protobuf:"bytes,1,opt,name=keyECDSA,proto3" json:"keyECDSA,omitempty"`
	Threshold                 uint64   `protobuf:"varint,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
	Parties                   []*Party `protobuf:"bytes,3,rep,name=parties,proto3" json:"parties,omitempty"`
	IsUpdateRequired          bool     `protobuf:"varint,5,opt,name=isUpdateRequired,proto3" json:"isUpdateRequired,omitempty"`
	LastSignature             string   `protobuf:"bytes,6,opt,name=lastSignature,proto3" json:"lastSignature,omitempty"`
	AvailableResignBlockDelta uint64   `protobuf:"varint,8,opt,name=availableResignBlockDelta,proto3" json:"availableResignBlockDelta,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_997c11d1d0285e72, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetKeyECDSA() string {
	if m != nil {
		return m.KeyECDSA
	}
	return ""
}

func (m *Params) GetThreshold() uint64 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

func (m *Params) GetParties() []*Party {
	if m != nil {
		return m.Parties
	}
	return nil
}

func (m *Params) GetIsUpdateRequired() bool {
	if m != nil {
		return m.IsUpdateRequired
	}
	return false
}

func (m *Params) GetLastSignature() string {
	if m != nil {
		return m.LastSignature
	}
	return ""
}

func (m *Params) GetAvailableResignBlockDelta() uint64 {
	if m != nil {
		return m.AvailableResignBlockDelta
	}
	return 0
}

func init() {
	proto.RegisterEnum("rarimo.rarimocore.rarimocore.ParamsUpdateType", ParamsUpdateType_name, ParamsUpdateType_value)
	proto.RegisterType((*Party)(nil), "rarimo.rarimocore.rarimocore.Party")
	proto.RegisterType((*Params)(nil), "rarimo.rarimocore.rarimocore.Params")
}

func init() { proto.RegisterFile("rarimocore/params.proto", fileDescriptor_997c11d1d0285e72) }

var fileDescriptor_997c11d1d0285e72 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcd, 0x8e, 0xd3, 0x30,
	0x18, 0x4c, 0xfa, 0x93, 0xb6, 0x46, 0xa0, 0xc8, 0x07, 0x30, 0xa8, 0x8a, 0xaa, 0xc2, 0xa1, 0xaa,
	0xd4, 0x14, 0xc1, 0x15, 0x0e, 0xfd, 0x89, 0xa8, 0x40, 0x82, 0x2a, 0x2d, 0x17, 0x2e, 0xc8, 0x49,
	0x3e, 0x5a, 0xab, 0x69, 0x1d, 0x6c, 0xa7, 0x22, 0x6f, 0x01, 0x6f, 0xc5, 0xb1, 0xc7, 0x3d, 0xae,
	0xda, 0x17, 0x59, 0xad, 0x93, 0xfe, 0xac, 0x56, 0xbb, 0x27, 0x7b, 0xbe, 0xf9, 0xac, 0x99, 0xb1,
	0x06, 0xbd, 0x10, 0x54, 0xb0, 0x35, 0x0f, 0xb9, 0x80, 0x7e, 0x42, 0x05, 0x5d, 0x4b, 0x37, 0x11,
	0x5c, 0x71, 0xdc, 0xcc, 0x09, 0xf7, 0xcc, 0x5f, 0x5c, 0xdb, 0x1c, 0x55, 0xa7, 0x54, 0xa8, 0x0c,
	0x3f, 0x47, 0x56, 0x92, 0x06, 0x5f, 0x20, 0x23, 0x66, 0xcb, 0xec, 0x34, 0xfc, 0x02, 0x61, 0x82,
	0x6a, 0x34, 0x8a, 0x04, 0x48, 0x49, 0x4a, 0x9a, 0x38, 0x42, 0xcd, 0x84, 0x21, 0x4f, 0x37, 0x8a,
	0x94, 0x0b, 0x26, 0x87, 0xf8, 0x15, 0xaa, 0x6f, 0x41, 0xb0, 0x5f, 0x0c, 0x22, 0x52, 0x69, 0x99,
	0x9d, 0xba, 0x7f, 0xc2, 0xed, 0x7f, 0x25, 0x64, 0x4d, 0xb5, 0xbf, 0xdb, 0xb5, 0x15, 0x64, 0xde,
	0x68, 0x3c, 0x1b, 0x14, 0xa2, 0x27, 0x8c, 0x9b, 0xa8, 0xa1, 0x96, 0x02, 0xe4, 0x92, 0xc7, 0x91,
	0x16, 0xae, 0xf8, 0xe7, 0x01, 0xfe, 0x88, 0x6a, 0x09, 0x15, 0x8a, 0x81, 0x24, 0xe5, 0x56, 0xb9,
	0xf3, 0xe4, 0xdd, 0x6b, 0xf7, 0xb1, 0x94, 0xae, 0x8e, 0xe8, 0x1f, 0xdf, 0xe0, 0x2e, 0xb2, 0x99,
	0xfc, 0x9e, 0x44, 0x54, 0x81, 0x0f, 0xbf, 0x53, 0x26, 0x20, 0x22, 0x55, 0xed, 0xf3, 0xde, 0x1c,
	0xbf, 0x41, 0x4f, 0x63, 0x2a, 0xd5, 0x8c, 0x2d, 0x36, 0x54, 0xa5, 0x02, 0x88, 0xa5, 0x9d, 0xde,
	0x1d, 0xe2, 0x0f, 0xe8, 0x25, 0xdd, 0x52, 0x16, 0xd3, 0x20, 0x06, 0x1f, 0x24, 0x5b, 0x6c, 0x86,
	0x31, 0x0f, 0x57, 0x63, 0x88, 0x15, 0x25, 0x75, 0x6d, 0xff, 0xe1, 0x85, 0x6e, 0x0f, 0xd9, 0xf9,
	0x97, 0xe4, 0xda, 0xf3, 0x2c, 0x01, 0xfc, 0x0c, 0xa1, 0xd1, 0x64, 0xf0, 0xf5, 0x93, 0xf7, 0x73,
	0xe6, 0xcd, 0x6d, 0x03, 0x37, 0x50, 0xf5, 0xdb, 0x7c, 0xe2, 0xf9, 0xb6, 0x39, 0xfc, 0xfc, 0x7f,
	0xef, 0x98, 0xbb, 0xbd, 0x63, 0x5e, 0xef, 0x1d, 0xf3, 0xef, 0xc1, 0x31, 0x76, 0x07, 0xc7, 0xb8,
	0x3a, 0x38, 0xc6, 0x8f, 0xb7, 0x0b, 0xa6, 0x62, 0x1a, 0xb8, 0x21, 0x5f, 0xf7, 0xf3, 0xf8, 0xc5,
	0xd1, 0xd3, 0xbd, 0xf8, 0xd3, 0xbf, 0x28, 0x89, 0xca, 0x12, 0x90, 0x81, 0xa5, 0x4b, 0xf2, 0xfe,
	0x26, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xaf, 0x4a, 0x63, 0x3f, 0x02, 0x00, 0x00,
}

func (m *Party) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Party) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Party) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Verified {
		i--
		if m.Verified {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintParams(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AvailableResignBlockDelta != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.AvailableResignBlockDelta))
		i--
		dAtA[i] = 0x40
	}
	if len(m.LastSignature) > 0 {
		i -= len(m.LastSignature)
		copy(dAtA[i:], m.LastSignature)
		i = encodeVarintParams(dAtA, i, uint64(len(m.LastSignature)))
		i--
		dAtA[i] = 0x32
	}
	if m.IsUpdateRequired {
		i--
		if m.IsUpdateRequired {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.Parties) > 0 {
		for iNdEx := len(m.Parties) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Parties[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Threshold != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.Threshold))
		i--
		dAtA[i] = 0x10
	}
	if len(m.KeyECDSA) > 0 {
		i -= len(m.KeyECDSA)
		copy(dAtA[i:], m.KeyECDSA)
		i = encodeVarintParams(dAtA, i, uint64(len(m.KeyECDSA)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Party) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.Verified {
		n += 2
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.KeyECDSA)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.Threshold != 0 {
		n += 1 + sovParams(uint64(m.Threshold))
	}
	if len(m.Parties) > 0 {
		for _, e := range m.Parties {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.IsUpdateRequired {
		n += 2
	}
	l = len(m.LastSignature)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.AvailableResignBlockDelta != 0 {
		n += 1 + sovParams(uint64(m.AvailableResignBlockDelta))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Party) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Party: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Party: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Verified", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			m.Verified = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyECDSA", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyECDSA = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Threshold", wireType)
			}
			m.Threshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Threshold |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parties", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Parties = append(m.Parties, &Party{})
			if err := m.Parties[len(m.Parties)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsUpdateRequired", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			m.IsUpdateRequired = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastSignature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastSignature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AvailableResignBlockDelta", wireType)
			}
			m.AvailableResignBlockDelta = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AvailableResignBlockDelta |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
