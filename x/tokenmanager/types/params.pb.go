// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tokenmanager/params.proto

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

type NetworkType int32

const (
	NetworkType_EVM    NetworkType = 0
	NetworkType_Solana NetworkType = 1
	NetworkType_Near   NetworkType = 2
	NetworkType_Other  NetworkType = 3
)

var NetworkType_name = map[int32]string{
	0: "EVM",
	1: "Solana",
	2: "Near",
	3: "Other",
}

var NetworkType_value = map[string]int32{
	"EVM":    0,
	"Solana": 1,
	"Near":   2,
	"Other":  3,
}

func (x NetworkType) String() string {
	return proto.EnumName(NetworkType_name, int32(x))
}

func (NetworkType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_55f297a3c8945a13, []int{0}
}

type NetworkTypeBinding struct {
	CoreType  uint32 `protobuf:"varint,1,opt,name=coreType,proto3" json:"coreType,omitempty"`
	SaverType uint32 `protobuf:"varint,2,opt,name=saverType,proto3" json:"saverType,omitempty"`
}

func (m *NetworkTypeBinding) Reset()         { *m = NetworkTypeBinding{} }
func (m *NetworkTypeBinding) String() string { return proto.CompactTextString(m) }
func (*NetworkTypeBinding) ProtoMessage()    {}
func (*NetworkTypeBinding) Descriptor() ([]byte, []int) {
	return fileDescriptor_55f297a3c8945a13, []int{0}
}
func (m *NetworkTypeBinding) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NetworkTypeBinding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NetworkTypeBinding.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NetworkTypeBinding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkTypeBinding.Merge(m, src)
}
func (m *NetworkTypeBinding) XXX_Size() int {
	return m.Size()
}
func (m *NetworkTypeBinding) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkTypeBinding.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkTypeBinding proto.InternalMessageInfo

func (m *NetworkTypeBinding) GetCoreType() uint32 {
	if m != nil {
		return m.CoreType
	}
	return 0
}

func (m *NetworkTypeBinding) GetSaverType() uint32 {
	if m != nil {
		return m.SaverType
	}
	return 0
}

type NetworkParams struct {
	// network name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// contract address hex
	Contract string                `protobuf:"bytes,2,opt,name=contract,proto3" json:"contract,omitempty"`
	Types    []*NetworkTypeBinding `protobuf:"bytes,3,rep,name=types,proto3" json:"types,omitempty"`
	Type     NetworkType           `protobuf:"varint,4,opt,name=type,proto3,enum=rarimo.rarimocore.tokenmanager.NetworkType" json:"type,omitempty"`
}

func (m *NetworkParams) Reset()         { *m = NetworkParams{} }
func (m *NetworkParams) String() string { return proto.CompactTextString(m) }
func (*NetworkParams) ProtoMessage()    {}
func (*NetworkParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_55f297a3c8945a13, []int{1}
}
func (m *NetworkParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NetworkParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NetworkParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NetworkParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkParams.Merge(m, src)
}
func (m *NetworkParams) XXX_Size() int {
	return m.Size()
}
func (m *NetworkParams) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkParams.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkParams proto.InternalMessageInfo

func (m *NetworkParams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkParams) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *NetworkParams) GetTypes() []*NetworkTypeBinding {
	if m != nil {
		return m.Types
	}
	return nil
}

func (m *NetworkParams) GetType() NetworkType {
	if m != nil {
		return m.Type
	}
	return NetworkType_EVM
}

