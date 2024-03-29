// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: icon/types/v1/types.proto

package icon

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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
	BlockIDFlag_BLOCK_ID_FLAG_UNKNOWN BlockIDFlag = 0
	BlockIDFlag_BLOCK_ID_FLAG_ABSENT  BlockIDFlag = 1
	BlockIDFlag_BLOCK_ID_FLAG_COMMIT  BlockIDFlag = 2
	BlockIDFlag_BLOCK_ID_FLAG_NIL     BlockIDFlag = 3
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
	SignedMsgType_SIGNED_MSG_TYPE_UNKNOWN SignedMsgType = 0
	// Votes
	SignedMsgType_SIGNED_MSG_TYPE_PREVOTE   SignedMsgType = 1
	SignedMsgType_SIGNED_MSG_TYPE_PRECOMMIT SignedMsgType = 2
	// Proposals
	SignedMsgType_SIGNED_MSG_TYPE_PROPOSAL SignedMsgType = 32
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

func (m *SignedHeader) GetHeader() *BTPHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SignedHeader) GetSignatures() [][]byte {
	if m != nil {
		return m.Signatures
	}
	return nil
}

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

func (m *BTPHeader) GetMainHeight() uint64 {
	if m != nil {
		return m.MainHeight
	}
	return 0
}

func (m *BTPHeader) GetRound() uint32 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *BTPHeader) GetNextProofContextHash() []byte {
	if m != nil {
		return m.NextProofContextHash
	}
	return nil
}

func (m *BTPHeader) GetNetworkSectionToRoot() []*MerkleNode {
	if m != nil {
		return m.NetworkSectionToRoot
	}
	return nil
}

func (m *BTPHeader) GetNetworkId() uint64 {
	if m != nil {
		return m.NetworkId
	}
	return 0
}

func (m *BTPHeader) GetUpdateNumber() uint64 {
	if m != nil {
		return m.UpdateNumber
	}
	return 0
}

func (m *BTPHeader) GetPrevNetworkSectionHash() []byte {
	if m != nil {
		return m.PrevNetworkSectionHash
	}
	return nil
}

func (m *BTPHeader) GetMessageCount() uint64 {
	if m != nil {
		return m.MessageCount
	}
	return 0
}

func (m *BTPHeader) GetMessageRoot() []byte {
	if m != nil {
		return m.MessageRoot
	}
	return nil
}

func (m *BTPHeader) GetNextValidators() [][]byte {
	if m != nil {
		return m.NextValidators
	}
	return nil
}

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

func (m *MerkleNode) GetDir() int32 {
	if m != nil {
		return m.Dir
	}
	return 0
}

func (m *MerkleNode) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type MerkleProofs struct {
	Proofs []*MerkleNode `protobuf:"bytes,1,rep,name=proofs,proto3" json:"proofs,omitempty"`
}

func (m *MerkleProofs) Reset()         { *m = MerkleProofs{} }
func (m *MerkleProofs) String() string { return proto.CompactTextString(m) }
func (*MerkleProofs) ProtoMessage()    {}
func (*MerkleProofs) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb5e993a7dd5d0ff, []int{3}
}
func (m *MerkleProofs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MerkleProofs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MerkleProofs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MerkleProofs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerkleProofs.Merge(m, src)
}
func (m *MerkleProofs) XXX_Size() int {
	return m.Size()
}
func (m *MerkleProofs) XXX_DiscardUnknown() {
	xxx_messageInfo_MerkleProofs.DiscardUnknown(m)
}

var xxx_messageInfo_MerkleProofs proto.InternalMessageInfo

func (m *MerkleProofs) GetProofs() []*MerkleNode {
	if m != nil {
		return m.Proofs
	}
	return nil
}

func init() {
	proto.RegisterEnum("icon.types.v1.BlockIDFlag", BlockIDFlag_name, BlockIDFlag_value)
	proto.RegisterEnum("icon.types.v1.SignedMsgType", SignedMsgType_name, SignedMsgType_value)
	proto.RegisterType((*SignedHeader)(nil), "icon.types.v1.SignedHeader")
	proto.RegisterType((*BTPHeader)(nil), "icon.types.v1.BTPHeader")
	proto.RegisterType((*MerkleNode)(nil), "icon.types.v1.MerkleNode")
	proto.RegisterType((*MerkleProofs)(nil), "icon.types.v1.MerkleProofs")
}

func init() { proto.RegisterFile("icon/types/v1/types.proto", fileDescriptor_eb5e993a7dd5d0ff) }

