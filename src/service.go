package src

import (
	"context"
	"time"

	"github.com/basic-relayer/src/constants"
)

//start a schedule to run after some time

func Start() {
	ctx, _ := context.WithCancel(context.Background())

	for {
		select {
		case <-ctx.Done():
			break
		default:
			fetchPackets()
			time.Sleep(time.Second * time.Duration(constants.TIME_CONSTANT))
		}
	}
}
