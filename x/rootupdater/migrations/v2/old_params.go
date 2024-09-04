package v2

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// OldParamas defines the parameters for the module.
type OldParamas struct {
	ContractAddress     string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	Root                string `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
	LastSignedRoot      string `protobuf:"bytes,3,opt,name=last_signed_root,json=lastSignedRoot,proto3" json:"last_signed_root,omitempty"`
	LastSignedRootIndex string `protobuf:"bytes,4,opt,name=last_signed_root_index,json=lastSignedRootIndex,proto3" json:"last_signed_root_index,omitempty"`
	EventName           string `protobuf:"bytes,5,opt,name=event_name,json=eventName,proto3" json:"event_name,omitempty"`
	RootTimestamp       int64  `protobuf:"varint,6,opt,name=root_timestamp,json=rootTimestamp,proto3" json:"root_timestamp,omitempty"`
	BlockHeight         uint64 `protobuf:"varint,7,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
}

func (m *OldParamas) Reset()         { *m = OldParamas{} }
func (m *OldParamas) String() string { return proto.CompactTextString(m) }
func (*OldParamas) ProtoMessage()    {}
func (*OldParamas) Descriptor() ([]byte, []int) {
	return fileDescriptor_31e4e5f76db9f10e, []int{0}
}
func (m *OldParamas) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OldParamas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
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
func (m *OldParamas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *OldParamas) XXX_Size() int {
	return m.Size()
}
func (m *OldParamas) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *OldParamas) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *OldParamas) GetRoot() string {
	if m != nil {
		return m.Root
	}
	return ""
}

func (m *OldParamas) GetLastSignedRoot() string {
	if m != nil {
		return m.LastSignedRoot
	}
	return ""
}

func (m *OldParamas) GetLastSignedRootIndex() string {
	if m != nil {
		return m.LastSignedRootIndex
	}
	return ""
}

func (m *OldParamas) GetEventName() string {
	if m != nil {
		return m.EventName
	}
	return ""
}

func (m *OldParamas) GetRootTimestamp() int64 {
	if m != nil {
		return m.RootTimestamp
	}
	return 0
}

func (m *OldParamas) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*OldParamas)(nil), "rarimo.rarimocore.rootupdater.Params")
}

func init() { proto.RegisterFile("rootupdater/params.proto", fileDescriptor_31e4e5f76db9f10e) }

