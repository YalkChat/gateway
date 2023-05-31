package chat

import (
	"log"
	"time"

	"yalk/chat/clients"
)

func (server *Server) Sender(c *clients.Client, ctx *EventContext) {
	defer func() {
		ctx.WaitGroup.Done()
		// <-ctx.NotifyChannel
	}()

	// Run:
	for {
		select {
		case <-ctx.NotifyChannel:
			log.Println("Sender - got shutdown signal")
			// break Run
			return
		case payload := <-c.Msgs:
			err := clients.ClientWriteWithTimeout(ctx.Request.Context(), time.Second*5, ctx.Connection, payload)
			if err != nil {
				// break Run
				return
			}
		}
	}
}
