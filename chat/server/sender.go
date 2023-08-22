package server

import (
	"log"
	"time"

	"yalk/chat/chatmodels"
)

func (server *Server) Sender(c *chatmodels.Client, ctx *chatmodels.EventContext) {
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
			log.Printf("Sending payload to %d", c.ID)
			// TODO: Method on server Types
			err := chatmodels.ClientWriteWithTimeout(c.ID, ctx.Request.Context(), time.Second*5, ctx.Connection, payload)
			if err != nil {
				// break Run
				return
			}
		}
	}
}
