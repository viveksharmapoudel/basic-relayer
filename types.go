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
	NextProofContextHash   []byte
	NetworkSectionToRoot   []*icon.MerkleNode
	NetworkId              uint64
	UpdateNumber           uint64
	PrevNetworkSectionHash []byte
	MessageCount           uint64
	MessageRoot            []byte
	NextProofContext       []byte
}

type Validators []types.HexBytes

type NetworkTypeSection struct {
	NextProofContextHash types.HexBytes
	NetworkSectionsRoot  types.HexBytes
}

type NetworkTypeSectionDecision struct {
	SrcNetworkID           types.HexBytes
	DstType                int64
	Height                 int64
	Round                  int32
	NetworkTypeSectionHash types.HexBytes
	// mod                    module.NetworkTypeModule
}

type NetworkSection struct {
	Nid          types.HexInt
	UpdateNumber types.HexInt
	Prev         types.HexBytes
	MessageCount types.HexInt
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
		Nid:          types.NewHexInt(int64(header.NetworkId)),
		UpdateNumber: types.NewHexInt(int64(header.UpdateNumber)),
		// Prev:         header.PrevNetworkSectionHash,
		MessageCount: types.NewHexInt(int64(header.MessageCount)),
		// MessageRoot:  header.MessageRoot,
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
		types.NewHexBytes(SrcNetworkID),
		DstType,
		Height,
		Round,
		types.NewHexBytes(networkTypeSection.Hash()),
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
