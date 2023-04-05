package main

import (
	"fmt"

	"github.com/basic-relayer/icon"
	"github.com/icon-project/goloop/common/codec"
	"github.com/icon-project/goloop/common/crypto"
	"github.com/icon-project/icon-bridge/cmd/iconbridge/chain/icon/types"
)

type BtpBlockHeaderFormat struct {
	MainHeight             uint64
	Round                  uint32
	NextProofContextHash   types.HexBytes
	NetworkSectionToRoot   []*icon.MerkleNode
	NetworkId              uint64
	UpdateNumber           uint64
	PrevNetworkSectionHash types.HexBytes
	MessageCount           uint64
	MessageRoot            types.HexBytes
	NextProofContext       []byte
}

type NetworkTypeSection struct {
	NextProofContextHash types.HexBytes
	NetworkSectionsRoot  types.HexBytes
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
	Prev         types.HexBytes
	MessageCount int64
	MessageRoot  types.HexBytes
}

type Secp256k1Proof struct {
	Signatures []*crypto.Signature
	bytes      []byte
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

func (h *NetworkTypeSection) Hash() []byte {
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
