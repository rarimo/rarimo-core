// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: identity/state.proto

package types

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

type Node struct {
	// Node key (identity state hash)
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Node priority (should be random)
	Priority uint64 `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	// Node left son key
	Left string `protobuf:"bytes,4,opt,name=left,proto3" json:"left,omitempty"`
	// Node right son key
	Right string `protobuf:"bytes,5,opt,name=right,proto3" json:"right,omitempty"`
	// Merkle hash. H = Hash(Hash(left_key|right_key)|self_key)
	Hash string `protobuf:"bytes,6,opt,name=hash,proto3" json:"hash,omitempty"`
	// Hash(left_key|right_key)
	ChildrenHash string `protobuf:"bytes,7,opt,name=childrenHash,proto3" json:"childrenHash,omitempty"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_a879c8de568b961b, []int{0}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Node.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return m.Size()
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Node) GetPriority() uint64 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Node) GetLeft() string {
	if m != nil {
		return m.Left
	}
	return ""
}

func (m *Node) GetRight() string {
	if m != nil {
		return m.Right
	}
	return ""
}

func (m *Node) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Node) GetChildrenHash() string {
	if m != nil {
		return m.ChildrenHash
	}
	return ""
}

type StateInfo struct {
	// State info index (issuer id)
	Index string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	// State hash in hex with 0x
	Hash string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// Creation timestamps
	CreatedAtTimestamp uint64 `protobuf:"varint,3,opt,name=createdAtTimestamp,proto3" json:"createdAtTimestamp,omitempty"`
	// Index of last update/create operation
	LastUpdateOperationIndex string `protobuf:"bytes,4,opt,name=lastUpdateOperationIndex,proto3" json:"lastUpdateOperationIndex,omitempty"`
}

func (m *StateInfo) Reset()         { *m = StateInfo{} }
func (m *StateInfo) String() string { return proto.CompactTextString(m) }
func (*StateInfo) ProtoMessage()    {}
func (*StateInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a879c8de568b961b, []int{1}
}
func (m *StateInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StateInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StateInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StateInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateInfo.Merge(m, src)
}
func (m *StateInfo) XXX_Size() int {
	return m.Size()
}
func (m *StateInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_StateInfo.DiscardUnknown(m)
}

var xxx_messageInfo_StateInfo proto.InternalMessageInfo

func (m *StateInfo) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *StateInfo) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *StateInfo) GetCreatedAtTimestamp() uint64 {
	if m != nil {
		return m.CreatedAtTimestamp
	}
	return 0
}

func (m *StateInfo) GetLastUpdateOperationIndex() string {
	if m != nil {
		return m.LastUpdateOperationIndex
	}
	return ""
}

func init() {
	proto.RegisterType((*Node)(nil), "rarimo.rarimocore.identity.Node")
	proto.RegisterType((*StateInfo)(nil), "rarimo.rarimocore.identity.StateInfo")
}

func init() { proto.RegisterFile("identity/state.proto", fileDescriptor_a879c8de568b961b) }

