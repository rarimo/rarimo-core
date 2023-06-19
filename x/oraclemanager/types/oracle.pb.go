// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: oraclemanager/oracle.proto

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

type OracleStatus int32

const (
	OracleStatus_Inactive OracleStatus = 0
	OracleStatus_Active   OracleStatus = 1
	OracleStatus_Jailed   OracleStatus = 2
	OracleStatus_Freezed  OracleStatus = 3
	OracleStatus_Slashed  OracleStatus = 4
)

var OracleStatus_name = map[int32]string{
	0: "Inactive",
	1: "Active",
	2: "Jailed",
	3: "Freezed",
	4: "Slashed",
}

var OracleStatus_value = map[string]int32{
	"Inactive": 0,
	"Active":   1,
	"Jailed":   2,
	"Freezed":  3,
	"Slashed":  4,
}

func (x OracleStatus) String() string {
	return proto.EnumName(OracleStatus_name, int32(x))
}

func (OracleStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6c1a9fae0dab29d9, []int{0}
}

type OracleIndex struct {
	Chain   string `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	Account string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
}

func (m *OracleIndex) Reset()         { *m = OracleIndex{} }
func (m *OracleIndex) String() string { return proto.CompactTextString(m) }
func (*OracleIndex) ProtoMessage()    {}
func (*OracleIndex) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c1a9fae0dab29d9, []int{0}
}
func (m *OracleIndex) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OracleIndex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OracleIndex.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OracleIndex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleIndex.Merge(m, src)
}
func (m *OracleIndex) XXX_Size() int {
	return m.Size()
}
func (m *OracleIndex) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleIndex.DiscardUnknown(m)
}

var xxx_messageInfo_OracleIndex proto.InternalMessageInfo

func (m *OracleIndex) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *OracleIndex) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type Delegation struct {
	Delegator string `protobuf:"bytes,1,opt,name=delegator,proto3" json:"delegator,omitempty"`
	Amount    string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *Delegation) Reset()         { *m = Delegation{} }
func (m *Delegation) String() string { return proto.CompactTextString(m) }
func (*Delegation) ProtoMessage()    {}
func (*Delegation) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c1a9fae0dab29d9, []int{1}
}
func (m *Delegation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Delegation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Delegation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Delegation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Delegation.Merge(m, src)
}
func (m *Delegation) XXX_Size() int {
	return m.Size()
}
func (m *Delegation) XXX_DiscardUnknown() {
	xxx_messageInfo_Delegation.DiscardUnknown(m)
}

var xxx_messageInfo_Delegation proto.InternalMessageInfo

func (m *Delegation) GetDelegator() string {
	if m != nil {
		return m.Delegator
	}
	return ""
}

func (m *Delegation) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

type Oracle struct {
	Index                 *OracleIndex `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Status                OracleStatus `protobuf:"varint,2,opt,name=status,proto3,enum=rarimo.rarimocore.oraclemanager.OracleStatus" json:"status,omitempty"`
	Stake                 string       `protobuf:"bytes,3,opt,name=stake,proto3" json:"stake,omitempty"`
	MissedCount           uint64       `protobuf:"varint,4,opt,name=missedCount,proto3" json:"missedCount,omitempty"`
	ViolationsCount       uint64       `protobuf:"varint,5,opt,name=violationsCount,proto3" json:"violationsCount,omitempty"`
	FreezeEndBlock        uint64       `protobuf:"varint,6,opt,name=freezeEndBlock,proto3" json:"freezeEndBlock,omitempty"`
	VotesCount            uint64       `protobuf:"varint,7,opt,name=votesCount,proto3" json:"votesCount,omitempty"`
	CreateOperationsCount uint64       `protobuf:"varint,8,opt,name=createOperationsCount,proto3" json:"createOperationsCount,omitempty"`
	Delegations           []Delegation `protobuf:"bytes,9,rep,name=delegations,proto3" json:"delegations"`
}

