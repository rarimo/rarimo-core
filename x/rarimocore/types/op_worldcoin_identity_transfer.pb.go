// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rarimocore/op_worldcoin_identity_transfer.proto

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

type WorldCoinIdentityTransfer struct {
	Contract string `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	Chain    string `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
	// Hex 0x uint256
	PrevState string `protobuf:"bytes,3,opt,name=prevState,proto3" json:"prevState,omitempty"`
	// Hex 0x uint256
	State string `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	// Dec uint256
	Timestamp string `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// simplifies voting flow
	BlockNumber uint64 `protobuf:"varint,6,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
}

func (m *WorldCoinIdentityTransfer) Reset()         { *m = WorldCoinIdentityTransfer{} }
func (m *WorldCoinIdentityTransfer) String() string { return proto.CompactTextString(m) }
func (*WorldCoinIdentityTransfer) ProtoMessage()    {}
func (*WorldCoinIdentityTransfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_a897c477db0590c9, []int{0}
}
func (m *WorldCoinIdentityTransfer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WorldCoinIdentityTransfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WorldCoinIdentityTransfer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WorldCoinIdentityTransfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorldCoinIdentityTransfer.Merge(m, src)
}
func (m *WorldCoinIdentityTransfer) XXX_Size() int {
	return m.Size()
}
func (m *WorldCoinIdentityTransfer) XXX_DiscardUnknown() {
	xxx_messageInfo_WorldCoinIdentityTransfer.DiscardUnknown(m)
}

var xxx_messageInfo_WorldCoinIdentityTransfer proto.InternalMessageInfo

func (m *WorldCoinIdentityTransfer) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *WorldCoinIdentityTransfer) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *WorldCoinIdentityTransfer) GetPrevState() string {
	if m != nil {
		return m.PrevState
	}
	return ""
}

func (m *WorldCoinIdentityTransfer) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *WorldCoinIdentityTransfer) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *WorldCoinIdentityTransfer) GetBlockNumber() uint64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*WorldCoinIdentityTransfer)(nil), "rarimo.rarimocore.rarimocore.WorldCoinIdentityTransfer")
}

func init() {
	proto.RegisterFile("rarimocore/op_worldcoin_identity_transfer.proto", fileDescriptor_a897c477db0590c9)
}

var fileDescriptor_a897c477db0590c9 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x63, 0x68, 0x2b, 0x6a, 0xb6, 0x88, 0x21, 0xa0, 0xca, 0x8a, 0x98, 0xba, 0x90, 0x22,
	0xf1, 0x0f, 0x60, 0x82, 0x81, 0xa1, 0x20, 0x21, 0xb1, 0x44, 0x8e, 0x6b, 0xa8, 0x45, 0xed, 0xb3,
	0x2e, 0x57, 0xa0, 0xff, 0x82, 0xbf, 0xc4, 0xc6, 0xd8, 0x91, 0x11, 0x25, 0x7f, 0x04, 0xc5, 0x89,
	0xda, 0x4c, 0x77, 0xef, 0x3d, 0x7d, 0x6f, 0x78, 0x7c, 0x86, 0x12, 0x8d, 0x05, 0x05, 0xa8, 0x67,
	0xe0, 0xf3, 0x0f, 0xc0, 0xd5, 0x42, 0x81, 0x71, 0xb9, 0x59, 0x68, 0x47, 0x86, 0x36, 0x39, 0xa1,
	0x74, 0xe5, 0x8b, 0xc6, 0xcc, 0x23, 0x10, 0xc4, 0x93, 0x16, 0xc8, 0xf6, 0x5c, 0xef, 0x3d, 0xff,
	0x66, 0xfc, 0xf4, 0xa9, 0xe9, 0xb8, 0x01, 0xe3, 0x6e, 0xbb, 0x8a, 0xc7, 0xae, 0x21, 0x3e, 0xe3,
	0x47, 0x0a, 0x1c, 0xa1, 0x54, 0x94, 0xb0, 0x94, 0x4d, 0xc7, 0xf3, 0x9d, 0x8e, 0x4f, 0xf8, 0x50,
	0x2d, 0xa5, 0x71, 0xc9, 0x41, 0x08, 0x5a, 0x11, 0x4f, 0xf8, 0xd8, 0xa3, 0x7e, 0x7f, 0x20, 0x49,
	0x3a, 0x39, 0x0c, 0xc9, 0xde, 0x68, 0x98, 0x32, 0x24, 0x83, 0x96, 0x09, 0xa2, 0x61, 0xc8, 0x58,
	0x5d, 0x92, 0xb4, 0x3e, 0x19, 0xb6, 0xcc, 0xce, 0x88, 0x53, 0x7e, 0x5c, 0xac, 0x40, 0xbd, 0xdd,
	0xaf, 0x6d, 0xa1, 0x31, 0x19, 0xa5, 0x6c, 0x3a, 0x98, 0xf7, 0xad, 0xeb, 0xbb, 0x9f, 0x4a, 0xb0,
	0x6d, 0x25, 0xd8, 0x5f, 0x25, 0xd8, 0x57, 0x2d, 0xa2, 0x6d, 0x2d, 0xa2, 0xdf, 0x5a, 0x44, 0xcf,
	0x97, 0xaf, 0x86, 0x96, 0xeb, 0x22, 0x53, 0x60, 0xbb, 0xdd, 0xba, 0x73, 0x11, 0xf6, 0xfb, 0xec,
	0x8f, 0x49, 0x1b, 0xaf, 0xcb, 0x62, 0x14, 0x46, 0xbb, 0xfa, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x5d,
	0x0f, 0x47, 0x3a, 0x67, 0x01, 0x00, 0x00,
}

func (m *WorldCoinIdentityTransfer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WorldCoinIdentityTransfer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WorldCoinIdentityTransfer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockNumber != 0 {
		i = encodeVarintOpWorldcoinIdentityTransfer(dAtA, i, uint64(m.BlockNumber))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Timestamp) > 0 {
		i -= len(m.Timestamp)
		copy(dAtA[i:], m.Timestamp)
		i = encodeVarintOpWorldcoinIdentityTransfer(dAtA, i, uint64(len(m.Timestamp)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintOpWorldcoinIdentityTransfer(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PrevState) > 0 {
		i -= len(m.PrevState)
		copy(dAtA[i:], m.PrevState)
		i = encodeVarintOpWorldcoinIdentityTransfer(dAtA, i, uint64(len(m.PrevState)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintOpWorldcoinIdentityTransfer(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintOpWorldcoinIdentityTransfer(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOpWorldcoinIdentityTransfer(dAtA []byte, offset int, v uint64) int {
	offset -= sovOpWorldcoinIdentityTransfer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WorldCoinIdentityTransfer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovOpWorldcoinIdentityTransfer(uint64(l))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovOpWorldcoinIdentityTransfer(uint64(l))
	}
	l = len(m.PrevState)
	if l > 0 {
		n += 1 + l + sovOpWorldcoinIdentityTransfer(uint64(l))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovOpWorldcoinIdentityTransfer(uint64(l))
	}
	l = len(m.Timestamp)
	if l > 0 {
		n += 1 + l + sovOpWorldcoinIdentityTransfer(uint64(l))
	}
	if m.BlockNumber != 0 {
		n += 1 + sovOpWorldcoinIdentityTransfer(uint64(m.BlockNumber))
	}
	return n
}

func sovOpWorldcoinIdentityTransfer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOpWorldcoinIdentityTransfer(x uint64) (n int) {
	return sovOpWorldcoinIdentityTransfer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WorldCoinIdentityTransfer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOpWorldcoinIdentityTransfer
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
			return fmt.Errorf("proto: WorldCoinIdentityTransfer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WorldCoinIdentityTransfer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpWorldcoinIdentityTransfer
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
				return ErrInvalidLengthOpWorldcoinIdentityTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpWorldcoinIdentityTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpWorldcoinIdentityTransfer
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
				return ErrInvalidLengthOpWorldcoinIdentityTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpWorldcoinIdentityTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevState", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpWorldcoinIdentityTransfer
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
				return ErrInvalidLengthOpWorldcoinIdentityTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpWorldcoinIdentityTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrevState = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpWorldcoinIdentityTransfer
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
				return ErrInvalidLengthOpWorldcoinIdentityTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpWorldcoinIdentityTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpWorldcoinIdentityTransfer
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
				return ErrInvalidLengthOpWorldcoinIdentityTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpWorldcoinIdentityTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Timestamp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockNumber", wireType)
			}
			m.BlockNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpWorldcoinIdentityTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOpWorldcoinIdentityTransfer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOpWorldcoinIdentityTransfer
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
func skipOpWorldcoinIdentityTransfer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOpWorldcoinIdentityTransfer
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
					return 0, ErrIntOverflowOpWorldcoinIdentityTransfer
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
					return 0, ErrIntOverflowOpWorldcoinIdentityTransfer
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
				return 0, ErrInvalidLengthOpWorldcoinIdentityTransfer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOpWorldcoinIdentityTransfer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOpWorldcoinIdentityTransfer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOpWorldcoinIdentityTransfer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOpWorldcoinIdentityTransfer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOpWorldcoinIdentityTransfer = fmt.Errorf("proto: unexpected end of group")
)
