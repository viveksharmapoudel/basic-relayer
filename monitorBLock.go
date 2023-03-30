package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/basic-relayer/icon"
	types "github.com/basic-relayer/icon"

	cosmos_codec "github.com/cosmos/cosmos-sdk/codec"
	cosmos_types "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/proto"
	btpClient "github.com/icon-project/btp/chain/icon/client"
	"github.com/icon-project/btp/common/log"
	"github.com/icon-project/goloop/common"
	"github.com/icon-project/goloop/common/codec"
	"github.com/icon-project/goloop/common/db"
	"github.com/icon-project/goloop/common/trie/ompt"

	"github.com/gorilla/websocket"
)

type IconChainProcessor struct {
	client btpClient.Client
}

func NewIconChainProcessor(url string) *IconChainProcessor {
	return &IconChainProcessor{
		client: *btpClient.NewClient(url, log.GlobalLogger()),
	}
}

type BtpHeaderInterface interface {
	proto.Message
}

func MakeCodec() cosmos_codec.ProtoCodecMarshaler {
	interfaceRegistry := cosmos_types.NewInterfaceRegistry()
	interfaceRegistry.RegisterInterface(
		"/icon.types.v1.BTPHeader",
		(*BtpHeaderInterface)(nil),
		&types.BTPHeader{},
	)
	marshaler := cosmos_codec.NewProtoCodec(interfaceRegistry)

	// moduleBasics.RegisterInterfaces(interfaceRegistry)
	return marshaler
}

func (icp *IconChainProcessor) QueryCycle(ctx context.Context, height int64, networkID int64) {
	btpMessages := make(chan []string, 1000)
	// btpBlockNotifications := make(chan *btpClient.BlockNotification, 1000)
	monitorErr := make(chan error, 1)

	req := &btpClient.BTPRequest{
		Height:    btpClient.NewHexInt(height),
		NetworkID: btpClient.NewHexInt(1),
		ProofFlag: btpClient.NewHexInt(1),
	}

	var filters []*btpClient.EventFilter
	filters = append(filters, &btpClient.EventFilter{
		// Addr: "cxfffe383e4780084e48e477935099b03193d952fe",
		Signature: "SendPacket(bytes)",
	}, &btpClient.EventFilter{
		Signature: "CreateClient(str,bytes)",
	})

	// reqIconBlock := &btpClient.BlockRequest{
	// 	Height:       btpClient.NewHexInt(height),
	// 	EventFilters: filters,
	// }

	log.Println("values->", height, "<-->", networkID)
	go icp.monitorBTP2Block(req, btpMessages, monitorErr)
	// icp.MonitorIconBlock(reqIconBlock, btpBlockNotifications, monitorErr)
	for {
		select {
		case bn := <-btpMessages:
			// Do something with the block notification
			fmt.Println("btp messages", bn)
			// rlpdecode the message
			// get the message type

			// Query for more information based on the block notification

		// case noti := <-btpBlockNotifications:
		// 	go icp.handleBlockEventRequest(noti)

		case err := <-monitorErr:
			// Handle the error
			fmt.Printf("Error received: %s\n", err.Error())
			return

		case <-ctx.Done():
			// Context has been cancelled, stop the loop
			return
		}
	}
}

func (icp *IconChainProcessor) MonitorIconBlock(req *btpClient.BlockRequest, incomingEvent chan *btpClient.BlockNotification, errChan chan error) {

	go func() {
		err := icp.client.MonitorBlock(req, func(conn *websocket.Conn, v *btpClient.BlockNotification) error {

			fmt.Printf("THis is IconMonitorBlock:-> %v\n", v)

			if len(v.Indexes) > 0 && len(v.Events) > 0 {

				incomingEvent <- v
			}

			return nil
		}, func(conn *websocket.Conn) {
			log.Println(fmt.Sprintf("ReceiveLoop monitorBTP2Block"))
		}, func(conn *websocket.Conn, err error) {
			log.Println(fmt.Sprintf("onError %s err:%+v", conn.LocalAddr().String(), err))
			_ = conn.Close()
			errChan <- err
		})
		if err != nil {
			errChan <- err
		}
	}()
}