func (m *Oracle) Reset()         { *m = Oracle{} }
func (m *Oracle) String() string { return proto.CompactTextString(m) }
func (*Oracle) ProtoMessage()    {}
func (*Oracle) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c1a9fae0dab29d9, []int{2}
}
func (m *Oracle) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Oracle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Oracle.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Oracle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Oracle.Merge(m, src)
}
func (m *Oracle) XXX_Size() int {
	return m.Size()
}
func (m *Oracle) XXX_DiscardUnknown() {
	xxx_messageInfo_Oracle.DiscardUnknown(m)
}

var xxx_messageInfo_Oracle proto.InternalMessageInfo

func (m *Oracle) GetIndex() *OracleIndex {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *Oracle) GetStatus() OracleStatus {
	if m != nil {
		return m.Status
	}
	return OracleStatus_Inactive
}

func (m *Oracle) GetStake() string {
	if m != nil {
		return m.Stake
	}
	return ""
}

func (m *Oracle) GetMissedCount() uint64 {
	if m != nil {
		return m.MissedCount
	}
	return 0
}

func (m *Oracle) GetViolationsCount() uint64 {
	if m != nil {
		return m.ViolationsCount
	}
	return 0
}

func (m *Oracle) GetFreezeEndBlock() uint64 {
	if m != nil {
		return m.FreezeEndBlock
	}
	return 0
}

func (m *Oracle) GetVotesCount() uint64 {
	if m != nil {
		return m.VotesCount
	}
	return 0
}

func (m *Oracle) GetCreateOperationsCount() uint64 {
	if m != nil {
		return m.CreateOperationsCount
	}
	return 0
}

func (m *Oracle) GetDelegations() []Delegation {
	if m != nil {
		return m.Delegations
	}
	return nil
}

func init() {
	proto.RegisterEnum("rarimo.rarimocore.oraclemanager.OracleStatus", OracleStatus_name, OracleStatus_value)
	proto.RegisterType((*OracleIndex)(nil), "rarimo.rarimocore.oraclemanager.OracleIndex")
	proto.RegisterType((*Delegation)(nil), "rarimo.rarimocore.oraclemanager.Delegation")
	proto.RegisterType((*Oracle)(nil), "rarimo.rarimocore.oraclemanager.Oracle")
}

func init() { proto.RegisterFile("oraclemanager/oracle.proto", fileDescriptor_6c1a9fae0dab29d9) }

