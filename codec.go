package main

// import (
// 	"fmt"

// 	types "github.com/basic-relayer/icon"
// 	cosmos_codec "github.com/cosmos/cosmos-sdk/codec"
// 	cosmos_types "github.com/cosmos/cosmos-sdk/codec/types"
// 	"github.com/cosmos/cosmos-sdk/std"
// 	"github.com/gogo/protobuf/proto"
// )

// func MarshalJSONAny(m cosmos_codec.Codec, msg proto.Message) ([]byte, error) {
// 	any, err := cosmos_types.NewAnyWithValue(msg)

// 	if err != nil {
// 		return nil, err
// 	}

// 	op, err := m.MarshalJSON(any)
// 	fmt.Println(op)
// 	return op, err
// }

// func UnmarshalJSONAny(m cosmos_codec.Codec, iface interface{}, bz []byte) error {
// 	any := &cosmos_types.Any{}

// 	err := m.UnmarshalJSON(bz, any)
// 	if err != nil {
// 		return err
// 	}

// 	return m.UnpackAny(any, iface)
// }

// func MakeCodec() cosmos_codec.ProtoCodecMarshaler {
// 	interfaceRegistry := cosmos_types.NewInterfaceRegistry()
// 	marshaler := cosmos_codec.NewProtoCodec(interfaceRegistry)
// 	std.RegisterInterfaces(interfaceRegistry)
// 	interfaceRegistry.RegisterInterface(
// 		"/icon.types.v1",
// 		(*BtpHeaderInterface)(nil),
// 		// &types.BTPHeader{},
// 	)
// 	interfaceRegistry.RegisterInterface(
// 		"/icon.types.v1",
// 		(*SignedHeaderInterface)(nil),
// 		&types.SignedHeader{},
// 	)
// 	interfaceRegistry.RegisterImplementations(
// 		(*BtpHeaderInterface)(nil),
// 		&types.BTPHeader{},
// 	)
// 	interfaceRegistry.RegisterImplementations(
// 		(*BtpHeaderInterface)(nil),
// 		&types.BlockUpdate{},
// 	)

// 	fmt.Println("check this out ", marshaler.InterfaceRegistry().ListImplementations("//icon.types.v1"))

// 	return marshaler
// }

// // func MakeCodec() codec.ProtoCodecMarshaler {
// // 	interfaceRegistry := types.NewInterfaceRegistry()
// // 	marshaler := codec.NewProtoCodec(interfaceRegistry)
// // 	std.RegisterInterfaces(interfaceRegistry)
// // 	moduleBasics.RegisterInterfaces(interfaceRegistry)
// // 	return marshaler

// // }