type BlockHeaderResult struct {
	StateHash        []byte
	PatchReceiptHash []byte
	ReceiptHash      common.HexBytes
	ExtensionData    []byte
}
type TxResult struct {
	Status             int64
	To                 []byte
	CumulativeStepUsed []byte
	StepUsed           []byte
	StepPrice          []byte
	LogsBloom          []byte
	EventLogs          []btpClient.EventLog
	ScoreAddress       []byte
	EventLogsHash      common.HexBytes
	TxIndex            btpClient.HexInt
	BlockHeight        btpClient.HexInt
}

func (icp *IconChainProcessor) handleBlockEventRequest(request *btpClient.BlockNotification) {

	blockHeaderEncoded, err := icp.client.GetBlockHeaderByHeight(&btpClient.BlockHeightParam{
		Height: request.Height,
	})
	var blockHeader btpClient.BlockHeader
	_, err = codec.RLP.UnmarshalFromBytes(blockHeaderEncoded, &blockHeader)
	if err != nil {
		fmt.Println(err)
		return
	}

	var receiptHash BlockHeaderResult
	_, err = codec.RLP.UnmarshalFromBytes(blockHeader.Result, &receiptHash)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, index := range request.Indexes[0] {
		p := &btpClient.ProofEventsParam{
			Index:     index,
			BlockHash: request.Hash,
			Events:    request.Events[0][i],
		}

		proofs, err := icp.client.GetProofForEvents(p)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Processing receipt index
		serializedReceipt, err := mptProve(index, proofs[0], receiptHash.ReceiptHash)
		if err != nil {
			fmt.Println(err)
			return
		}
		var result TxResult
		_, err = codec.RLP.UnmarshalFromBytes(serializedReceipt, &result)
		if err != nil {
			// request.err = errors.Wrapf(err, "Unmarshal Receipt: %v", err)
			fmt.Println(err)
			return
		}

		fmt.Printf("Receipt: %x\n", serializedReceipt)

		for j := 0; j < len(p.Events); j++ {
			// nextEP is pointer to event where sequence has caught up
			serializedEventLog, err := mptProve(
				p.Events[j], proofs[j+1], common.HexBytes(result.EventLogsHash))
			if err != nil {
				fmt.Println(err)
				return
			}
			var el btpClient.EventLog
			_, err = codec.RLP.UnmarshalFromBytes(serializedEventLog, &el)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("this is eventLog:%x\n", el.Indexed[1])
		}

	}

}

type TempBtpBlockHeader struct {
	MainHeight             uint64
	Round                  uint64
	NextProofContextHash   btpClient.HexBytes
	NetworkSectionToRoot   []*icon.MerkleNode
	NetworkId              uint64
	UpdateNumber           uint64
	PrevNetworkSectionHash btpClient.HexBytes
	MessageCount           int64
	MessageRoot            btpClient.HexBytes
	NextValidators         []btpClient.HexBytes
}

