package main

import (
	"fmt"

	"github.com/icon-project/btp/common/log"
	"github.com/icon-project/goloop/common"
	"github.com/icon-project/goloop/common/crypto"
	"github.com/icon-project/icon-bridge/cmd/iconbridge/chain/icon/types"
)

func ProveBTPMessage(height, networkID types.HexInt) error {

	// initalize new client
	cl := NewClient(ENDPOINT)

	// btp header
	header, err := cl.GetBtpHeader(height, networkID)
	if err != nil {
		log.Println(err)
		return err
	}
	// fmt.Printf("PrevNetowkHash: %x\n", header.PrevNetworkSectionHash)

	btpproof, err := cl.GetBTPProof(height, networkID)
	if err != nil {
		log.Println(err)
		return err
	}

	networkTypeSection := NetworkTypeSection{
		NextProofContextHash: header.NextProofContextHash,
		NetworkSectionsRoot:  header.GetNetworkSectionRoot(),
	}

	// fmt.Printf("Network Section is %x \n ", header.GetNetworkSectionRoot())

	decision := NewNetworkTypeSectionDecision(
		GetSourceNetworkUID(3),
		int64(header.NetworkId),
		int64(header.MainHeight),
		int32(header.Round),
		networkTypeSection,
	)

	// fmt.Printf("network type Section encode %x \n", networkTypeSection.Encode())
	// fmt.Printf("network type Section hash %x \n", networkTypeSection.Hash())
	//
	// fmt.Printf("networkSectionDecision Encode is: %x \n", decision.Encode())
	// fmt.Printf("networkSectionDecision Hash is: %x \n", decision.Hash())

	validatorsList, err := ValidatorsByProofContext(int64(header.MainHeight), int64(header.NetworkId))
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Printf("validatorList %x \n", validatorsList)

	verified, err := VerifyBtpProof(decision, btpproof, validatorsList)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println("is Verified", verified)
	return nil
}

func ValidatorsByProofContext(height int64, networkId int64) ([]common.Address, error) {

	chainProcessor := NewIconChainProcessor(ENDPOINT)

	typeInfo, err := chainProcessor.GetNetworkTypeInfo(height, networkId)
	if err != nil {
		return nil, err
	}

	type ValidatorList struct {
		Validators []common.Address
	}
	var vals ValidatorList
	_, err = Base64ToData(typeInfo.NextProofContext, &vals)
	if err != nil {
		return nil, err
	}

	return vals.Validators, nil
}

func PubkeyToAddress(pubKey *crypto.PublicKey) (*common.Address, error) {
	pubBytes := pubKey.SerializeUncompressed()
	return common.NewAddress(Keccak256(pubBytes[1:])[12:])
}

func VerifyBtpProof(decision *NetworkTypeSectionDecision, proof *Secp256k1Proof, listValidators []common.Address) (bool, error) {

	if len(proof.Signatures) == 0 {
		return false, fmt.Errorf("signatures is empty")
	}

	requiredVotes := (2 * len(listValidators)) / 3
	if requiredVotes < 1 {
		requiredVotes = 1
	}

	// fmt.Println("required votes:", requiredVotes)

	numVotes := 0
	validators := make(map[common.Address]struct{})
	for _, val := range listValidators {
		validators[val] = struct{}{}
	}

	fmt.Printf("%x\n ", decision.Hash())
	for _, sig := range proof.Signatures {

		pubkey, err := sig.RecoverPublicKey(decision.Hash())
		if err != nil {
			fmt.Println("error when generating pubkey", err)
			continue
		}

		address, err := PubkeyToAddress(pubkey)
		if err != nil {
			fmt.Errorf("error %v \n ", err)
			continue
		}
		if address == nil {
			continue
		}
		fmt.Println("address", address)
		if _, ok := validators[*address]; !ok {
			continue
		}
		fmt.Println("add val")
		delete(validators, *address)
		if numVotes++; numVotes >= requiredVotes {
			return true, nil
		}
	}

	return false, nil

}
