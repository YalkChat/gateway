package chat

import (
	"log"
	"time"
)

func Sender(ctx *MessageChannelsContext) {
	defer func() {
		ctx.WaitGroup.Done()
	}()

Run:
	for {
		select {
		case <-ctx.NotifyChannel:
			log.Println("Sender - got shutdown signal")
			break Run
		case payload := <-ctx.ClientData.Msgs:
			err := ClientWriteWithTimeout(ctx.Request.Context(), time.Second*5, ctx.Connection, payload)
			if err != nil {
				break Run
			}
		}
	}
}