func (icp *IconChainProcessor) monitorBTP2Block(req *btpClient.BTPRequest, msgsChan chan []string, errChan chan error) {

	go func() {
		err := icp.client.MonitorBTP(req, func(conn *websocket.Conn, v *btpClient.BTPNotification) error {
			fmt.Println("btp-header", v.Header)
			h, err := base64.StdEncoding.DecodeString(v.Header)
			if err != nil {
				return err
			}

			bh := types.BTPHeader{}
			// bh := &BtpBlockHeaderFormat{}
			if _, err = codec.RLP.UnmarshalFromBytes(h, &bh); err != nil {
				return err
			}

			var bhBTP2 TempBtpBlockHeader
			if _, err = codec.RLP.UnmarshalFromBytes(h, &bhBTP2); err != nil {
				return err
			}

			fmt.Printf("BhBTP2-- %x \n", bhBTP2)

			// encoded, err := proto.Marshal(bh)
			// if err != nil {
			// 	fmt.Println(err)
			// }

			cod := MakeCodec()
			encoded, err := MarshalJSONAny(cod, &bh)
			if err != nil {
				fmt.Println(err)
			}

			op, err := json.Marshal(encoded)
			fmt.Printf("this is from json marhsal: %x\n", op)

			jsonDumpData(bhBTP2, btpClient.NewHexBytes(encoded))
			// encoded = []byte("0x080b10011a030a141e2206080112020a14280130143a030a141e400a4a0314141e")

			// hb := btpClient.NewHexBytes(encoded)

			// base64Encode := base64.StdEncoding.EncodeToString(encoded)

			log.Printf("BTP header %x \n", encoded)
			// log.Printf("BTP header %v \n", bh)
			// var op BtpHeaderInterface
			// err = UnmarshalJSONAny(cod, &op, encoded)
			// _, err = cod.MarshalInterfaceJSON(encoded, op)
			// fmt.Println("check error: ", err)
			// fmt.Println("value check:", op)
			// fetch all the filtered events related to that block

			// msgs, err := icp.GetBtpMessage(int64(bh.MainHeight))
			// if err != nil {
			// 	return err
			// }
			// fmt.Println("messages", msgs)

			// result := make([][]byte, 0)
			// for _, mg := range msgs {
			// 	m, err := base64.StdEncoding.DecodeString(mg)
			// 	if err != nil {
			// 		fmt.Println(err)
			// 	}
			// 	result = append(result, m)
			// }

			// mt, err := mbt.NewMerkleBinaryTree(mbt.HashFuncByUID("eth"), result)
			// if err != nil {
			// 	fmt.Println(err)
			// }
			// fmt.Println("roothash of MBT:", mt.Root())

			return nil
		}, func(conn *websocket.Conn) {
			log.Println(fmt.Sprintf("ReceiveLoop monitorBTP2Block"))
		}, func(conn *websocket.Conn, err error) {
			log.Println(fmt.Sprintf("onError %s err:%+v", conn.LocalAddr().String(), err))
			_ = conn.Close()
			errChan <- err
		})

		if err != nil {
			errChan <- err
		}
	}()
}

func (icp *IconChainProcessor) GetBtpMessage(height int64) ([]string, error) {
	pr := btpClient.BTPBlockParam{
		Height:    btpClient.NewHexInt(height),
		NetworkId: btpClient.NewHexInt(1),
	}
	mgs, err := icp.client.GetBTPMessage(&pr)
	if err != nil {
		return nil, err
	}
	return mgs, nil
}

func (icp *IconChainProcessor) GetAllTheEvents(height int64) error {

	return nil
}

type HeaderANdBuf struct {
	Header          TempBtpBlockHeader `json:"header"`
	EncodedProtobuf btpClient.HexBytes `json:"encoded_protobuf"`
}

func jsonDumpData(header TempBtpBlockHeader, protobufEncoded btpClient.HexBytes) {

	filename := "headerDump.json"

	var headerAndbufs []HeaderANdBuf

	// Check if the JSON file exists
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		// Read existing JSON data from file
		jsonData, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println("Error reading JSON from file:", err)
			os.Exit(1)
		}

		// Unmarshal JSON data into a slice of structs
		err = json.Unmarshal(jsonData, &headerAndbufs)
		if err != nil {
			fmt.Println("Error unmarshaling JSON data:", err)
			os.Exit(1)
		}
	}

	fmt.Println("check******************************", header)

	newHeaderAndBuf := HeaderANdBuf{
		Header:          header,
		EncodedProtobuf: protobufEncoded,
	}

	fmt.Println("this is the header header buf ", headerAndbufs)
	// Append the new person to the people slice
	headerAndbufs = append(headerAndbufs, newHeaderAndBuf)

	// Marshal the slice of structs to JSON format
	jsonData, err := json.MarshalIndent(headerAndbufs, "", "  ")
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

func mptProve(key btpClient.HexInt, proofs [][]byte, hash []byte) ([]byte, error) {
	db := db.NewMapDB()
	defer db.Close()
	index, err := key.Value()
	if err != nil {
		return nil, err
	}
	indexKey, err := codec.RLP.MarshalToBytes(index)
	if err != nil {
		return nil, err
	}
	mpt := ompt.NewMPTForBytes(db, hash)
	trie, err1 := mpt.Prove(indexKey, proofs)
	if err1 != nil {
		return nil, err1

	}
	return trie, nil
}

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

// 0a07736f6d6575726c120a736f6d652076616c7565 hex value of let test_any= Any{
// 	type_url:"someurl".to_string(),
// 	value:"some value".as_bytes().to_vec()

//  };
