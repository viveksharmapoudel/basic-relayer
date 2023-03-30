package main

// import (
// 	"fmt"

// 	"github.com/icon-project/goloop/common"
// 	"github.com/icon-project/icon-bridge/cmd/iconbridge/chain/icon/types"
// 	"github.com/icon-project/icon-bridge/common/log"
// )

// func ProveBTPMessage(height, networkID types.HexInt) error {

// 	// initalize new client
// 	cl := NewClient(ENDPOINT)

// 	// btp header
// 	header, err := cl.GetBtpHeader(height, networkID)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	// fmt.Printf("PrevNetowkHash: %x\n", header.PrevNetworkSectionHash)

// 	btpproof, err := cl.GetBTPProof(height, networkID)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}

// 	decision := NewNetworkTypeSectionDecision(
// 		GetSourceNetworkUID(3),
// 		header.NetworkID,
// 		header.MainHeight,
// 		header.Round,
// 		NetworkTypeSection{
// 			NextProofContextHash: header.NextProofContextHash,
// 			NetworkSectionsRoot:  header.GetNetworkSectionRoot(),
// 		},
// 	)
// 	fmt.Printf("current netowrkSectionRoot: %x\n ", header.GetNetworkSectionRoot())

// 	validatorsList, err := cl.GetValidatorSet(height)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}

// 	fmt.Println("validatorList", validatorsList[0].String())

// 	verified, err := VerifyBtpProof(decision, btpproof, validatorsList)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}

// 	fmt.Println("is Verified", verified)
// 	return nil
// }

// func VerifyBtpProof(decision *NetworkTypeSectionDecision, proof *Secp256k1Proof, listValidators []common.Address) (bool, error) {

// 	if len(proof.Signatures) == 0 {
// 		return false, fmt.Errorf("signatures is empty")
// 	}

// 	requiredVotes := (2 * len(listValidators)) / 3
// 	if requiredVotes < 1 {
// 		requiredVotes = 1
// 	}

// 	fmt.Println("required votes:", requiredVotes)

// 	numVotes := 0
// 	validators := make(map[common.Address]struct{})
// 	for _, val := range listValidators {
// 		validators[val] = struct{}{}
// 	}

// 	for _, sig := range proof.Signatures {
// 		pubkey, err := sig.RecoverPublicKey(decision.Hash())
// 		if err != nil {
// 			fmt.Println("this is errror ", err)
// 			continue
// 		}

// 		address := common.NewAccountAddressFromPublicKey(pubkey)
// 		if address == nil {
// 			continue
// 		}
// 		fmt.Println("address", address)
// 		if _, ok := validators[*address]; !ok {
// 			continue
// 		}
// 		fmt.Println("add val")
// 		delete(validators, *address)
// 		if numVotes++; numVotes >= requiredVotes {
// 			return true, nil
// 		}
// 	}

// 	return false, nil

// }
