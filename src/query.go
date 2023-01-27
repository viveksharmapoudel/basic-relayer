package src

import (
	"log"
)

// function to getBalance of an account
func GetBalance(address string) int {

	//getBalance logic
	log.Printf("Balance of account:%s is %d\n", address, 100)
	return 100
}

func GetBlockHash(blockNumber int) string {
	log.Printf("The Blockhash of block_number:%d is %s", blockNumber, "sdiiuidsjdskdsjdsljsdj")
	return ""
}