var fileDescriptor_31e4e5f76db9f10e = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x9b, 0xb6, 0x56, 0x1a, 0xb5, 0x96, 0x28, 0x12, 0x84, 0x2e, 0x55, 0x10, 0xd6, 0x83,
	0xbb, 0x48, 0x9f, 0x40, 0x4f, 0x8a, 0x20, 0xb2, 0x7a, 0xf2, 0xb2, 0xa4, 0xbb, 0x61, 0x1b, 0x6c,
	0x36, 0x4b, 0x32, 0x95, 0xfa, 0x14, 0xfa, 0x58, 0x1e, 0x7b, 0xf4, 0x28, 0xed, 0x8b, 0x48, 0x66,
	0x2d, 0x54, 0x4f, 0x33, 0x7c, 0xff, 0x97, 0x39, 0xfc, 0xa1, 0xdc, 0x1a, 0x03, 0xb3, 0x2a, 0x17,
	0x20, 0x6d, 0x5c, 0x09, 0x2b, 0xb4, 0x8b, 0x2a, 0x6b, 0xc0, 0xb0, 0x81, 0x15, 0x56, 0x69, 0x13,
	0xd5, 0x23, 0x33, 0x56, 0x46, 0x1b, 0xee, 0xf1, 0x61, 0x61, 0x0a, 0x83, 0x66, 0xec, 0xb7, 0xfa,
	0xd1, 0xe9, 0x7b, 0x93, 0x76, 0x1e, 0xf0, 0x0a, 0x3b, 0xa7, 0xfd, 0xcc, 0x94, 0x60, 0x45, 0x06,
	0xa9, 0xc8, 0x73, 0x2b, 0x9d, 0xe3, 0x64, 0x48, 0xc2, 0x6e, 0xb2, 0xbf, 0xe6, 0x57, 0x35, 0x66,
	0x8c, 0xb6, 0xfd, 0x69, 0xde, 0xc4, 0x18, 0x77, 0x16, 0xd2, 0xfe, 0x54, 0x38, 0x48, 0x9d, 0x2a,
	0x4a, 0x99, 0xa7, 0x98, 0xb7, 0x30, 0xef, 0x79, 0xfe, 0x88, 0x38, 0xf1, 0xe6, 0x88, 0x1e, 0xfd,
	0x37, 0x53, 0x55, 0xe6, 0x72, 0xce, 0xdb, 0xe8, 0x1f, 0xfc, 0xf5, 0x6f, 0x7d, 0xc4, 0x06, 0x94,
	0xca, 0x57, 0x59, 0x42, 0x5a, 0x0a, 0x2d, 0xf9, 0x16, 0x8a, 0x5d, 0x24, 0xf7, 0x42, 0x4b, 0x76,
	0x46, 0x7b, 0x78, 0x07, 0x94, 0x96, 0x0e, 0x84, 0xae, 0x78, 0x67, 0x48, 0xc2, 0x56, 0xb2, 0xe7,
	0xe9, 0xd3, 0x1a, 0xb2, 0x13, 0xba, 0x3b, 0x9e, 0x9a, 0xec, 0x25, 0x9d, 0x48, 0x55, 0x4c, 0x80,
	0x6f, 0x0f, 0x49, 0xd8, 0x4e, 0x76, 0x90, 0xdd, 0x20, 0xba, 0xbe, 0xfb, 0x5c, 0x06, 0x64, 0xb1,
	0x0c, 0xc8, 0xf7, 0x32, 0x20, 0x1f, 0xab, 0xa0, 0xb1, 0x58, 0x05, 0x8d, 0xaf, 0x55, 0xd0, 0x78,
	0xbe, 0x2c, 0x14, 0x4c, 0x66, 0xe3, 0x28, 0x33, 0x3a, 0xae, 0x4b, 0xfe, 0x1d, 0x17, 0xbe, 0xec,
	0x78, 0x1e, 0x6f, 0x7e, 0x0d, 0xbc, 0x55, 0xd2, 0x8d, 0x3b, 0xd8, 0xf2, 0xe8, 0x27, 0x00, 0x00,
	0xff, 0xff, 0x0e, 0x6b, 0x5d, 0x49, 0xb6, 0x01, 0x00, 0x00,
}

func (m *OldParamas) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OldParamas) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OldParamas) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockHeight != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x38
	}
	if m.RootTimestamp != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RootTimestamp))
		i--
		dAtA[i] = 0x30
	}
	if len(m.EventName) > 0 {
		i -= len(m.EventName)
		copy(dAtA[i:], m.EventName)
		i = encodeVarintParams(dAtA, i, uint64(len(m.EventName)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.LastSignedRootIndex) > 0 {
		i -= len(m.LastSignedRootIndex)
		copy(dAtA[i:], m.LastSignedRootIndex)
		i = encodeVarintParams(dAtA, i, uint64(len(m.LastSignedRootIndex)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.LastSignedRoot) > 0 {
		i -= len(m.LastSignedRoot)
		copy(dAtA[i:], m.LastSignedRoot)
		i = encodeVarintParams(dAtA, i, uint64(len(m.LastSignedRoot)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Root) > 0 {
		i -= len(m.Root)
		copy(dAtA[i:], m.Root)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Root)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ContractAddress)))
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
func (m *OldParamas) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.Root)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.LastSignedRoot)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.LastSignedRootIndex)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.EventName)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.RootTimestamp != 0 {
		n += 1 + sovParams(uint64(m.RootTimestamp))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovParams(uint64(m.BlockHeight))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OldParamas) Unmarshal(dAtA []byte) error {
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
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
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
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Root", wireType)
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
			m.Root = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastSignedRoot", wireType)
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
			m.LastSignedRoot = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastSignedRootIndex", wireType)
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
			m.LastSignedRootIndex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventName", wireType)
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
			m.EventName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RootTimestamp", wireType)
			}
			m.RootTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RootTimestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
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
