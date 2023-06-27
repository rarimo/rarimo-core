// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rarimocore/op_fee_token_management.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
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

type FeeTokenManagementType int32

const (
	FeeTokenManagementType_ADD_FEE_TOKEN      FeeTokenManagementType = 0
	FeeTokenManagementType_REMOVE_FEE_TOKEN   FeeTokenManagementType = 1
	FeeTokenManagementType_UPDATE_FEE_TOKEN   FeeTokenManagementType = 2
	FeeTokenManagementType_WITHDRAW_FEE_TOKEN FeeTokenManagementType = 3
)

var FeeTokenManagementType_name = map[int32]string{
	0: "ADD_FEE_TOKEN",
	1: "REMOVE_FEE_TOKEN",
	2: "UPDATE_FEE_TOKEN",
	3: "WITHDRAW_FEE_TOKEN",
}

var FeeTokenManagementType_value = map[string]int32{
	"ADD_FEE_TOKEN":      0,
	"REMOVE_FEE_TOKEN":   1,
	"UPDATE_FEE_TOKEN":   2,
	"WITHDRAW_FEE_TOKEN": 3,
}

func (x FeeTokenManagementType) String() string {
	return proto.EnumName(FeeTokenManagementType_name, int32(x))
}

func (FeeTokenManagementType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_85e0bebf8043d77d, []int{0}
}

type FeeTokenManagement struct {
	OpType   FeeTokenManagementType `protobuf:"varint,1,opt,name=opType,proto3,enum=rarimo.rarimocore.rarimocore.FeeTokenManagementType" json:"opType,omitempty"`
	Token    types.FeeToken         `protobuf:"bytes,2,opt,name=token,proto3" json:"token"`
	Chain    string                 `protobuf:"bytes,3,opt,name=chain,proto3" json:"chain,omitempty"`
	Receiver string                 `protobuf:"bytes,4,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Nonce    string                 `protobuf:"bytes,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *FeeTokenManagement) Reset()         { *m = FeeTokenManagement{} }
func (m *FeeTokenManagement) String() string { return proto.CompactTextString(m) }
func (*FeeTokenManagement) ProtoMessage()    {}
func (*FeeTokenManagement) Descriptor() ([]byte, []int) {
	return fileDescriptor_85e0bebf8043d77d, []int{0}
}
func (m *FeeTokenManagement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeeTokenManagement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeeTokenManagement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeeTokenManagement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeeTokenManagement.Merge(m, src)
}
func (m *FeeTokenManagement) XXX_Size() int {
	return m.Size()
}
func (m *FeeTokenManagement) XXX_DiscardUnknown() {
	xxx_messageInfo_FeeTokenManagement.DiscardUnknown(m)
}

var xxx_messageInfo_FeeTokenManagement proto.InternalMessageInfo

func (m *FeeTokenManagement) GetOpType() FeeTokenManagementType {
	if m != nil {
		return m.OpType
	}
	return FeeTokenManagementType_ADD_FEE_TOKEN
}

func (m *FeeTokenManagement) GetToken() types.FeeToken {
	if m != nil {
		return m.Token
	}
	return types.FeeToken{}
}

func (m *FeeTokenManagement) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *FeeTokenManagement) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *FeeTokenManagement) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func init() {
	proto.RegisterEnum("rarimo.rarimocore.rarimocore.FeeTokenManagementType", FeeTokenManagementType_name, FeeTokenManagementType_value)
	proto.RegisterType((*FeeTokenManagement)(nil), "rarimo.rarimocore.rarimocore.FeeTokenManagement")
}

func init() {
	proto.RegisterFile("rarimocore/op_fee_token_management.proto", fileDescriptor_85e0bebf8043d77d)
}

