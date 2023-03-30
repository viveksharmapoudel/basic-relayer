package main

import (
	"fmt"

	"github.com/icon-project/goloop/common/codec"
	"github.com/icon-project/goloop/common/crypto"
	"github.com/icon-project/goloop/module"
)

type BtpBlockHeaderFormat struct {
	MainHeight             int64
	Round                  int32
	NextProofContextHash   []byte
	NetworkSectionToRoot   []module.MerkleNode
	NetworkID              int64
	UpdateNumber           int64
	PrevNetworkSectionHash []byte
	MessageCount           int64
	MessagesRoot           []byte
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
	bytes      []byte
}

func NewNetworkSection(
	header *BtpBlockHeaderFormat,
) *NetworkSection {

	return &NetworkSection{
		Nid:          header.NetworkID,
		UpdateNumber: header.UpdateNumber,
		Prev:         header.PrevNetworkSectionHash,
		MessageCount: header.MessageCount,
		MessageRoot:  header.MessagesRoot,
	}
}

func (h *NetworkSection) Hash() []byte {
	return Keccak256(codec.BC.MustMarshalToBytes(h))
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
		networkTypeSection.Hash(),
	}
}

func (h *NetworkTypeSectionDecision) Hash() []byte {
	return Keccak256(codec.BC.MustMarshalToBytes(h))
}

func (h *NetworkTypeSection) Hash() []byte {
	return Keccak256(codec.BC.MustMarshalToBytes(h))
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
