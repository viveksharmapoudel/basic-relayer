package main

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func TestValidatorGet(t *testing.T) {

	processor := NewIconChainProcessor(ENDPOINT)

	op, _ := processor.GetNetworkTypeInfo(6499860, 01)
	fmt.Println(op)

	type secp256k1ProofContext struct {
		Validators []common.Address
	}
	var validators secp256k1ProofContext
	_, err := Base64ToData(string(op.NextProofContext), &validators)
	if err != nil {
		fmt.Println("error occured", err)
	}

	fmt.Printf("validators %x \n ", validators)

}

func TestRecovery(t *testing.T) {
	// Replace with your Ethereum address.
	ethAddress := "d6d594b040bff300eee91f7665ac8dcf89eb0871015306"

	// Remove the '0x' prefix and convert the address to bytes.
	addressBytes := common.Hex2Bytes(ethAddress)

	// Calculate the hash using SHA3-160 (Keccak-1600).
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(addressBytes)
	hashed := hasher.Sum(nil)

	// Add the ICON address prefix "hx".
	iconAddress := "hx" + fmt.Sprintf("%x", hashed)

	fmt.Println("Ethereum address:", ethAddress)
	fmt.Println("ICON address:", iconAddress)
}

func TestCheckEthPublicKeytoAddress(t *testing.T) {

	// signature, _ := hex.DecodeString("6c8b2bc2c3d31e34bd4ed9db6eff7d5dc647b13c58ae77d54e0b05141cb7a7995102587f1fa33fd56815463c6b78e100217c29ddca20fcace80510e3dab03a1600")
	pubkey, _ := hex.DecodeString("02c929887fe727f07a6d437d61bdf7803d9e2f10ae58c717da312fd135e20927ed")

	curve := crypto.S256()
	pubKey, err := crypto.UnmarshalPubkey(pubkey)
	if err != nil {
		panic(err)
	}

	// Verify the curve of the public key
	if pubKey.Curve != curve {
		panic("Invalid curve")
	}

	// Derive the Ethereum address from the public key
	address := crypto.PubkeyToAddress(*pubKey)

	fmt.Println("Ethereum Address:", address.Hex())
}
