package server

import (
	"log"
	"yalk/newchat/client"
	"yalk/newchat/event"
)

func (s *serverImpl) StartReceiver(client client.Client, eventChannel chan event.Event, quit chan struct{}) {
	for {
		select {
		case <-quit:
			// Handle cleanup if needed
			return
		default:
			messageType, receivedEvent, err := client.ReadMessage()
			if err != nil {
				// Handle the error, possibly by logging it and breaking the loop to stop the receiver
				log.Printf("Error reading event from client %s: %v", client.ID(), err)
				break
			}
			// Forward the event to the server for handling
			if err := s.HandleEvent(receivedEvent); err != nil {
				// Handle the error, possibly by logging it
				log.Printf("Error handling event from client %s: %v", client.ID(), err)
			}
		}
	}
}
