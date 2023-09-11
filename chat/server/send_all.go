package server

import (
	"fmt"
	"yalk/chat/models/events"

	"nhooyr.io/websocket"
)

func (s *serverImpl) SendAll(baseEvent *events.BaseEvent) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	messageBytes, err := s.serializer.Serialize(baseEvent)
	if err != nil {
		return fmt.Errorf("error serializing baseEvent: %v", err)
	}

	for _, client := range s.clients {
		if err := client.SendMessageWithTimeout(websocket.MessageText, messageBytes); err != nil {
			// Handle the error based on your application's needs
			fmt.Printf("failed to send message to client %d: %v\n", client.ID(), err)
		}
	}
	return nil
}
