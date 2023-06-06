package chat

import (
	"fmt"
	"log"
	"time"

	"yalk/chat/clients"
	"yalk/logger"
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
			logger.Info("SEND", fmt.Sprintf("Sending payload to %d", c.ID))
			err := clients.ClientWriteWithTimeout(c.ID, ctx.Request.Context(), time.Second*5, ctx.Connection, payload)
			if err != nil {
				// break Run
				return
			}
		}
	}
}
