package main

// icon_types "github.com/icon-project/ibc-integration/libraries/go/common/icon"

type IBtpHeader interface {
}

// func TestBtpHeader(t *testing.T) {

// 	interfaceRegistry := types.NewInterfaceRegistry()
// 	interfaceRegistry.RegisterInterface("icon.types.v1.HeaderI", (*IBtpHeader)(nil))
// 	interfaceRegistry.RegisterImplementations((*IBtpHeader)(nil), &icon_types.BTPHeader{})

// 	marshaler := codec.NewProtoCodec(interfaceRegistry)

// 	m := &icon_types.BTPHeader{
// 		MainHeight: 20,
// 	}

// 	fmt.Println("check this value::::")
// 	fmt.Println(marshaler)

// 	fmt.Println("check this out ", marshaler.InterfaceRegistry().ListAllInterfaces())
// 	fmt.Println("implementations: ", interfaceRegistry.ListImplementations("icon.types.v1.HeaderI"))

// 	x, _ := types.NewAnyWithValue(m)

// 	b, _ := marshaler.MarshalJSON(x)
// 	fmt.Printf("output: %s \n", b)

// }
