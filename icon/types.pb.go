// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: icon/types/v1/types.proto

package icon

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

// BlockIdFlag indicates which BlcokID the signature is for
type BlockIDFlag int32

const (
	BlockIDFlagUnknown BlockIDFlag = 0
	BlockIDFlagAbsent  BlockIDFlag = 1
	BlockIDFlagCommit  BlockIDFlag = 2
	BlockIDFlagNil     BlockIDFlag = 3
)

var BlockIDFlag_name = map[int32]string{
	0: "BLOCK_ID_FLAG_UNKNOWN",
	1: "BLOCK_ID_FLAG_ABSENT",
	2: "BLOCK_ID_FLAG_COMMIT",
	3: "BLOCK_ID_FLAG_NIL",
}

var BlockIDFlag_value = map[string]int32{
	"BLOCK_ID_FLAG_UNKNOWN": 0,
	"BLOCK_ID_FLAG_ABSENT":  1,
	"BLOCK_ID_FLAG_COMMIT":  2,
	"BLOCK_ID_FLAG_NIL":     3,
}

func (x BlockIDFlag) String() string {
	return proto.EnumName(BlockIDFlag_name, int32(x))
}

func (BlockIDFlag) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eb5e993a7dd5d0ff, []int{0}
}

// SignedMsgType is a type of signed message in the consensus.
type SignedMsgType int32

const (
	UnknownType SignedMsgType = 0
	// Votes
	PrevoteType   SignedMsgType = 1
	PrecommitType SignedMsgType = 2
	// Proposals
	ProposalType SignedMsgType = 32
)

var SignedMsgType_name = map[int32]string{
	0:  "SIGNED_MSG_TYPE_UNKNOWN",
	1:  "SIGNED_MSG_TYPE_PREVOTE",
	2:  "SIGNED_MSG_TYPE_PRECOMMIT",
	32: "SIGNED_MSG_TYPE_PROPOSAL",
}

var SignedMsgType_value = map[string]int32{
	"SIGNED_MSG_TYPE_UNKNOWN":   0,
	"SIGNED_MSG_TYPE_PREVOTE":   1,
	"SIGNED_MSG_TYPE_PRECOMMIT": 2,
	"SIGNED_MSG_TYPE_PROPOSAL":  32,
}

func (x SignedMsgType) String() string {
	return proto.EnumName(SignedMsgType_name, int32(x))
}

func (SignedMsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eb5e993a7dd5d0ff, []int{1}
}

type SignedHeader struct {
	Header     *BTPHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Signatures [][]byte   `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

func (m *SignedHeader) Reset()         { *m = SignedHeader{} }
func (m *SignedHeader) String() string { return proto.CompactTextString(m) }
func (*SignedHeader) ProtoMessage()    {}
func (*SignedHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb5e993a7dd5d0ff, []int{0}
}
func (m *SignedHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignedHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignedHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignedHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedHeader.Merge(m, src)
}
func (m *SignedHeader) XXX_Size() int {
	return m.Size()
}
func (m *SignedHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedHeader.DiscardUnknown(m)
}

var xxx_messageInfo_SignedHeader proto.InternalMessageInfo

type BTPHeader struct {
	MainHeight             uint64        `protobuf:"varint,1,opt,name=main_height,json=mainHeight,proto3" json:"main_height,omitempty"`
	Round                  uint32        `protobuf:"varint,2,opt,name=round,proto3" json:"round,omitempty"`
	NextProofContextHash   []byte        `protobuf:"bytes,3,opt,name=next_proof_context_hash,json=nextProofContextHash,proto3" json:"next_proof_context_hash,omitempty"`
	NetworkSectionToRoot   []*MerkleNode `protobuf:"bytes,4,rep,name=network_section_to_root,json=networkSectionToRoot,proto3" json:"network_section_to_root,omitempty"`
	NetworkId              uint64        `protobuf:"varint,5,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	UpdateNumber           uint64        `protobuf:"varint,6,opt,name=update_number,json=updateNumber,proto3" json:"update_number,omitempty"`
	PrevNetworkSectionHash []byte        `protobuf:"bytes,7,opt,name=prev_network_section_hash,json=prevNetworkSectionHash,proto3" json:"prev_network_section_hash,omitempty"`
	MessageCount           uint64        `protobuf:"varint,8,opt,name=message_count,json=messageCount,proto3" json:"message_count,omitempty"`
	MessageRoot            []byte        `protobuf:"bytes,9,opt,name=message_root,json=messageRoot,proto3" json:"message_root,omitempty"`
	NextValidators         [][]byte      `protobuf:"bytes,10,rep,name=nextValidators,proto3" json:"nextValidators,omitempty"`
}