// Params defines the parameters for the module.
type Params struct {
	Networks []*NetworkParams `protobuf:"bytes,1,rep,name=networks,proto3" json:"networks,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_55f297a3c8945a13, []int{2}
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

func (m *Params) GetNetworks() []*NetworkParams {
	if m != nil {
		return m.Networks
	}
	return nil
}

func init() {
	proto.RegisterEnum("rarimo.rarimocore.tokenmanager.NetworkType", NetworkType_name, NetworkType_value)
	proto.RegisterType((*NetworkTypeBinding)(nil), "rarimo.rarimocore.tokenmanager.NetworkTypeBinding")
	proto.RegisterType((*NetworkParams)(nil), "rarimo.rarimocore.tokenmanager.NetworkParams")
	proto.RegisterType((*Params)(nil), "rarimo.rarimocore.tokenmanager.Params")
}

func init() { proto.RegisterFile("tokenmanager/params.proto", fileDescriptor_55f297a3c8945a13) }

var fileDescriptor_55f297a3c8945a13 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0xc6, 0x33, 0x4d, 0xda, 0xdb, 0x9c, 0xd2, 0x4b, 0x98, 0x8d, 0xb1, 0xc8, 0x50, 0xba, 0x2a,
	0x4a, 0x13, 0x88, 0x0b, 0x97, 0x42, 0x41, 0x50, 0xd0, 0x2a, 0xa9, 0xb8, 0x70, 0x37, 0xad, 0x63,
	0x0d, 0x6d, 0x66, 0xc2, 0x64, 0x50, 0xfb, 0x16, 0x3e, 0x96, 0xb8, 0xea, 0xd2, 0xa5, 0xb4, 0x2f,
	0x22, 0x99, 0x29, 0xb1, 0x45, 0x10, 0x5d, 0x9d, 0x7f, 0xf9, 0x7e, 0xf9, 0xce, 0x1c, 0xd8, 0x55,
	0x62, 0xca, 0x78, 0x4a, 0x39, 0x9d, 0x30, 0x19, 0x66, 0x54, 0xd2, 0x34, 0x0f, 0x32, 0x29, 0x94,
	0xc0, 0x44, 0x52, 0x99, 0xa4, 0x22, 0x30, 0x61, 0x2c, 0x24, 0x0b, 0x36, 0x3f, 0x6e, 0xed, 0x6c,
	0x49, 0x13, 0x7e, 0x2f, 0x8c, 0xb0, 0x33, 0x00, 0x3c, 0x60, 0xea, 0x49, 0xc8, 0xe9, 0xf5, 0x3c,
	0x63, 0xfd, 0x84, 0xdf, 0x25, 0x7c, 0x82, 0x5b, 0x50, 0x2f, 0x18, 0x45, 0xcb, 0x47, 0x6d, 0xd4,
	0x6d, 0xc6, 0x65, 0x8d, 0xf7, 0xc0, 0xcd, 0xe9, 0x23, 0x93, 0x7a, 0x58, 0xd1, 0xc3, 0xaf, 0x46,
	0xe7, 0x0d, 0x41, 0x73, 0x0d, 0xbc, 0xd2, 0x06, 0x31, 0x06, 0x87, 0xd3, 0xd4, 0x70, 0xdc, 0x58,
	0xe7, 0x86, 0xcf, 0x95, 0xa4, 0x63, 0xa5, 0x11, 0x6e, 0x5c, 0xd6, 0xf8, 0x14, 0xaa, 0x6a, 0x9e,
	0xb1, 0xdc, 0xb7, 0xdb, 0x76, 0xb7, 0x11, 0x45, 0xc1, 0xcf, 0xab, 0x05, 0xdf, 0xed, 0xc7, 0x06,
	0x80, 0x8f, 0xc1, 0x29, 0x12, 0xdf, 0x69, 0xa3, 0xee, 0xff, 0xe8, 0xe0, 0x0f, 0xa0, 0x58, 0x0b,
	0x3b, 0x43, 0xa8, 0xad, 0x97, 0x38, 0x83, 0x3a, 0x37, 0xe3, 0xdc, 0x47, 0xda, 0x57, 0xef, 0x97,
	0x38, 0x03, 0x88, 0x4b, 0xf9, 0xfe, 0x11, 0x34, 0x36, 0xfe, 0x84, 0xff, 0x81, 0x7d, 0x72, 0x73,
	0xe1, 0x59, 0x18, 0xa0, 0x36, 0x14, 0x33, 0xca, 0xa9, 0x87, 0x70, 0x1d, 0x9c, 0x01, 0xa3, 0xd2,
	0xab, 0x60, 0x17, 0xaa, 0x97, 0xea, 0x81, 0x49, 0xcf, 0xee, 0x9f, 0xbf, 0x2e, 0x09, 0x5a, 0x2c,
	0x09, 0xfa, 0x58, 0x12, 0xf4, 0xb2, 0x22, 0xd6, 0x62, 0x45, 0xac, 0xf7, 0x15, 0xb1, 0x6e, 0xa3,
	0x49, 0xa2, 0x66, 0x74, 0x14, 0x8c, 0x45, 0x1a, 0x1a, 0x3b, 0xeb, 0xd0, 0x2b, 0x6c, 0x85, 0xcf,
	0xe1, 0xd6, 0xf5, 0xf5, 0xe3, 0x8c, 0x6a, 0xfa, 0xfe, 0x87, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x5d, 0x64, 0x35, 0x12, 0x55, 0x02, 0x00, 0x00,
}

func (m *NetworkTypeBinding) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkTypeBinding) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NetworkTypeBinding) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SaverType != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SaverType))
		i--
		dAtA[i] = 0x10
	}
	if m.CoreType != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.CoreType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *NetworkParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NetworkParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Types) > 0 {
		for iNdEx := len(m.Types) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Types[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Name)))
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
	if len(m.Networks) > 0 {
		for iNdEx := len(m.Networks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Networks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
func (m *NetworkTypeBinding) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CoreType != 0 {
		n += 1 + sovParams(uint64(m.CoreType))
	}
	if m.SaverType != 0 {
		n += 1 + sovParams(uint64(m.SaverType))
	}
	return n
}

func (m *NetworkParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if len(m.Types) > 0 {
		for _, e := range m.Types {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.Type != 0 {
		n += 1 + sovParams(uint64(m.Type))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Networks) > 0 {
		for _, e := range m.Networks {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NetworkTypeBinding) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: NetworkTypeBinding: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkTypeBinding: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoreType", wireType)
			}
			m.CoreType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CoreType |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SaverType", wireType)
			}
			m.SaverType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SaverType |= uint32(b&0x7F) << shift
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
func (m *NetworkParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: NetworkParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
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
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Types", wireType)
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
			m.Types = append(m.Types, &NetworkTypeBinding{})
			if err := m.Types[len(m.Types)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= NetworkType(b&0x7F) << shift
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
				return fmt.Errorf("proto: wrong wireType = %d for field Networks", wireType)
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
			m.Networks = append(m.Networks, &NetworkParams{})
			if err := m.Networks[len(m.Networks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
