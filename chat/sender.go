package chat

import (
	"log"
	"time"
)

func (server *Server) Sender(client *Client, ctx *EventContext) {
	defer ctx.WaitGroup.Done()

Run:
	for {
		select {
		case <-ctx.NotifyChannel:
			log.Println("Sender - got shutdown signal")
			break Run
		case payload := <-client.Msgs:
			err := ClientWriteWithTimeout(ctx.Request.Context(), time.Second*5, ctx.Connection, payload)
			if err != nil {
				break Run
			}
		}
	}
}
