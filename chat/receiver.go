package chat

import (
	"fmt"
	"io"
	"log"

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
		fmt.Printf("Payload len: %v\n", len(payload))
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
			fmt.Printf("Message received: %s", payload)
			// err = server.handlePayload(payload, session.UserID)
			// if err != nil {
			// log.Printf("Sender - errors in broadcast: %v", err)
			// wg.Done()
			// return
			// }
		}
	}
}
