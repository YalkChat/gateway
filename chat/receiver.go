package chat

import (
	"fmt"
	"io"
	"log"
	"yalk/logger"

	"nhooyr.io/websocket"
)

func (server *Server) Receiver(clientID uint, ctx *EventContext) {
	// TODO: var closingReason string
	defer func() {
		// Signalign that client is closing
		ctx.WaitGroup.Done()
		// ctx.NotifyChannel <- true // TODO: Verify why it was here
	}()

Run:
	for {
		messageType, payload, err := ctx.Connection.Read(ctx.Request.Context())

		logger.Info("RCV", fmt.Sprintf("Received payload lenght: %d", len(payload)))

		if err != nil && err != io.EOF {

			statusCode := websocket.CloseStatus(err)

			if statusCode == websocket.StatusGoingAway {
				log.Println("Graceful sender shutdown")
				ctx.PingTicket.Stop()
				break Run

			} else {
				log.Println("Sender - Error in reading from websocket context, client closed? Check main.go")
				break Run
			}
		}

		if messageType.String() == "MessageText" && err == nil {
			logger.Info("RCV", fmt.Sprintf("Message received: %s", payload))
			if err := server.HandleIncomingEvent(clientID, payload); err != nil {
				log.Printf("Sender - errors in broadcast: %v", err)
				return
			}
		}
	}
}
