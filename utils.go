package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/icon-project/goloop/common/codec"
	"golang.org/x/crypto/sha3"
)

func appendKeccak256(out []byte, data ...[]byte) []byte {
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(out)
}

// hash fucntion type : string
// switchCase

func Keccak256(data ...[]byte) []byte {
	return appendKeccak256(nil, data...)
}

func Sha3Keccak256(data ...[]byte) []byte {
	h := sha3.NewLegacyKeccak256()
	for _, b := range data {
		h.Write(b)
	}
	var digest [32]byte
	h.Sum(digest[:0])
	return digest[:]
}

func GetSourceNetworkUID(nid int) []byte {
	return []byte(fmt.Sprintf("0x%x.icon", nid))
}

func Base64ToData(encoded string, v interface{}) ([]byte, error) {
	if encoded == "" {
		return nil, fmt.Errorf("Encoded string is empty ")
	}

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	return codec.RLP.UnmarshalFromBytes(decoded, v)
}

func jsonDumpDataFile(filename string, bufs interface{}) {
	// Marshal the slice of structs to JSON format
	jsonData, err := json.MarshalIndent(bufs, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling slice of structs to JSON:", err)
		os.Exit(1)
	}

	// Write JSON data to file
	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully created or appended JSON in headerDump.json")
}

func readExistingData(filename string, opPointer interface{}) error {

	// Check if the JSON file exists
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		// Read existing JSON data from file
		jsonData, err := ioutil.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("Error reading JSON from file: %v", err)
		}

		// Unmarshal JSON data into a slice of structs
		err = json.Unmarshal(jsonData, opPointer)
		if err != nil {
			return fmt.Errorf("Error unmarshaling JSON data: %v", err)
		}
	}

	return nil
}
