package main

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"github.com/gorilla/websocket"
// 	"github.com/icon-project/btp/chain/icon"
// 	"github.com/icon-project/btp/common/log"
// 	"github.com/icon-project/icon-bridge/cmd/iconbridge/chain/icon/types"
// )

// func FetchEvent(height int) {
// 	urls := []string{

// 		"http://170.187.251.26:9082/api/v3/default",
// 	}
// 	l := log.New()
// 	ctx := context.Background()
// 	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
// 	defer cancel()
// 	seq := 0

// 	// dstAddr := "btp://0x38.bsc/0x0401343c50FF963A7b02b20B31fA9B0B159354D4"
// 	blockReq := &types.BlockRequest{
// 		EventFilters: []*types.EventFilter{{
// 			Addr:      types.Address(CONTRACT_ADDRESS),
// 			Signature: BTP_SIGNATURE,
// 			// Indexed:   []*string{&dstAddr},
// 		}},
// 		Height: types.NewHexInt(int64(height)),
// 	}

// 	for i, url := range urls {
// 		go func(i int, url string) {
// 			l := l.WithFields(log.Fields{"i": i, "url": url})

// 			cl := icon.NewClient(url, l)

// 			h, s := height, seq
// 			err := cl.MonitorBlock(ctx, blockReq,
// 				func(conn *websocket.Conn, v *types.BlockNotification) error {
// 					_h, _ := v.Height.Int()
// 					if _h != h {
// 						err := fmt.Errorf("invalid block height: %d, expected: %d", _h, h+1)
// 						l.Info(err)
// 						return err
// 					}
// 					fmt.Println("this is the test log ", v)
// 					h++
// 					s++
// 					return nil
// 				},
// 				func(conn *websocket.Conn) {
// 					l.WithFields(log.Fields{"local": conn.LocalAddr().String()}).Debug("connected")
// 				},
// 				func(conn *websocket.Conn, err error) {
// 					l.WithFields(log.Fields{"error": err, "local": conn.LocalAddr().String()}).Warn("disconnected")
// 					_ = conn.Close()
// 				})
// 			if err.Error() == "context deadline exceeded" {
// 				return
// 			}

// 		}(i, url)
// 	}
// 	time.Sleep(time.Second * 10)

// }
