package main

import (
	types "github.com/basic-relayer/icon"
	cosmos_codec "github.com/cosmos/cosmos-sdk/codec"
	cosmos_types "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/proto"
)

func MarshalJSONAny(m cosmos_codec.Codec, msg proto.Message) ([]byte, error) {
	any, err := cosmos_types.NewAnyWithValue(msg)
	if err != nil {
		return nil, err
	}
	return m.MarshalJSON(any)
}

func UnmarshalJSONAny(m cosmos_codec.Codec, iface interface{}, bz []byte) error {
	any := &cosmos_types.Any{}

	err := m.UnmarshalJSON(bz, any)
	if err != nil {
		return err
	}

	return m.UnpackAny(any, iface)
}

func MakeCodec() cosmos_codec.ProtoCodecMarshaler {
	interfaceRegistry := cosmos_types.NewInterfaceRegistry()
	interfaceRegistry.RegisterInterface(
		"/icon.types.v1.BTPHeader",
		(*BtpHeaderInterface)(nil),
		&types.BTPHeader{},
	)
	interfaceRegistry.RegisterInterface(
		"/icon.types.v1.SignedHeader",
		(*SignedHeaderInterface)(nil),
		&types.SignedHeader{},
	)
	marshaler := cosmos_codec.NewProtoCodec(interfaceRegistry)

	// moduleBasics.RegisterInterfaces(interfaceRegistry)
	return marshaler
}
