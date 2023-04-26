package chat

import (
	"fmt"
	"io"
	"log"
	"yalk-backend/logger"

	"nhooyr.io/websocket"
)

func Receiver(ctx *MessageChannelsContext) {
	defer func() {
		ctx.WaitGroup.Done()
		ctx.NotifyChannel <- true
	}()
Run:
	for {
		t, payload, err := ctx.Connection.Read(ctx.Request.Context())
		logger.Info("RCV", fmt.Sprintf("Receiced payload lenght: %d", len(payload)))

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
		if t.String() == "MessageText" && err == nil {
			logger.Info("RCV", fmt.Sprintf("Message received: %s", payload))
			err = handlePayload(payload, "test", ctx)
			if err != nil {
				log.Printf("Sender - errors in broadcast: %v", err)
				return
			}
		}
	}
}