var fileDescriptor_a879c8de568b961b = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0xad, 0xdb, 0xb4, 0x50, 0x8b, 0x01, 0x59, 0x1d, 0xac, 0x0e, 0x51, 0xd5, 0xa9, 0x0b, 0xc9,
	0xc0, 0xc6, 0x06, 0x53, 0xbb, 0x80, 0x54, 0x60, 0x61, 0x73, 0xe3, 0x6b, 0x62, 0x91, 0xc4, 0x96,
	0x73, 0x43, 0xf3, 0x17, 0x4c, 0xec, 0xfc, 0x0d, 0x63, 0x47, 0x46, 0xd4, 0xfe, 0x08, 0xb2, 0x5d,
	0x2a, 0x90, 0x60, 0xba, 0x77, 0xf7, 0x9e, 0x9f, 0xef, 0xf4, 0xe8, 0x48, 0x49, 0xa8, 0x51, 0x61,
	0x9b, 0x36, 0x28, 0x10, 0x12, 0x63, 0x35, 0x6a, 0x36, 0xb6, 0xc2, 0xaa, 0x4a, 0x27, 0xa1, 0x64,
	0xda, 0x42, 0xf2, 0xad, 0x1b, 0x8f, 0x72, 0x9d, 0x6b, 0x2f, 0x4b, 0x1d, 0x0a, 0x2f, 0xa6, 0xaf,
	0x84, 0x46, 0xb7, 0x5a, 0x02, 0x3b, 0xa7, 0xbd, 0x67, 0x68, 0x39, 0x99, 0x90, 0xd9, 0x70, 0xe9,
	0x20, 0x1b, 0xd3, 0x53, 0x63, 0x95, 0xb6, 0x0a, 0x5b, 0xde, 0x9d, 0x90, 0x59, 0xb4, 0x3c, 0xf6,
	0x8c, 0xd1, 0xa8, 0x84, 0x35, 0xf2, 0xc8, 0xcb, 0x3d, 0x66, 0x23, 0xda, 0xb7, 0x2a, 0x2f, 0x90,
	0xf7, 0xfd, 0x30, 0x34, 0x4e, 0x59, 0x88, 0xa6, 0xe0, 0x83, 0xa0, 0x74, 0x98, 0x4d, 0xe9, 0x59,
	0x56, 0xa8, 0x52, 0x5a, 0xa8, 0xe7, 0x8e, 0x3b, 0xf1, 0xdc, 0xaf, 0xd9, 0xf4, 0x8d, 0xd0, 0xe1,
	0xbd, 0x3b, 0x6d, 0x51, 0xaf, 0xb5, 0xf3, 0x56, 0xb5, 0x84, 0xcd, 0x61, 0xbf, 0xd0, 0x1c, 0xbd,
	0xbb, 0x3f, 0xbc, 0x13, 0xca, 0x32, 0x0b, 0x02, 0x41, 0x5e, 0xe3, 0x83, 0xaa, 0xa0, 0x41, 0x51,
	0x19, 0xde, 0xf3, 0xfb, 0xff, 0xc1, 0xb0, 0x2b, 0xca, 0x4b, 0xd1, 0xe0, 0xa3, 0x91, 0x02, 0xe1,
	0xce, 0x80, 0x15, 0xa8, 0x74, 0xbd, 0xf0, 0x9f, 0x85, 0xeb, 0xfe, 0xe5, 0x6f, 0xe6, 0xef, 0xbb,
	0x98, 0x6c, 0x77, 0x31, 0xf9, 0xdc, 0xc5, 0xe4, 0x65, 0x1f, 0x77, 0xb6, 0xfb, 0xb8, 0xf3, 0xb1,
	0x8f, 0x3b, 0x4f, 0x49, 0xae, 0xb0, 0x14, 0xab, 0x24, 0xd3, 0x55, 0x1a, 0xc2, 0x38, 0x94, 0x0b,
	0x17, 0x4a, 0xba, 0x49, 0x8f, 0xf1, 0x61, 0x6b, 0xa0, 0x59, 0x0d, 0x7c, 0x1a, 0x97, 0x5f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x26, 0x02, 0x98, 0x1b, 0xd7, 0x01, 0x00, 0x00,
}

func (m *Node) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Node) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Node) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChildrenHash) > 0 {
		i -= len(m.ChildrenHash)
		copy(dAtA[i:], m.ChildrenHash)
		i = encodeVarintState(dAtA, i, uint64(len(m.ChildrenHash)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintState(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Right) > 0 {
		i -= len(m.Right)
		copy(dAtA[i:], m.Right)
		i = encodeVarintState(dAtA, i, uint64(len(m.Right)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Left) > 0 {
		i -= len(m.Left)
		copy(dAtA[i:], m.Left)
		i = encodeVarintState(dAtA, i, uint64(len(m.Left)))
		i--
		dAtA[i] = 0x22
	}
	if m.Priority != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Priority))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintState(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StateInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StateInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StateInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LastUpdateOperationIndex) > 0 {
		i -= len(m.LastUpdateOperationIndex)
		copy(dAtA[i:], m.LastUpdateOperationIndex)
		i = encodeVarintState(dAtA, i, uint64(len(m.LastUpdateOperationIndex)))
		i--
		dAtA[i] = 0x22
	}
	if m.CreatedAtTimestamp != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.CreatedAtTimestamp))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintState(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintState(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintState(dAtA []byte, offset int, v uint64) int {
	offset -= sovState(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Node) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.Priority != 0 {
		n += 1 + sovState(uint64(m.Priority))
	}
	l = len(m.Left)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.Right)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.ChildrenHash)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	return n
}

func (m *StateInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.CreatedAtTimestamp != 0 {
		n += 1 + sovState(uint64(m.CreatedAtTimestamp))
	}
	l = len(m.LastUpdateOperationIndex)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	return n
}

func sovState(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozState(x uint64) (n int) {
	return sovState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Node) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: Node: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Node: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Priority", wireType)
			}
			m.Priority = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Priority |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Left", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Left = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Right", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Right = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChildrenHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChildrenHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func (m *StateInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: StateInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StateInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAtTimestamp", wireType)
			}
			m.CreatedAtTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAtTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdateOperationIndex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastUpdateOperationIndex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func skipState(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
				return 0, ErrInvalidLengthState
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupState
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthState
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthState        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowState          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupState = fmt.Errorf("proto: unexpected end of group")
)