var fileDescriptor_eb5e993a7dd5d0ff = []byte{
	// 692 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcd, 0x6e, 0xda, 0x4a,
	0x14, 0xc6, 0x90, 0x70, 0x6f, 0x0e, 0x70, 0x2f, 0x19, 0x25, 0x37, 0x46, 0xb7, 0xa1, 0x94, 0x48,
	0x15, 0xca, 0x02, 0x4a, 0xda, 0x2e, 0x4a, 0x57, 0xfc, 0x85, 0x58, 0x01, 0x63, 0x19, 0x97, 0xfe,
	0x28, 0x92, 0x3b, 0xc1, 0x53, 0x63, 0x05, 0x3c, 0xc8, 0x33, 0xd0, 0xe6, 0x05, 0xba, 0xee, 0x33,
	0x74, 0x59, 0xf5, 0x09, 0xfa, 0x04, 0x55, 0x57, 0x59, 0x76, 0x59, 0x91, 0x5d, 0x9f, 0xa2, 0x9a,
	0x31, 0x49, 0x43, 0x12, 0xa9, 0x1b, 0x34, 0xdf, 0xf7, 0x9d, 0xef, 0xcc, 0xf9, 0x61, 0x0c, 0x19,
	0x6f, 0x40, 0xfd, 0x12, 0x3f, 0x9d, 0x10, 0x56, 0x9a, 0x95, 0xc3, 0x43, 0x71, 0x12, 0x50, 0x4e,
	0x51, 0x4a, 0x48, 0xc5, 0x90, 0x99, 0x95, 0xf3, 0xaf, 0x21, 0xd9, 0xf3, 0x5c, 0x9f, 0x38, 0x07,
	0x04, 0x3b, 0x24, 0x40, 0x0f, 0x20, 0x3e, 0x94, 0x27, 0x55, 0xc9, 0x29, 0x85, 0xc4, 0x9e, 0x5a,
	0x5c, 0x8a, 0x2f, 0xd6, 0x2c, 0x23, 0x8c, 0x34, 0x17, 0x71, 0x28, 0x0b, 0xc0, 0x3c, 0xd7, 0xc7,
	0x7c, 0x1a, 0x10, 0xa6, 0x46, 0x73, 0xb1, 0x42, 0xd2, 0xbc, 0xc2, 0xe4, 0xbf, 0xc4, 0x60, 0xed,
	0xd2, 0x85, 0xee, 0x42, 0x62, 0x8c, 0x3d, 0xdf, 0x1e, 0x12, 0xcf, 0x1d, 0x72, 0x79, 0xc9, 0x8a,
	0x09, 0x82, 0x3a, 0x90, 0x0c, 0xda, 0x80, 0xd5, 0x80, 0x4e, 0x7d, 0x47, 0x8d, 0xe6, 0x94, 0x42,
	0xca, 0x0c, 0x01, 0x7a, 0x0c, 0x5b, 0x3e, 0x79, 0xc7, 0xed, 0x49, 0x40, 0xe9, 0x1b, 0x7b, 0x40,
	0x7d, 0x2e, 0xd0, 0x10, 0xb3, 0xa1, 0x1a, 0xcb, 0x29, 0x85, 0xa4, 0xb9, 0x21, 0x64, 0x43, 0xa8,
	0xf5, 0x50, 0x3c, 0xc0, 0x6c, 0x88, 0x0c, 0x61, 0xe3, 0x6f, 0x69, 0x70, 0x62, 0x33, 0x32, 0xe0,
	0x1e, 0xf5, 0x6d, 0x4e, 0xed, 0x80, 0x52, 0xae, 0xae, 0xe4, 0x62, 0x85, 0xc4, 0x5e, 0xe6, 0x5a,
	0x7b, 0x1d, 0x12, 0x9c, 0x8c, 0x88, 0x4e, 0x1d, 0x22, 0x32, 0x4a, 0x67, 0x2f, 0x34, 0x5a, 0xd4,
	0xa4, 0x94, 0xa3, 0x6d, 0x80, 0x8b, 0x8c, 0x9e, 0xa3, 0xae, 0xca, 0xf2, 0xd7, 0x16, 0x8c, 0xe6,
	0xa0, 0x1d, 0x48, 0x4d, 0x27, 0x0e, 0xe6, 0xc4, 0xf6, 0xa7, 0xe3, 0x63, 0x12, 0xa8, 0x71, 0x19,
	0x91, 0x0c, 0x49, 0x5d, 0x72, 0xe8, 0x09, 0x64, 0x26, 0x01, 0x99, 0xd9, 0xd7, 0x4b, 0x93, 0xed,
	0xfc, 0x25, 0xdb, 0xf9, 0x4f, 0x04, 0xe8, 0x4b, 0x05, 0xc8, 0x86, 0x76, 0x20, 0x35, 0x26, 0x8c,
	0x61, 0x97, 0xd8, 0x03, 0x3a, 0xf5, 0xb9, 0xfa, 0x77, 0x98, 0x7f, 0x41, 0xd6, 0x05, 0x87, 0xee,
	0xc1, 0x05, 0x0e, 0x5b, 0x5d, 0x93, 0x29, 0x13, 0x0b, 0x4e, 0xb6, 0x71, 0x1f, 0xfe, 0x11, 0x03,
	0xeb, 0xe3, 0x91, 0xe7, 0x60, 0x4e, 0x03, 0xa6, 0x82, 0x5c, 0xdc, 0x35, 0x36, 0xff, 0x08, 0xe0,
	0xf7, 0x48, 0x50, 0x1a, 0x62, 0x0d, 0x2f, 0xfc, 0x67, 0xac, 0x9a, 0xe2, 0x28, 0xb6, 0x35, 0xc3,
	0xa3, 0x29, 0x91, 0xdb, 0x4a, 0x9a, 0x21, 0xc8, 0x57, 0x21, 0x19, 0xba, 0xe4, 0x42, 0x18, 0x2a,
	0x43, 0x5c, 0x2e, 0x8e, 0xa9, 0xca, 0x9f, 0xa6, 0xbe, 0x08, 0xdc, 0x65, 0x90, 0xa8, 0x8d, 0xe8,
	0xe0, 0x44, 0x6b, 0xec, 0x8f, 0xb0, 0x8b, 0x32, 0xb0, 0x59, 0x6b, 0x77, 0xeb, 0x87, 0xb6, 0xd6,
	0xb0, 0xf7, 0xdb, 0xd5, 0x96, 0xfd, 0x4c, 0x3f, 0xd4, 0xbb, 0xcf, 0xf5, 0x74, 0x04, 0xa9, 0xb0,
	0xb1, 0x2c, 0x55, 0x6b, 0xbd, 0xa6, 0x6e, 0xa5, 0x95, 0x9b, 0x4a, 0xbd, 0xdb, 0xe9, 0x68, 0x56,
	0x3a, 0x8a, 0x36, 0x61, 0x7d, 0x59, 0xd1, 0xb5, 0x76, 0x3a, 0xb6, 0xfb, 0x5e, 0x81, 0x54, 0xf8,
	0x1a, 0x3a, 0xcc, 0xb5, 0x4e, 0x27, 0x04, 0xfd, 0x0f, 0x5b, 0x3d, 0xad, 0xa5, 0x37, 0x1b, 0x76,
	0xa7, 0xd7, 0xb2, 0xad, 0x97, 0x46, 0xf3, 0xca, 0xcd, 0xb7, 0x88, 0x86, 0xd9, 0xec, 0x77, 0xad,
	0x66, 0x5a, 0x41, 0xdb, 0x90, 0xb9, 0x45, 0xbc, 0xac, 0xe0, 0x0e, 0xa8, 0x37, 0xe5, 0xae, 0xd1,
	0xed, 0x55, 0xdb, 0xe9, 0x5c, 0xed, 0xb3, 0xf2, 0x75, 0x9e, 0x55, 0xce, 0xe6, 0x59, 0xe5, 0xc7,
	0x3c, 0xab, 0x7c, 0x38, 0xcf, 0x46, 0xce, 0xce, 0xb3, 0x91, 0xef, 0xe7, 0xd9, 0x08, 0xac, 0x0f,
	0xe8, 0x78, 0x79, 0x7c, 0x35, 0x10, 0xa5, 0x32, 0x43, 0x3c, 0x6f, 0x43, 0x79, 0xb5, 0x3d, 0xf2,
	0x8e, 0x03, 0x1c, 0x78, 0x84, 0x95, 0x5c, 0x5a, 0x1a, 0xd0, 0xf1, 0x98, 0xfa, 0x25, 0x61, 0x78,
	0x2a, 0x7e, 0x3e, 0x46, 0x63, 0x9a, 0xf5, 0xe2, 0x53, 0x34, 0xa5, 0x89, 0x14, 0xd2, 0x58, 0xec,
	0x97, 0xbf, 0x85, 0xf8, 0x48, 0xe2, 0xa3, 0x7e, 0x79, 0x1e, 0xcd, 0x2c, 0xe1, 0xa3, 0x96, 0x51,
	0xeb, 0x10, 0x8e, 0x1d, 0xcc, 0xf1, 0xcf, 0xe8, 0xbf, 0x42, 0xab, 0x54, 0xa4, 0x58, 0xa9, 0xf4,
	0xcb, 0xc7, 0x71, 0xf9, 0x69, 0x79, 0xf8, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x8b, 0x16, 0x4e,
	0x77, 0x04, 0x00, 0x00,
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

func (m *MerkleProofs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MerkleProofs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MerkleProofs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proofs) > 0 {
		for iNdEx := len(m.Proofs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Proofs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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

func (m *MerkleProofs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Proofs) > 0 {
		for _, e := range m.Proofs {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
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
func (m *MerkleProofs) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MerkleProofs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MerkleProofs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proofs", wireType)
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
			m.Proofs = append(m.Proofs, &MerkleNode{})
			if err := m.Proofs[len(m.Proofs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
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