var fileDescriptor_85e0bebf8043d77d = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x33, 0xfd, 0x43, 0x47, 0x94, 0x3a, 0x94, 0x12, 0x8b, 0xc4, 0xe2, 0x2a, 0x08, 0x26,
	0x52, 0x7d, 0x81, 0x96, 0xa4, 0xf8, 0x57, 0x2b, 0x21, 0x5a, 0x70, 0x13, 0xa6, 0xe1, 0x1a, 0x83,
	0x26, 0x13, 0xa6, 0x41, 0xec, 0x5b, 0xf8, 0x58, 0x5d, 0x76, 0xe9, 0x4a, 0xa4, 0x7d, 0x01, 0x1f,
	0x41, 0x32, 0x53, 0x62, 0xc0, 0xe2, 0x2a, 0xf7, 0x9e, 0xcc, 0x77, 0xee, 0x61, 0xe6, 0x62, 0x9d,
	0x53, 0x1e, 0x46, 0xcc, 0x67, 0x1c, 0x4c, 0x96, 0x78, 0x8f, 0x00, 0x5e, 0xca, 0x9e, 0x21, 0xf6,
	0x22, 0x1a, 0xd3, 0x00, 0x22, 0x88, 0x53, 0x23, 0xe1, 0x2c, 0x65, 0x64, 0x5f, 0x9e, 0x34, 0x7e,
	0x81, 0x42, 0xd9, 0xda, 0x13, 0x94, 0x84, 0xb8, 0x99, 0x50, 0x4e, 0xa3, 0x89, 0x04, 0x5b, 0x8d,
	0x80, 0x05, 0x4c, 0x94, 0x66, 0x56, 0x49, 0xf5, 0xf0, 0x1b, 0x61, 0xd2, 0x07, 0x70, 0x33, 0x6c,
	0x90, 0xcf, 0x22, 0xd7, 0xb8, 0xc6, 0x12, 0x77, 0x9a, 0x80, 0x8a, 0xda, 0x48, 0xdf, 0xe9, 0x9c,
	0x19, 0xff, 0x8d, 0x35, 0xfe, 0x3a, 0x64, 0xac, 0xb3, 0xf2, 0x20, 0x16, 0xae, 0x8a, 0x5c, 0x6a,
	0xa9, 0x8d, 0xf4, 0xad, 0x8e, 0xbe, 0xc6, 0xac, 0x98, 0x3b, 0xb7, 0xeb, 0x55, 0x66, 0x9f, 0x07,
	0x8a, 0x23, 0x61, 0xd2, 0xc0, 0x55, 0xff, 0x89, 0x86, 0xb1, 0x5a, 0x6e, 0x23, 0x7d, 0xd3, 0x91,
	0x0d, 0x69, 0xe1, 0x0d, 0x0e, 0x3e, 0x84, 0xaf, 0xc0, 0xd5, 0x8a, 0xf8, 0x91, 0xf7, 0x19, 0x11,
	0xb3, 0xd8, 0x07, 0xb5, 0x2a, 0x09, 0xd1, 0x1c, 0x31, 0xdc, 0x5c, 0x9f, 0x97, 0xec, 0xe2, 0xed,
	0xae, 0x65, 0x79, 0x7d, 0xdb, 0xf6, 0xdc, 0xe1, 0x95, 0x7d, 0x53, 0x57, 0x48, 0x03, 0xd7, 0x1d,
	0x7b, 0x30, 0xbc, 0xb7, 0x0b, 0x2a, 0xca, 0xd4, 0xbb, 0x5b, 0xab, 0xeb, 0x16, 0xd5, 0x12, 0x69,
	0x62, 0x32, 0xba, 0x70, 0xcf, 0x2d, 0xa7, 0x3b, 0x2a, 0xe8, 0xe5, 0xde, 0xe5, 0x6c, 0xa1, 0xa1,
	0xf9, 0x42, 0x43, 0x5f, 0x0b, 0x0d, 0xbd, 0x2f, 0x35, 0x65, 0xbe, 0xd4, 0x94, 0x8f, 0xa5, 0xa6,
	0x3c, 0x9c, 0x04, 0x61, 0xfa, 0x42, 0xc7, 0x86, 0xcf, 0x22, 0x53, 0x5e, 0xc6, 0xea, 0x73, 0x2c,
	0x36, 0xe1, 0xcd, 0x2c, 0xac, 0x45, 0x3a, 0x4d, 0x60, 0x32, 0xae, 0x89, 0x67, 0x3b, 0xfd, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x4b, 0x7d, 0xae, 0x04, 0x31, 0x02, 0x00, 0x00,
}

func (m *FeeTokenManagement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeeTokenManagement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeeTokenManagement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Nonce) > 0 {
		i -= len(m.Nonce)
		copy(dAtA[i:], m.Nonce)
		i = encodeVarintOpFeeTokenManagement(dAtA, i, uint64(len(m.Nonce)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintOpFeeTokenManagement(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintOpFeeTokenManagement(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.Token.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOpFeeTokenManagement(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.OpType != 0 {
		i = encodeVarintOpFeeTokenManagement(dAtA, i, uint64(m.OpType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintOpFeeTokenManagement(dAtA []byte, offset int, v uint64) int {
	offset -= sovOpFeeTokenManagement(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FeeTokenManagement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OpType != 0 {
		n += 1 + sovOpFeeTokenManagement(uint64(m.OpType))
	}
	l = m.Token.Size()
	n += 1 + l + sovOpFeeTokenManagement(uint64(l))
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovOpFeeTokenManagement(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovOpFeeTokenManagement(uint64(l))
	}
	l = len(m.Nonce)
	if l > 0 {
		n += 1 + l + sovOpFeeTokenManagement(uint64(l))
	}
	return n
}

func sovOpFeeTokenManagement(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOpFeeTokenManagement(x uint64) (n int) {
	return sovOpFeeTokenManagement(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FeeTokenManagement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOpFeeTokenManagement
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
			return fmt.Errorf("proto: FeeTokenManagement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeeTokenManagement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OpType", wireType)
			}
			m.OpType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpFeeTokenManagement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OpType |= FeeTokenManagementType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpFeeTokenManagement
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
				return ErrInvalidLengthOpFeeTokenManagement
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOpFeeTokenManagement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Token.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpFeeTokenManagement
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
				return ErrInvalidLengthOpFeeTokenManagement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpFeeTokenManagement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpFeeTokenManagement
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
				return ErrInvalidLengthOpFeeTokenManagement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpFeeTokenManagement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpFeeTokenManagement
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
				return ErrInvalidLengthOpFeeTokenManagement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpFeeTokenManagement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonce = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOpFeeTokenManagement(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOpFeeTokenManagement
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
func skipOpFeeTokenManagement(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOpFeeTokenManagement
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
					return 0, ErrIntOverflowOpFeeTokenManagement
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
					return 0, ErrIntOverflowOpFeeTokenManagement
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
				return 0, ErrInvalidLengthOpFeeTokenManagement
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOpFeeTokenManagement
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOpFeeTokenManagement
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOpFeeTokenManagement        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOpFeeTokenManagement          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOpFeeTokenManagement = fmt.Errorf("proto: unexpected end of group")
)