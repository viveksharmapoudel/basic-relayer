package main

import (
	"encoding/base64"
	"fmt"

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
