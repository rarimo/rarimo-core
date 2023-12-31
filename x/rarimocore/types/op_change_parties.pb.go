// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rarimocore/op_change_parties.proto

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

type ChangeParties struct {
	Parties      []*Party `protobuf:"bytes,1,rep,name=parties,proto3" json:"parties,omitempty"`
	NewPublicKey string   `protobuf:"bytes,2,opt,name=newPublicKey,proto3" json:"newPublicKey,omitempty"`
	Signature    string   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *ChangeParties) Reset()         { *m = ChangeParties{} }
func (m *ChangeParties) String() string { return proto.CompactTextString(m) }
func (*ChangeParties) ProtoMessage()    {}
func (*ChangeParties) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcce331765093f26, []int{0}
}
func (m *ChangeParties) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChangeParties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChangeParties.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChangeParties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeParties.Merge(m, src)
}
func (m *ChangeParties) XXX_Size() int {
	return m.Size()
}
func (m *ChangeParties) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeParties.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeParties proto.InternalMessageInfo

func (m *ChangeParties) GetParties() []*Party {
	if m != nil {
		return m.Parties
	}
	return nil
}

func (m *ChangeParties) GetNewPublicKey() string {
	if m != nil {
		return m.NewPublicKey
	}
	return ""
}

func (m *ChangeParties) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func init() {
	proto.RegisterType((*ChangeParties)(nil), "rarimo.rarimocore.rarimocore.ChangeParties")
}

func init() {
	proto.RegisterFile("rarimocore/op_change_parties.proto", fileDescriptor_bcce331765093f26)
}

var fileDescriptor_bcce331765093f26 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2a, 0x4a, 0x2c, 0xca,
	0xcc, 0xcd, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0xcf, 0x2f, 0x88, 0x4f, 0xce, 0x48, 0xcc, 0x4b, 0x4f,
	0x8d, 0x2f, 0x48, 0x2c, 0x2a, 0xc9, 0x4c, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92,
	0x81, 0xa8, 0xd1, 0x43, 0x28, 0x45, 0x62, 0x4a, 0x89, 0x23, 0x99, 0x50, 0x90, 0x58, 0x94, 0x98,
	0x0b, 0xd5, 0xa6, 0x34, 0x81, 0x91, 0x8b, 0xd7, 0x19, 0x6c, 0x5e, 0x00, 0xc4, 0x38, 0x21, 0x5b,
	0x2e, 0x76, 0xa8, 0xc9, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0xca, 0x7a, 0xf8, 0x8c, 0xd6,
	0x03, 0xe9, 0xab, 0x0c, 0x82, 0xe9, 0x11, 0x52, 0xe2, 0xe2, 0xc9, 0x4b, 0x2d, 0x0f, 0x28, 0x4d,
	0xca, 0xc9, 0x4c, 0xf6, 0x4e, 0xad, 0x94, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x11, 0x13,
	0x92, 0xe1, 0xe2, 0x2c, 0xce, 0x4c, 0xcf, 0x4b, 0x2c, 0x29, 0x2d, 0x4a, 0x95, 0x60, 0x06, 0x2b,
	0x40, 0x08, 0x38, 0x79, 0x9d, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72,
	0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x41,
	0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xc4, 0x05, 0x50, 0x4a, 0x17,
	0xec, 0xb1, 0x0a, 0x7d, 0x24, 0x5f, 0x96, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x7d, 0x69,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x5c, 0xe6, 0xc4, 0x42, 0x01, 0x00, 0x00,
}

func (m *ChangeParties) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChangeParties) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChangeParties) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintOpChangeParties(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NewPublicKey) > 0 {
		i -= len(m.NewPublicKey)
		copy(dAtA[i:], m.NewPublicKey)
		i = encodeVarintOpChangeParties(dAtA, i, uint64(len(m.NewPublicKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Parties) > 0 {
		for iNdEx := len(m.Parties) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Parties[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOpChangeParties(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintOpChangeParties(dAtA []byte, offset int, v uint64) int {
	offset -= sovOpChangeParties(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChangeParties) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Parties) > 0 {
		for _, e := range m.Parties {
			l = e.Size()
			n += 1 + l + sovOpChangeParties(uint64(l))
		}
	}
	l = len(m.NewPublicKey)
	if l > 0 {
		n += 1 + l + sovOpChangeParties(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovOpChangeParties(uint64(l))
	}
	return n
}

func sovOpChangeParties(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOpChangeParties(x uint64) (n int) {
	return sovOpChangeParties(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChangeParties) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOpChangeParties
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
			return fmt.Errorf("proto: ChangeParties: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChangeParties: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parties", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpChangeParties
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
				return ErrInvalidLengthOpChangeParties
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOpChangeParties
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Parties = append(m.Parties, &Party{})
			if err := m.Parties[len(m.Parties)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewPublicKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpChangeParties
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
				return ErrInvalidLengthOpChangeParties
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpChangeParties
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewPublicKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpChangeParties
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
				return ErrInvalidLengthOpChangeParties
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpChangeParties
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOpChangeParties(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOpChangeParties
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
func skipOpChangeParties(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOpChangeParties
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
					return 0, ErrIntOverflowOpChangeParties
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
					return 0, ErrIntOverflowOpChangeParties
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
				return 0, ErrInvalidLengthOpChangeParties
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOpChangeParties
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOpChangeParties
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOpChangeParties        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOpChangeParties          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOpChangeParties = fmt.Errorf("proto: unexpected end of group")
)
