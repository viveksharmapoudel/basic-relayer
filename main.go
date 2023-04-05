package main

import (
	"context"
	"fmt"

	"github.com/icon-project/btp/chain/icon/client"
	"github.com/icon-project/btp/common/mbt"
	"github.com/icon-project/icon-bridge/cmd/iconbridge/chain/icon/types"
	"golang.org/x/crypto/sha3"

	mt "github.com/txaty/go-merkletree"
)

// "cosmossdk.io/errors"

const (
	// ENDPOINT              = "http://138.197.69.76:9000/api/v3/icon_dex"
	ENDPOINT              = "http://localhost:9082/api/v3/"
	WSS_ENDPOINT          = "wss://ctz.solidwallet.io/api/v3/icon_dex/event"
	SEND_PACKET_SIGNATURE = "Message(str,int,bytes)"
	CONTRACT_ADDRESS      = "cx0000000000000000000000000000000000000000"
	BTP_SIGNATURE         = "BTPMessage(int,int)"
)

func Sha3FIPS256(l ...[]byte) []byte {
	h := sha3.New256()
	for _, b := range l {
		h.Write(b)
	}
	var digest [32]byte
	h.Sum(digest[:0])
	return digest[:]
}

//   Typess

func main() {
	// cmd.Execute()
	// FetchEvent(1500)

	height := types.HexInt("0x1b")
	networkID := types.HexInt("0x1")
	ProveBTPMessage(height, networkID)

	// testQueryCycle()
	// testBtpMessageProof()
	// testMerkleProof()

	// op := cosmos_types.Any{
	// 	TypeUrl: "someurl",
	// 	Value: []byte("some value"),
	// }

	// b, _:= op.Marshal()
	// fmt.Printf("check this is %x\n", b)

	// cdc := MakeCodec()

}

func testQueryCycle() {

	processor := NewIconChainProcessor(ENDPOINT)

	ctx := context.Background()
	processor.QueryCycle(ctx, 10, 1)
}

func testBtpMessageProof() {

	var msgs [][]byte

	msgVals := []string{
		"0x766976656b",
		"0x736861726d6120706f7564656c206e6570616c",
		"0x6b746d20696272697a",
	}

	for _, msgVal := range msgVals {
		hxBytes := client.HexBytes(msgVal)
		msg, _ := hxBytes.Value()
		msgs = append(msgs, msg)
	}

	binaryTree, err := mbt.NewMerkleBinaryTree(mbt.Sha3Keccak256, msgs)
	if err != nil {
		fmt.Println(err)
		return
	}

	messageProofs, err := binaryTree.Proof(2, 2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("root: %x \n ", messageProofs.ProofInRight)

	// fa675dbe4bcaf91bf8f155c4cc6c10cd50956fc874809babb518dc29e15357eb

	// op := Sha3Keccak256([]byte(messageVal))

	// {50 0

	// // {fa5 0
	// 	d090304264eeee3c3562152f2dc355601b0b423a948824fd0a012c11c3fc2fb4
	// 	[] 1 4 e23299ec4ca5962e04434e6c3d9e411e3c2c0d353e9e818822bef835f0f423ac 2
	// 	 f6b057629e6ae866f026d357ebf7aaab8ceee97efce5a5caaf0dad7a58383b16 }

	// fmt.Printf("op %x \n", op)

	// message -> contruct each message merkle tree -> content
}

type testData struct {
	data []byte
}

func (t *testData) Serialize() ([]byte, error) {
	return t.data, nil
}

func testMerkleProof() {

	var msgs []mt.DataBlock

	msgVals := []string{
		"0x766976656b",
		"0x736861726d6120706f7564656c206e6570616c",
		"0x6b746d20696272697a",
		"0x766976656b",
		"0x736861726d6120706f7564656c206e6570616c",
		"0x6b746d20696272697a",
	}

	for _, msgVal := range msgVals {
		msgs = append(msgs,
			&testData{
				data: []byte(msgVal),
			},
		)
	}
	tree, err := mt.New(nil, msgs)
	if err != nil {
		fmt.Println(err)
		return
	}
	p, _ := tree.Proof(msgs[1])
	fmt.Printf("%x\n", p)
	fmt.Printf("%x", msgs[1])

}