var fileDescriptor_6c1a9fae0dab29d9 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xb5, 0x9b, 0xc4, 0x69, 0xc6, 0x55, 0x89, 0x56, 0x05, 0x59, 0x15, 0x72, 0xa3, 0x1c, 0x90,
	0x05, 0xd4, 0x96, 0x52, 0xae, 0x1c, 0x30, 0x14, 0xa9, 0x48, 0x50, 0xc9, 0xb9, 0x71, 0xdb, 0xae,
	0x07, 0x77, 0x55, 0xdb, 0x1b, 0xed, 0x6e, 0xa3, 0xc2, 0x57, 0xf0, 0x59, 0x39, 0xf6, 0xc8, 0x09,
	0xa1, 0xe4, 0x47, 0x90, 0x77, 0x0d, 0x75, 0x2b, 0xa4, 0x70, 0xda, 0x79, 0x2f, 0x6f, 0x66, 0xdf,
	0xbc, 0x8d, 0xe1, 0x50, 0x48, 0xca, 0x4a, 0xac, 0x68, 0x4d, 0x0b, 0x94, 0x89, 0x45, 0xf1, 0x42,
	0x0a, 0x2d, 0xc8, 0x91, 0xa4, 0x92, 0x57, 0x22, 0xb6, 0x07, 0x13, 0x12, 0xe3, 0x7b, 0xea, 0xc3,
	0x83, 0x42, 0x14, 0xc2, 0x68, 0x93, 0xa6, 0xb2, 0x6d, 0xd3, 0xd7, 0xe0, 0x9f, 0x1b, 0xd9, 0x59,
	0x9d, 0xe3, 0x0d, 0x39, 0x80, 0x01, 0xbb, 0xa4, 0xbc, 0x0e, 0xdc, 0x89, 0x1b, 0x8d, 0x32, 0x0b,
	0x48, 0x00, 0x43, 0xca, 0x98, 0xb8, 0xae, 0x75, 0xb0, 0x63, 0xf8, 0x3f, 0x70, 0x9a, 0x02, 0xbc,
	0xc3, 0x12, 0x0b, 0xaa, 0xb9, 0xa8, 0xc9, 0x53, 0x18, 0xe5, 0x16, 0x09, 0xd9, 0x4e, 0xb8, 0x23,
	0xc8, 0x13, 0xf0, 0x68, 0xd5, 0x19, 0xd2, 0xa2, 0xe9, 0xaa, 0x07, 0x9e, 0xf5, 0x40, 0x52, 0x18,
	0xf0, 0xc6, 0x87, 0x69, 0xf6, 0x67, 0x2f, 0xe3, 0x2d, 0x4b, 0xc5, 0x1d, 0xef, 0x99, 0x6d, 0x25,
	0xa7, 0xe0, 0x29, 0x4d, 0xf5, 0xb5, 0x32, 0xd7, 0xec, 0xcf, 0x8e, 0xff, 0x73, 0xc8, 0xdc, 0x34,
	0x65, 0x6d, 0x73, 0x93, 0x84, 0xd2, 0xf4, 0x0a, 0x83, 0x9e, 0x4d, 0xc2, 0x00, 0x32, 0x01, 0xbf,
	0xe2, 0x4a, 0x61, 0xfe, 0xd6, 0x2c, 0xd2, 0x9f, 0xb8, 0x51, 0x3f, 0xeb, 0x52, 0x24, 0x82, 0x47,
	0x4b, 0x2e, 0x4a, 0x13, 0x88, 0xb2, 0xaa, 0x81, 0x51, 0x3d, 0xa4, 0xc9, 0x33, 0xd8, 0xff, 0x22,
	0x11, 0xbf, 0xe1, 0x69, 0x9d, 0xa7, 0xa5, 0x60, 0x57, 0x81, 0x67, 0x84, 0x0f, 0x58, 0x12, 0x02,
	0x2c, 0x85, 0xc6, 0x76, 0xd8, 0xd0, 0x68, 0x3a, 0x0c, 0x79, 0x05, 0x8f, 0x99, 0x44, 0xaa, 0xf1,
	0x7c, 0x81, 0xb2, 0x7b, 0xef, 0xae, 0x91, 0xfe, 0xfb, 0x47, 0x32, 0x07, 0x3f, 0xff, 0xfb, 0x72,
	0x2a, 0x18, 0x4d, 0x7a, 0x91, 0x3f, 0x7b, 0xb1, 0x35, 0xab, 0xbb, 0xd7, 0x4e, 0xfb, 0xab, 0x9f,
	0x47, 0x4e, 0xd6, 0x9d, 0xf2, 0xfc, 0x13, 0xec, 0x75, 0xc3, 0x24, 0x7b, 0xb0, 0x7b, 0x56, 0x53,
	0xa6, 0xf9, 0x12, 0xc7, 0x0e, 0x01, 0xf0, 0xde, 0xd8, 0xda, 0x6d, 0xea, 0x0f, 0x94, 0x97, 0x98,
	0x8f, 0x77, 0x88, 0x0f, 0xc3, 0xf7, 0x66, 0xe5, 0x7c, 0xdc, 0x6b, 0xc0, 0xbc, 0xa4, 0xea, 0x12,
	0xf3, 0x71, 0x3f, 0xfd, 0xb8, 0x5a, 0x87, 0xee, 0xed, 0x3a, 0x74, 0x7f, 0xad, 0x43, 0xf7, 0xfb,
	0x26, 0x74, 0x6e, 0x37, 0xa1, 0xf3, 0x63, 0x13, 0x3a, 0x9f, 0x4f, 0x0a, 0xae, 0x4b, 0x7a, 0x11,
	0x33, 0x51, 0x25, 0xd6, 0x6c, 0x7b, 0x1c, 0x37, 0xa6, 0x93, 0x9b, 0xe4, 0xfe, 0xa7, 0xa2, 0xbf,
	0x2e, 0x50, 0x5d, 0x78, 0xe6, 0x3f, 0x7f, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x95, 0xa5, 0xb1,
	0xbc, 0x48, 0x03, 0x00, 0x00,
}

