package main

import (
	"encoding/hex"
	"fmt"
	"testing"

	types "github.com/basic-relayer/icon"
)

func TestNetworkDecisionCalculation(t *testing.T) {

	x := types.BTPHeader{
		MainHeight:           27,
		Round:                0,
		NextProofContextHash: []byte("d090304264eeee3c3562152f2dc355601b0b423a948824fd0a012c11c3fc2fb4"),
		// NetworkSectionToRoot: ,
		NetworkId:              1,
		UpdateNumber:           0,
		PrevNetworkSectionHash: []byte("b791b4b069c561ca31093f825f083f6cc3c8e5ad5135625becd2ff77a8ccfa1e"),
		MessageCount:           1,
		MessageRoot:            []byte("84d8e19eb09626e4a94212d3a9db54bc16a75dfd791858c0fab3032b944f657a"),
		// NextValidators: ["0xb040bff300eee91f7665ac8dcf89eb0871015306"]
	}

	proofContext, _ := hex.DecodeString(string(x.NextProofContextHash))
	prev, _ := hex.DecodeString(string(x.PrevNetworkSectionHash))

	networkSection := NetworkTypeSection{
		NextProofContextHash: proofContext,
		NetworkSectionsRoot:  prev,
	}

	networkSectionDecision := NewNetworkTypeSectionDecision(
		GetSourceNetworkUID(3),
		1,
		int64(x.MainHeight),
		int32(x.Round),
		networkSection,
	)

	fmt.Printf("network type Section encode %x \n", networkSection.Encode())
	fmt.Printf("network type Section hash %x \n", networkSection.Hash())

	fmt.Printf("networkSectionDecision Encode is: %x \n", networkSectionDecision.Encode())
	fmt.Printf("networkSectionDecision Hash is: %x \n", networkSectionDecision.Hash())

}
