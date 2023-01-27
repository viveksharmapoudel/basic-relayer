package src

import "log"

func BalanceTransfer(address string, bal int) error {

	log.Printf("balance transfered address:%s, balance:%d \n", address, bal)
	return nil
}

func RawTransaction(d string) error {

	//function for executing raw transaction
	log.Println("data->", d)
	return nil

}
