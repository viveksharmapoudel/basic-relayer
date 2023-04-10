package main

import (
	"fmt"

	"github.com/basic-relayer/icon"
	"github.com/icon-project/goloop/common/codec"
	"github.com/icon-project/goloop/common/crypto"
)

type BtpBlockHeaderFormat struct {
	MainHeight             uint64
	Round                  uint32
	NextProofContextHash   []byte
	NetworkSectionToRoot   []*icon.MerkleNode
	NetworkId              uint64
	UpdateNumber           uint64
	PrevNetworkSectionHash []byte
	MessageCount           uint64
	MessageRoot            []byte
	NextProofContext       []byte
}

type NetworkTypeSection struct {
	NextProofContextHash []byte
	NetworkSectionsRoot  []byte
}

type NetworkTypeSectionDecision struct {
	SrcNetworkID           []byte
	DstType                int64
	Height                 int64
	Round                  int32
	NetworkTypeSectionHash []byte
	// mod                    module.NetworkTypeModule
}

type NetworkSection struct {
	Nid          int64
	UpdateNumber int64
	Prev         []byte
	MessageCount int64
	MessageRoot  []byte
}

type Secp256k1Proof struct {
	Signatures []*crypto.Signature
	Bytes      []byte
}

func NewNetworkSection(
	header *BtpBlockHeaderFormat,
) *NetworkSection {
	return &NetworkSection{
		Nid:          int64(header.NetworkId),
		UpdateNumber: int64(header.UpdateNumber),
		Prev:         header.PrevNetworkSectionHash,
		MessageCount: int64(header.MessageCount),
		MessageRoot:  header.MessageRoot,
	}
}

func (h *NetworkSection) Hash() []byte {
	return Keccak256(codec.RLP.MustMarshalToBytes(h))
}

func NewNetworkTypeSectionDecision(SrcNetworkID []byte,
	DstType int64,
	Height int64,
	Round int32,
	networkTypeSection NetworkTypeSection,
) *NetworkTypeSectionDecision {
	return &NetworkTypeSectionDecision{
		SrcNetworkID,
		DstType,
		Height,
		Round,
		(networkTypeSection.Hash()),
	}
}

func (h *NetworkTypeSectionDecision) Hash() []byte {
	return Keccak256(codec.RLP.MustMarshalToBytes(h))
}

func (h *NetworkTypeSectionDecision) Encode() []byte {
	return codec.RLP.MustMarshalToBytes(h)
}

func (h NetworkTypeSection) Encode() []byte {
	return codec.RLP.MustMarshalToBytes(h)
}

func (h NetworkTypeSection) Hash() []byte {
	return Keccak256(codec.RLP.MustMarshalToBytes(h))
}

func (header *BtpBlockHeaderFormat) GetNetworkSectionRoot() []byte {
	networkSection := NewNetworkSection(header)
	base := networkSection.Hash()

	for _, root := range header.NetworkSectionToRoot {
		fmt.Println(root)
		// todo: implement
	}
	return base
}
