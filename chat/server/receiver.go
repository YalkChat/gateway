package server

import (
	"io"
	"log"
	"yalk/chat/chatmodels"

	"nhooyr.io/websocket"
)

func (server *Server) Receiver(clientID uint, ctx *chatmodels.EventContext) {
	defer func() {
		ctx.WaitGroup.Done()
		ctx.NotifyChannel <- true // TODO: Verify why it was heren
	}()

	// Run:
	for {
		messageType, jsonEventMessage, err := ctx.Connection.Read(ctx.Request.Context())

		log.Printf("Received payload lenght: %d", len(jsonEventMessage))

		if err != nil && err != io.EOF {

			statusCode := websocket.CloseStatus(err)

			if statusCode == websocket.StatusGoingAway {
				log.Println("Graceful sender shutdown")
				ctx.PingTicket.Stop()
				return
				// break Run

			} else {
				log.Println("Sender - Error in reading from websocket context, client closed? Check main.go")
				// break Run
				return
			}
		}

		if messageType.String() == "MessageText" && err == nil {
			log.Printf("Message received: %s", jsonEventMessage)

			rawEvent := &chatmodels.RawEvent{UserID: clientID}

			if err := rawEvent.Deserialize(jsonEventMessage); err != nil {
				log.Printf("Error unmarshaling RawEvent")
				return
			}
			if err := server.HandleIncomingEvent(clientID, rawEvent); err != nil {
				log.Printf("Sender - errors in broadcast: %v", err)
				return
			}
		}
	}
}