func (m *OracleIndex) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OracleIndex) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OracleIndex) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Delegation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Delegation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Delegation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Delegator) > 0 {
		i -= len(m.Delegator)
		copy(dAtA[i:], m.Delegator)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Delegator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Oracle) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Oracle) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Oracle) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Delegations) > 0 {
		for iNdEx := len(m.Delegations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Delegations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOracle(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.CreateOperationsCount != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.CreateOperationsCount))
		i--
		dAtA[i] = 0x40
	}
	if m.VotesCount != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.VotesCount))
		i--
		dAtA[i] = 0x38
	}
	if m.FreezeEndBlock != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.FreezeEndBlock))
		i--
		dAtA[i] = 0x30
	}
	if m.ViolationsCount != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.ViolationsCount))
		i--
		dAtA[i] = 0x28
	}
	if m.MissedCount != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.MissedCount))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Stake) > 0 {
		i -= len(m.Stake)
		copy(dAtA[i:], m.Stake)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Stake)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Status != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if m.Index != nil {
		{
			size, err := m.Index.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOracle(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOracle(dAtA []byte, offset int, v uint64) int {
	offset -= sovOracle(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OracleIndex) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	return n
}

func (m *Delegation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Delegator)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	return n
}

func (m *Oracle) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != nil {
		l = m.Index.Size()
		n += 1 + l + sovOracle(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovOracle(uint64(m.Status))
	}
	l = len(m.Stake)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	if m.MissedCount != 0 {
		n += 1 + sovOracle(uint64(m.MissedCount))
	}
	if m.ViolationsCount != 0 {
		n += 1 + sovOracle(uint64(m.ViolationsCount))
	}
	if m.FreezeEndBlock != 0 {
		n += 1 + sovOracle(uint64(m.FreezeEndBlock))
	}
	if m.VotesCount != 0 {
		n += 1 + sovOracle(uint64(m.VotesCount))
	}
	if m.CreateOperationsCount != 0 {
		n += 1 + sovOracle(uint64(m.CreateOperationsCount))
	}
	if len(m.Delegations) > 0 {
		for _, e := range m.Delegations {
			l = e.Size()
			n += 1 + l + sovOracle(uint64(l))
		}
	}
	return n
}

func sovOracle(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOracle(x uint64) (n int) {
	return sovOracle(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OracleIndex) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
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
			return fmt.Errorf("proto: OracleIndex: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OracleIndex: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
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
func (m *Delegation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
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
			return fmt.Errorf("proto: Delegation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Delegation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Delegator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
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
func (m *Oracle) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
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
			return fmt.Errorf("proto: Oracle: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Oracle: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Index == nil {
				m.Index = &OracleIndex{}
			}
			if err := m.Index.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= OracleStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stake", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stake = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissedCount", wireType)
			}
			m.MissedCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MissedCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ViolationsCount", wireType)
			}
			m.ViolationsCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ViolationsCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FreezeEndBlock", wireType)
			}
			m.FreezeEndBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FreezeEndBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotesCount", wireType)
			}
			m.VotesCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotesCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateOperationsCount", wireType)
			}
			m.CreateOperationsCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreateOperationsCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Delegations = append(m.Delegations, Delegation{})
			if err := m.Delegations[len(m.Delegations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
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
func skipOracle(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOracle
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
					return 0, ErrIntOverflowOracle
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
					return 0, ErrIntOverflowOracle
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
				return 0, ErrInvalidLengthOracle
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOracle
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOracle
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOracle        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOracle          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOracle = fmt.Errorf("proto: unexpected end of group")
)