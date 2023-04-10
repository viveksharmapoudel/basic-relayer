package main

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	icon_crypto "github.com/icon-project/goloop/common/crypto"
	"golang.org/x/crypto/sha3"
	"gotest.tools/assert"
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

	signature, _ := hex.DecodeString("6c8b2bc2c3d31e34bd4ed9db6eff7d5dc647b13c58ae77d54e0b05141cb7a7995102587f1fa33fd56815463c6b78e100217c29ddca20fcace80510e3dab03a1600")
	decision, _ := hex.DecodeString("6f078f51ea67ddd7e1f288a1e2e012b2ebec11d14f9c2ec7620c90ac044dcb21")
	expectedAddress, _ := hex.DecodeString("00b040bff300eee91f7665ac8dcf89eb0871015306")
	fmt.Println("expected", expectedAddress)

	ecdsaPubkey, err := crypto.SigToPub(decision, signature)
	if err != nil {
		panic(err)
	}

	pubKeyBytes := crypto.FromECDSAPub(ecdsaPubkey)
	hex := hex.EncodeToString(pubKeyBytes)

	fmt.Println("this is the public key ", hex)

	address := crypto.PubkeyToAddress(*ecdsaPubkey)

	fmt.Println("Ethereum Address: ", address.Hex())
}

func TestPubkeyToAddress(t *testing.T) {

	signature, _ := hex.DecodeString("6c8b2bc2c3d31e34bd4ed9db6eff7d5dc647b13c58ae77d54e0b05141cb7a7995102587f1fa33fd56815463c6b78e100217c29ddca20fcace80510e3dab03a1600")
	decision, _ := hex.DecodeString("6f078f51ea67ddd7e1f288a1e2e012b2ebec11d14f9c2ec7620c90ac044dcb21")
	// expectedAddress, _ := hex.DecodeString("00b040bff300eee91f7665ac8dcf89eb0871015306")

	sig, _ := icon_crypto.ParseSignature(signature)
	pubkey, err := sig.RecoverPublicKey(decision)

	if err != nil {
		fmt.Println("error when generating pubkey", err)
	}

	addr, _ := PubkeyToAddress(pubkey)
	fmt.Printf("public key %x \n", *addr)
}

func TestGetValidatorByContext(t *testing.T) {

	addresses, err := ValidatorsByProofContext(27, 1)
	if err != nil {
		fmt.Println("error fetching address", err)
		return
	}
	fmt.Printf("check the address %x \n", addresses)
}

func TestKeccak256Hash(t *testing.T) {
	expectedHash := "06b831d75b670298d5f5a48747bb73e9c935547de89bb8a04904ec6ebc1b5f57"
	// 	expectedAddress := "5c42b6096c4601ceabacdb471cb1cdfe6bc46586"
	// 	expectedSignature := "c8b2b5eeb7b54620a0246b2355e42ce6d3bdf1648cd8eae298ebbbe5c3bacc197d5e8bfddb0f1e33778b7fc558c54d35e47c88daa24fff243aa743088e5503d701"

	msg := []byte("test message")
	hash := Keccak256(msg)
	expectedHashByte, _ := hex.DecodeString(expectedHash)

	// address, _ := hex.DecodeString(expectedAddress)

	// signature, _ := hex.DecodeString(expectedSignature)
	// r := signature[:32]
	// s := signature[32:64]
	// v := uint8(signature[64])

	assert.Equal(t, expectedHashByte, hash)
	// assert.Equal(t, address, crypto.PubkeyToAddress(crypto.SignatureToPublicKey(signature, msg)).Bytes())
	// assert.True(t, crypto.VerifySignature(crypto.PubkeyToAddress(address), crypto.Keccak256Hash(msg).Bytes(), r, s, v))
}
