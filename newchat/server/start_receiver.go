package server

import (
	"encoding/json"
	"log"
	"yalk/newchat/client"
	"yalk/newchat/event"
	"yalk/newchat/models"
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
			// TODO: Missing EOF check but unsure if needed
			if messageType.String() == "MessageText" {
				baseEvent := &models.BaseEvent{} // Assuming BaseEvent is your new RawEvent
				err := json.Unmarshal(receivedEvent, baseEvent)
				if err != nil {
					log.Println("Error unmarshalling:", err)
					continue
				}
				// Forward the event to the server for handling
				if err := s.HandleEvent(baseEvent); err != nil {
					// Handle the error, possibly by logging it
					log.Printf("Error handling event from client %s: %v", client.ID(), err)
				}
			}
		}
	}
}
