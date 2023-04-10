package main

import (
	"context"
	"fmt"

	"github.com/basic-relayer/icon"
	types "github.com/basic-relayer/icon"

	"github.com/gogo/protobuf/proto"
	btpClient "github.com/icon-project/btp/chain/icon/client"
	"github.com/icon-project/btp/common/log"
	"github.com/icon-project/goloop/common"
	"github.com/icon-project/goloop/common/codec"
	"github.com/icon-project/goloop/common/db"
	"github.com/icon-project/goloop/common/trie/ompt"

	"github.com/gorilla/websocket"
	icon_bridge_types "github.com/icon-project/icon-bridge/cmd/iconbridge/chain/icon/types"
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

type SignedHeaderInterface interface {
	proto.Message
	Check()
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

func (icp *IconChainProcessor) monitorBTP2Block(req *btpClient.BTPRequest, msgsChan chan []string, errChan chan error) {

	go func() {
		err := icp.client.MonitorBTP(req, func(conn *websocket.Conn, v *btpClient.BTPNotification) error {

			var bh BtpBlockHeaderFormat
			_, err := Base64ToData(v.Header, &bh)
			if err != nil {
				log.Println("issue in one ", err)
				panic(err)
			}

			var btpproof Secp256k1Proof
			_, err = Base64ToData(v.Proof, &btpproof)
			if err != nil {
				log.Println("issue in two", err)
				return err
			}

			// decision := NewNetworkTypeSectionDecision(
			// 	GetSourceNetworkUID(3),
			// 	int64(bh.NetworkId),
			// 	int64(bh.MainHeight),
			// 	int32(bh.Round),
			// 	NetworkTypeSection{
			// 		NextProofContextHash: bh.NextProofContextHash,
			// 		NetworkSectionsRoot:  (bh.GetNetworkSectionRoot()),
			// 	},
			// )

			var signatures [][]byte
			for _, s := range btpproof.Signatures {
				b, _ := icon_bridge_types.HexBytes(s.String()).Value()
				fmt.Println("check the value ")
				signatures = append(signatures, b)
			}

			// fetch Validator
			validators, err := ValidatorsByProofContext(int64(bh.MainHeight), int64(bh.NetworkId))
			if err != nil {
				log.Println("issue in three", err)

				return err
			}

			var validatorHex [][]byte
			for _, v := range validators {
				validatorHex = append(validatorHex, v.Bytes())
			}

			signedHeader := icon.SignedHeader{
				Header: &types.BTPHeader{
					MainHeight:             bh.MainHeight,
					Round:                  bh.Round,
					NextProofContextHash:   bh.NextProofContextHash,
					NetworkSectionToRoot:   bh.NetworkSectionToRoot,
					NetworkId:              bh.NetworkId,
					UpdateNumber:           bh.UpdateNumber,
					PrevNetworkSectionHash: bh.PrevNetworkSectionHash,
					MessageCount:           bh.MessageCount,
					MessageRoot:            bh.MessageRoot,
					NextValidators:         validatorHex,
				},
				Signatures: signatures,
			}

			// yes, err := VerifyBtpProof(decision, &btpproof, validatorAddressList)
			// if err != nil {
			// 	fmt.Println("error when verifying the prooof")
			// 	return nil
			// }
			// fmt.Println(" it is verified:", yes)

			// cod := MakeCodec()
			// encoded, err := MarshalJSONAny(cod, &signedHeader)
			// if err != nil {
			// 	fmt.Println(err)
			// }

			// var decodedHeader SignedHeaderInterface
			// err = UnmarshalJSONAny(cod, &decodedHeader, encoded)
			// if err != nil {
			// 	fmt.Println(err)
			// }

			encoded, err := proto.Marshal(&signedHeader)
			if err != nil {
				fmt.Println("there is some error ")
				return err
			}
			saveSignedHeader(signedHeader, btpClient.NewHexBytes(encoded))
			// encoded = []byte("0x080b10011a030a141e2206080112020a14280130143a030a141e400a4a0314141e")

			// hb := btpClient.NewHexBytes(encoded)

			// base64Encode := base64.StdEncoding.EncodeToString(encoded)

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

type BTPQueryParam struct {
	Height btpClient.HexInt `json:"height,omitempty" validate:"optional,t_int"`
	Id     btpClient.HexInt `json:"id" validate:"required,t_int"`
}

type BTPNetworkTypeInfo struct {
	NetworkTypeName  string             `json:"networkTypeName"`
	NextProofContext string             `json:"nextProofContext"`
	OpenNetworkIDs   []btpClient.HexInt `json:"openNetworkIDs"`
	NetworkTypeID    btpClient.HexInt   `json:"networkTypeID"`
}

func (icp *IconChainProcessor) GetNetworkTypeInfo(height int64, networkId int64) (*BTPNetworkTypeInfo, error) {
	nti := &BTPNetworkTypeInfo{}
	if _, err := icp.client.Do("btp_getNetworkTypeInfo", &BTPQueryParam{
		Height: btpClient.NewHexInt(height),
		Id:     btpClient.NewHexInt(networkId),
	}, nti); err != nil {
		return nil, err
	}
	return nti, nil
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

type TempBtpBlockHeader struct {
	MainHeight             uint64
	Round                  uint32
	NextProofContextHash   icon_bridge_types.HexBytes
	NetworkSectionToRoot   []*icon.MerkleNode
	NetworkId              uint64
	UpdateNumber           uint64
	PrevNetworkSectionHash icon_bridge_types.HexBytes
	MessageCount           uint64
	MessageRoot            icon_bridge_types.HexBytes
	NextValidators         []icon_bridge_types.HexBytes
}

type TempSignedHeader struct {
	BTPHeader TempBtpBlockHeader
	Signature []icon_bridge_types.HexBytes
}

type HeaderANdBuf struct {
	Header          TempSignedHeader   `json:"signed_header"`
	EncodedProtobuf btpClient.HexBytes `json:"encoded_protobuf"`
}

func saveSignedHeader(signedHeader icon.SignedHeader, protobufEncoded btpClient.HexBytes) {

	filename := "headerDump.json"
	var headerAndbufs []HeaderANdBuf
	err := readExistingData(filename, &headerAndbufs)
	if err != nil {
		fmt.Println("there is some error while reading the data ")
		return
	}

	var validators []icon_bridge_types.HexBytes
	for _, v := range signedHeader.Header.NextValidators {
		validators = append(validators, icon_bridge_types.NewHexBytes(v))
	}

	var sig []icon_bridge_types.HexBytes
	for _, s := range signedHeader.Signatures {
		sig = append(sig, icon_bridge_types.NewHexBytes(s))
	}

	tempHeader := TempBtpBlockHeader{
		MainHeight:             signedHeader.Header.MainHeight,
		Round:                  signedHeader.Header.Round,
		NextProofContextHash:   icon_bridge_types.NewHexBytes(signedHeader.Header.NextProofContextHash),
		NetworkSectionToRoot:   signedHeader.Header.NetworkSectionToRoot,
		NetworkId:              signedHeader.Header.NetworkId,
		UpdateNumber:           signedHeader.Header.UpdateNumber,
		PrevNetworkSectionHash: icon_bridge_types.NewHexBytes(signedHeader.Header.PrevNetworkSectionHash),
		MessageCount:           signedHeader.Header.MessageCount,
		MessageRoot:            icon_bridge_types.NewHexBytes(signedHeader.Header.MessageRoot),
		NextValidators:         validators,
	}

	tempSignedHeader := TempSignedHeader{
		BTPHeader: tempHeader,
		Signature: sig,
	}

	newHeaderAndBuf := HeaderANdBuf{
		Header:          tempSignedHeader,
		EncodedProtobuf: protobufEncoded,
	}

	// Append the new person to the people slice
	headerAndbufs = append(headerAndbufs, newHeaderAndBuf)

	jsonDumpDataFile(filename, headerAndbufs)
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
