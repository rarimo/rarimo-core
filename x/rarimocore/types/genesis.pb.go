// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rarimocore/genesis.proto

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

// GenesisState defines the rarimocore module's genesis state.
type GenesisState struct {
	Params              Params            `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	OperationList       []Operation       `protobuf:"bytes,2,rep,name=operationList,proto3" json:"operationList"`
	ConfirmationList    []Confirmation    `protobuf:"bytes,3,rep,name=confirmationList,proto3" json:"confirmationList"`
	VoteList            []Vote            `protobuf:"bytes,4,rep,name=voteList,proto3" json:"voteList"`
	ViolationReportList []ViolationReport `protobuf:"bytes,5,rep,name=violationReportList,proto3" json:"violationReportList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c823eb3ed15d1a7, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetOperationList() []Operation {
	if m != nil {
		return m.OperationList
	}
	return nil
}

func (m *GenesisState) GetConfirmationList() []Confirmation {
	if m != nil {
		return m.ConfirmationList
	}
	return nil
}

func (m *GenesisState) GetVoteList() []Vote {
	if m != nil {
		return m.VoteList
	}
	return nil
}

func (m *GenesisState) GetViolationReportList() []ViolationReport {
	if m != nil {
		return m.ViolationReportList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "rarimo.rarimocore.rarimocore.GenesisState")
}

func init() { proto.RegisterFile("rarimocore/genesis.proto", fileDescriptor_7c823eb3ed15d1a7) }

var fileDescriptor_7c823eb3ed15d1a7 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xd2, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x06, 0xe0, 0xd6, 0xcd, 0x21, 0x99, 0x82, 0x44, 0xc5, 0x51, 0x34, 0xce, 0x21, 0x38, 0x84,
	0xb5, 0x32, 0xdf, 0x60, 0x0a, 0x82, 0x08, 0xca, 0x06, 0x5e, 0x88, 0x20, 0xdd, 0x88, 0x35, 0x60,
	0x77, 0x4a, 0x1a, 0x45, 0xdf, 0xc2, 0x67, 0xf1, 0x29, 0x76, 0xb9, 0x4b, 0xaf, 0x44, 0xda, 0x17,
	0x91, 0x9d, 0xc4, 0x9a, 0xa1, 0xd4, 0xab, 0x86, 0x9c, 0xf3, 0x7f, 0xf4, 0x84, 0x43, 0x1a, 0x32,
	0x94, 0x22, 0x86, 0x11, 0x48, 0x1e, 0x44, 0x7c, 0xcc, 0x53, 0x91, 0xfa, 0x89, 0x04, 0x05, 0x74,
	0x4b, 0x57, 0xfc, 0x9f, 0x06, 0xeb, 0xe8, 0xad, 0x47, 0x10, 0x01, 0x36, 0x06, 0xb3, 0x93, 0xce,
	0x78, 0x9b, 0x96, 0x96, 0x84, 0x32, 0x8c, 0x0d, 0xe6, 0x79, 0x56, 0x01, 0x12, 0x2e, 0x43, 0x25,
	0x60, 0x6c, 0x6a, 0xdb, 0x56, 0x6d, 0x04, 0xe3, 0x3b, 0x21, 0x63, 0xbb, 0xbc, 0x61, 0x95, 0x9f,
	0x40, 0x71, 0x73, 0xbd, 0x6b, 0x5f, 0x0b, 0x78, 0xc0, 0xc8, 0xad, 0xe4, 0x09, 0x48, 0xa5, 0x5b,
	0x5a, 0x6f, 0x15, 0xb2, 0x7c, 0xaa, 0x67, 0x1a, 0xa8, 0x50, 0x71, 0xda, 0x23, 0x35, 0xfd, 0x57,
	0x0d, 0xb7, 0xe9, 0xb6, 0xeb, 0xdd, 0x3d, 0xbf, 0x6c, 0x46, 0xff, 0x12, 0x7b, 0x7b, 0xd5, 0xc9,
	0xc7, 0x8e, 0xd3, 0x37, 0x49, 0x3a, 0x20, 0x2b, 0xc5, 0x00, 0xe7, 0x22, 0x55, 0x8d, 0x85, 0x66,
	0xa5, 0x5d, 0xef, 0xee, 0x97, 0x53, 0x17, 0xdf, 0x11, 0xa3, 0xcd, 0x1b, 0xf4, 0x86, 0xac, 0xda,
	0x93, 0xa3, 0x5b, 0x41, 0xf7, 0xa0, 0xdc, 0x3d, 0xb6, 0x52, 0x86, 0xfe, 0x25, 0xd1, 0x13, 0xb2,
	0x34, 0x7b, 0x38, 0x54, 0xab, 0xa8, 0xb6, 0xca, 0xd5, 0x2b, 0x50, 0xdc, 0x68, 0x45, 0x92, 0x72,
	0xb2, 0x56, 0xbc, 0x73, 0x1f, 0x9f, 0x19, 0xc1, 0x45, 0x04, 0x3b, 0xff, 0x80, 0xf3, 0x41, 0x63,
	0xff, 0xe5, 0xf5, 0xce, 0x26, 0x19, 0x73, 0xa7, 0x19, 0x73, 0x3f, 0x33, 0xe6, 0xbe, 0xe6, 0xcc,
	0x99, 0xe6, 0xcc, 0x79, 0xcf, 0x99, 0x73, 0x7d, 0x18, 0x09, 0x75, 0xff, 0x38, 0xf4, 0x47, 0x10,
	0x07, 0xda, 0x36, 0x9f, 0x0e, 0x2e, 0xc1, 0x73, 0x60, 0x6d, 0x84, 0x7a, 0x49, 0x78, 0x3a, 0xac,
	0xe1, 0x1e, 0x1c, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0x72, 0x7d, 0xb7, 0x84, 0xe5, 0x02, 0x00,
	0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ViolationReportList) > 0 {
		for iNdEx := len(m.ViolationReportList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ViolationReportList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.VoteList) > 0 {
		for iNdEx := len(m.VoteList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VoteList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ConfirmationList) > 0 {
		for iNdEx := len(m.ConfirmationList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ConfirmationList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.OperationList) > 0 {
		for iNdEx := len(m.OperationList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OperationList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.OperationList) > 0 {
		for _, e := range m.OperationList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ConfirmationList) > 0 {
		for _, e := range m.ConfirmationList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.VoteList) > 0 {
		for _, e := range m.VoteList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ViolationReportList) > 0 {
		for _, e := range m.ViolationReportList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperationList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperationList = append(m.OperationList, Operation{})
			if err := m.OperationList[len(m.OperationList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfirmationList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConfirmationList = append(m.ConfirmationList, Confirmation{})
			if err := m.ConfirmationList[len(m.ConfirmationList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VoteList = append(m.VoteList, Vote{})
			if err := m.VoteList[len(m.VoteList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ViolationReportList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ViolationReportList = append(m.ViolationReportList, ViolationReport{})
			if err := m.ViolationReportList[len(m.ViolationReportList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