func (m *BTPHeader) Reset()         { *m = BTPHeader{} }
func (m *BTPHeader) String() string { return proto.CompactTextString(m) }
func (*BTPHeader) ProtoMessage()    {}
func (*BTPHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb5e993a7dd5d0ff, []int{1}
}
func (m *BTPHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BTPHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BTPHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BTPHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BTPHeader.Merge(m, src)
}
func (m *BTPHeader) XXX_Size() int {
	return m.Size()
}
func (m *BTPHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BTPHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BTPHeader proto.InternalMessageInfo

type MerkleNode struct {
	Dir   int32  `protobuf:"varint,1,opt,name=Dir,proto3" json:"Dir,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *MerkleNode) Reset()         { *m = MerkleNode{} }
func (m *MerkleNode) String() string { return proto.CompactTextString(m) }
func (*MerkleNode) ProtoMessage()    {}
func (*MerkleNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb5e993a7dd5d0ff, []int{2}
}
func (m *MerkleNode) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MerkleNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MerkleNode.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MerkleNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerkleNode.Merge(m, src)
}
func (m *MerkleNode) XXX_Size() int {
	return m.Size()
}
func (m *MerkleNode) XXX_DiscardUnknown() {
	xxx_messageInfo_MerkleNode.DiscardUnknown(m)
}

var xxx_messageInfo_MerkleNode proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("icon.types.v1.BlockIDFlag", BlockIDFlag_name, BlockIDFlag_value)
	proto.RegisterEnum("icon.types.v1.SignedMsgType", SignedMsgType_name, SignedMsgType_value)
	proto.RegisterType((*SignedHeader)(nil), "icon.types.v1.SignedHeader")
	proto.RegisterType((*BTPHeader)(nil), "icon.types.v1.BTPHeader")
	proto.RegisterType((*MerkleNode)(nil), "icon.types.v1.MerkleNode")
}

func init() { proto.RegisterFile("icon/types/v1/types.proto", fileDescriptor_eb5e993a7dd5d0ff) }

var fileDescriptor_eb5e993a7dd5d0ff = []byte{
	// 785 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x4d, 0x8f, 0xda, 0x46,
	0x18, 0xc6, 0x90, 0xdd, 0x66, 0x07, 0x48, 0xcc, 0x88, 0x24, 0x06, 0x29, 0xae, 0xbb, 0x91, 0x2a,
	0x5a, 0x55, 0x10, 0xd2, 0xf6, 0x50, 0x7a, 0xe2, 0x2b, 0xac, 0x15, 0x30, 0x96, 0x71, 0xe8, 0x87,
	0x56, 0x72, 0x07, 0x3c, 0x35, 0x16, 0xc6, 0x83, 0xc6, 0x03, 0xe9, 0xfe, 0x83, 0x8a, 0x53, 0xff,
	0x00, 0x97, 0xb6, 0x87, 0xaa, 0xa7, 0x9e, 0xfb, 0x0b, 0x56, 0x3d, 0xed, 0xad, 0x3d, 0xb6, 0xec,
	0xad, 0xbf, 0xa2, 0x9a, 0x31, 0xbb, 0x0b, 0xbb, 0xcd, 0x05, 0xbd, 0xef, 0xf3, 0x3e, 0xcf, 0xc3,
	0xfb, 0x78, 0xec, 0x01, 0x05, 0x7f, 0x4c, 0xc2, 0x0a, 0x3b, 0x9b, 0xe3, 0xa8, 0xb2, 0xac, 0xc6,
	0x45, 0x79, 0x4e, 0x09, 0x23, 0x30, 0xcb, 0x47, 0xe5, 0x18, 0x59, 0x56, 0x8b, 0x79, 0x8f, 0x78,
	0x44, 0x4c, 0x2a, 0xbc, 0x8a, 0x49, 0xc7, 0xdf, 0x80, 0xcc, 0xc0, 0xf7, 0x42, 0xec, 0x9e, 0x60,
	0xe4, 0x62, 0x0a, 0x9f, 0x83, 0xc3, 0x89, 0xa8, 0x14, 0x49, 0x93, 0x4a, 0xe9, 0x17, 0x4a, 0x79,
	0xcf, 0xa5, 0xdc, 0xb0, 0xcd, 0x98, 0x69, 0x6d, 0x79, 0x50, 0x05, 0x20, 0xf2, 0xbd, 0x10, 0xb1,
	0x05, 0xc5, 0x91, 0x92, 0xd4, 0x52, 0xa5, 0x8c, 0xb5, 0x83, 0x1c, 0xff, 0x9e, 0x02, 0x47, 0xd7,
	0x2a, 0xf8, 0x2e, 0x48, 0xcf, 0x90, 0x1f, 0x3a, 0x13, 0xec, 0x7b, 0x13, 0x26, 0xfe, 0xe4, 0x9e,
	0x05, 0x38, 0x74, 0x22, 0x10, 0x98, 0x07, 0x07, 0x94, 0x2c, 0x42, 0x57, 0x49, 0x6a, 0x52, 0x29,
	0x6b, 0xc5, 0x0d, 0xfc, 0x14, 0x3c, 0x09, 0xf1, 0x77, 0xcc, 0x99, 0x53, 0x42, 0xbe, 0x75, 0xc6,
	0x24, 0x64, 0xbc, 0x9b, 0xa0, 0x68, 0xa2, 0xa4, 0x34, 0xa9, 0x94, 0xb1, 0xf2, 0x7c, 0x6c, 0xf2,
	0x69, 0x33, 0x1e, 0x9e, 0xa0, 0x68, 0x02, 0x4d, 0x2e, 0x63, 0x6f, 0x08, 0x9d, 0x3a, 0x11, 0x1e,
	0x33, 0x9f, 0x84, 0x0e, 0x23, 0x0e, 0x25, 0x84, 0x29, 0xf7, 0xb4, 0x54, 0x29, 0xfd, 0xa2, 0x70,
	0x2b, 0x5e, 0x0f, 0xd3, 0x69, 0x80, 0x0d, 0xe2, 0x62, 0xee, 0x28, 0x94, 0x83, 0x58, 0x68, 0x13,
	0x8b, 0x10, 0x06, 0x9f, 0x02, 0x70, 0xe5, 0xe8, 0xbb, 0xca, 0x81, 0x58, 0xff, 0x68, 0x8b, 0xe8,
	0x2e, 0x7c, 0x06, 0xb2, 0x8b, 0xb9, 0x8b, 0x18, 0x76, 0xc2, 0xc5, 0x6c, 0x84, 0xa9, 0x72, 0x28,
	0x18, 0x99, 0x18, 0x34, 0x04, 0x06, 0x3f, 0x03, 0x85, 0x39, 0xc5, 0x4b, 0xe7, 0xf6, 0x6a, 0x22,
	0xce, 0x3b, 0x22, 0xce, 0x63, 0x4e, 0x30, 0xf6, 0x16, 0x10, 0x81, 0x9e, 0x81, 0xec, 0x0c, 0x47,
	0x11, 0xf2, 0xb0, 0x33, 0x26, 0x8b, 0x90, 0x29, 0xf7, 0x63, 0xff, 0x2d, 0xd8, 0xe4, 0x18, 0x7c,
	0x0f, 0x5c, 0xf5, 0x71, 0xd4, 0x23, 0x61, 0x99, 0xde, 0x62, 0x22, 0xc6, 0xfb, 0xe0, 0x01, 0x7f,
	0x60, 0x43, 0x14, 0xf8, 0x2e, 0x62, 0x84, 0x46, 0x0a, 0x10, 0x07, 0x77, 0x0b, 0x3d, 0xfe, 0x04,
	0x80, 0x9b, 0x47, 0x02, 0x65, 0x90, 0x6a, 0xf9, 0xf1, 0x9b, 0x71, 0x60, 0xf1, 0x92, 0x9f, 0xd6,
	0x12, 0x05, 0x0b, 0x2c, 0x4e, 0x2b, 0x63, 0xc5, 0xcd, 0x87, 0x7f, 0x4a, 0x20, 0xdd, 0x08, 0xc8,
	0x78, 0xaa, 0xb7, 0x5e, 0x06, 0xc8, 0x83, 0x55, 0xf0, 0xa8, 0xd1, 0xed, 0x37, 0x5f, 0x39, 0x7a,
	0xcb, 0x79, 0xd9, 0xad, 0x77, 0x9c, 0xd7, 0xc6, 0x2b, 0xa3, 0xff, 0x85, 0x21, 0x27, 0x8a, 0x8f,
	0x57, 0x6b, 0x0d, 0xee, 0x70, 0x5f, 0x87, 0xd3, 0x90, 0xbc, 0x09, 0x61, 0x05, 0xe4, 0xf7, 0x25,
	0xf5, 0xc6, 0xa0, 0x6d, 0xd8, 0xb2, 0x54, 0x7c, 0xb4, 0x5a, 0x6b, 0xb9, 0x1d, 0x45, 0x7d, 0x14,
	0xe1, 0x90, 0xdd, 0x15, 0x34, 0xfb, 0xbd, 0x9e, 0x6e, 0xcb, 0xc9, 0x3b, 0x82, 0x26, 0x99, 0xcd,
	0x7c, 0x06, 0x3f, 0x00, 0xb9, 0x7d, 0x81, 0xa1, 0x77, 0xe5, 0x54, 0x11, 0xae, 0xd6, 0xda, 0x83,
	0x1d, 0xb6, 0xe1, 0x07, 0xc5, 0xfb, 0xdf, 0xff, 0xa4, 0x26, 0x7e, 0xf9, 0x59, 0x95, 0x78, 0xb2,
	0x6c, 0xfc, 0xbd, 0xf4, 0x22, 0xcf, 0x3e, 0x9b, 0x63, 0xf8, 0x11, 0x78, 0x32, 0xd0, 0x3b, 0x46,
	0xbb, 0xe5, 0xf4, 0x06, 0x1d, 0xc7, 0xfe, 0xca, 0x6c, 0xef, 0xa4, 0x7b, 0xb8, 0x5a, 0x6b, 0xe9,
	0x6d, 0xa4, 0xb7, 0xb1, 0x4d, 0xab, 0x3d, 0xec, 0xdb, 0x6d, 0x59, 0x8a, 0xd9, 0x26, 0xc5, 0x4b,
	0xc2, 0xb0, 0x60, 0x3f, 0x07, 0x85, 0xff, 0x61, 0x5f, 0x07, 0xcb, 0xad, 0xd6, 0x5a, 0xd6, 0xa4,
	0x78, 0x2c, 0x02, 0x09, 0x45, 0x19, 0x28, 0x77, 0x15, 0x7d, 0xb3, 0x3f, 0xa8, 0x77, 0x65, 0xad,
	0x28, 0xaf, 0xd6, 0x5a, 0xc6, 0xa4, 0x64, 0x4e, 0x22, 0x14, 0x70, 0xfe, 0x4d, 0xb2, 0xc6, 0x6f,
	0xd2, 0xf9, 0x3f, 0x6a, 0xe2, 0x7c, 0xa3, 0x4a, 0x17, 0x1b, 0x55, 0xfa, 0x7b, 0xa3, 0x4a, 0x3f,
	0x5c, 0xaa, 0x89, 0x8b, 0x4b, 0x35, 0xf1, 0xd7, 0xa5, 0x9a, 0x00, 0xb9, 0x31, 0x99, 0xed, 0x7f,
	0x2b, 0x0d, 0xc0, 0x1d, 0x22, 0x93, 0x5f, 0x23, 0xa6, 0xf4, 0xf5, 0xd3, 0xc0, 0x1f, 0x51, 0x44,
	0x7d, 0x1c, 0x55, 0x3c, 0x52, 0xe1, 0x2b, 0x91, 0xb0, 0xc2, 0x05, 0x9f, 0xf3, 0x9f, 0x1f, 0x93,
	0x29, 0xdd, 0xfe, 0xf2, 0xd7, 0x64, 0x56, 0xe7, 0x16, 0x42, 0x58, 0x1e, 0x56, 0xff, 0x88, 0xfb,
	0x53, 0xd1, 0x9f, 0x0e, 0xab, 0x9b, 0x64, 0x61, 0xaf, 0x3f, 0xed, 0x98, 0x8d, 0x1e, 0x66, 0xc8,
	0x45, 0x0c, 0xfd, 0x9b, 0x7c, 0xc8, 0x67, 0xb5, 0x9a, 0x18, 0xd6, 0x6a, 0xc3, 0xea, 0xe8, 0x50,
	0x5c, 0x61, 0x1f, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xd8, 0x7e, 0x9a, 0x04, 0x05, 0x00,
	0x00,
}

func (m *SignedHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignedHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignedHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signatures) > 0 {
		for iNdEx := len(m.Signatures) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signatures[iNdEx])
			copy(dAtA[i:], m.Signatures[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.Signatures[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BTPHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BTPHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BTPHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NextValidators) > 0 {
		for iNdEx := len(m.NextValidators) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.NextValidators[iNdEx])
			copy(dAtA[i:], m.NextValidators[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.NextValidators[iNdEx])))
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.MessageRoot) > 0 {
		i -= len(m.MessageRoot)
		copy(dAtA[i:], m.MessageRoot)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.MessageRoot)))
		i--
		dAtA[i] = 0x4a
	}
	if m.MessageCount != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MessageCount))
		i--
		dAtA[i] = 0x40
	}
	if len(m.PrevNetworkSectionHash) > 0 {
		i -= len(m.PrevNetworkSectionHash)
		copy(dAtA[i:], m.PrevNetworkSectionHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.PrevNetworkSectionHash)))
		i--
		dAtA[i] = 0x3a
	}
	if m.UpdateNumber != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.UpdateNumber))
		i--
		dAtA[i] = 0x30
	}
	if m.NetworkId != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.NetworkId))
		i--
		dAtA[i] = 0x28
	}
	if len(m.NetworkSectionToRoot) > 0 {
		for iNdEx := len(m.NetworkSectionToRoot) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NetworkSectionToRoot[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.NextProofContextHash) > 0 {
		i -= len(m.NextProofContextHash)
		copy(dAtA[i:], m.NextProofContextHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.NextProofContextHash)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Round != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Round))
		i--
		dAtA[i] = 0x10
	}
	if m.MainHeight != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MainHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MerkleNode) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MerkleNode) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MerkleNode) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if m.Dir != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Dir))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SignedHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Signatures) > 0 {
		for _, b := range m.Signatures {
			l = len(b)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *BTPHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MainHeight != 0 {
		n += 1 + sovTypes(uint64(m.MainHeight))
	}
	if m.Round != 0 {
		n += 1 + sovTypes(uint64(m.Round))
	}
	l = len(m.NextProofContextHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.NetworkSectionToRoot) > 0 {
		for _, e := range m.NetworkSectionToRoot {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.NetworkId != 0 {
		n += 1 + sovTypes(uint64(m.NetworkId))
	}
	if m.UpdateNumber != 0 {
		n += 1 + sovTypes(uint64(m.UpdateNumber))
	}
	l = len(m.PrevNetworkSectionHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.MessageCount != 0 {
		n += 1 + sovTypes(uint64(m.MessageCount))
	}
	l = len(m.MessageRoot)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.NextValidators) > 0 {
		for _, b := range m.NextValidators {
			l = len(b)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *MerkleNode) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Dir != 0 {
		n += 1 + sovTypes(uint64(m.Dir))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SignedHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: SignedHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignedHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &BTPHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signatures", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signatures = append(m.Signatures, make([]byte, postIndex-iNdEx))
			copy(m.Signatures[len(m.Signatures)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *BTPHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: BTPHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BTPHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MainHeight", wireType)
			}
			m.MainHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MainHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Round |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextProofContextHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NextProofContextHash = append(m.NextProofContextHash[:0], dAtA[iNdEx:postIndex]...)
			if m.NextProofContextHash == nil {
				m.NextProofContextHash = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkSectionToRoot", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NetworkSectionToRoot = append(m.NetworkSectionToRoot, &MerkleNode{})
			if err := m.NetworkSectionToRoot[len(m.NetworkSectionToRoot)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkId", wireType)
			}
			m.NetworkId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NetworkId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateNumber", wireType)
			}
			m.UpdateNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdateNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevNetworkSectionHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrevNetworkSectionHash = append(m.PrevNetworkSectionHash[:0], dAtA[iNdEx:postIndex]...)
			if m.PrevNetworkSectionHash == nil {
				m.PrevNetworkSectionHash = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageCount", wireType)
			}
			m.MessageCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MessageCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageRoot = append(m.MessageRoot[:0], dAtA[iNdEx:postIndex]...)
			if m.MessageRoot == nil {
				m.MessageRoot = []byte{}
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextValidators", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NextValidators = append(m.NextValidators, make([]byte, postIndex-iNdEx))
			copy(m.NextValidators[len(m.NextValidators)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *MerkleNode) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: MerkleNode: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MerkleNode: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dir", wireType)
			}
			m.Dir = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Dir |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
