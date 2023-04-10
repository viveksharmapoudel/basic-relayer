package main

import (
	"fmt"
	"testing"

	types "github.com/basic-relayer/icon"
)

type IBtpHeader interface {
}

func TestBtpHeader(t *testing.T) {

	codec := MakeCodec()

	interfaces := codec.InterfaceRegistry().ListAllInterfaces()

	if len(interfaces) == 0 {
		fmt.Println("no interface registered")
		t.Fail()
	}

	impls := codec.InterfaceRegistry().ListImplementations(interfaces[0])

	if len(impls) == 0 {
		fmt.Println("no implementation registered error")
		t.Fail()
	}

	header := types.BTPHeader{
		MainHeight: 20,
	}

	encodedByte, err := MarshalJSONAny(codec, &header)
	if err != nil {
		fmt.Println("error getting any encode", encodedByte)
		t.Fail()
	}

	fmt.Printf("output %s \n", encodedByte)

}
