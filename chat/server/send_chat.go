package server

import (
	"fmt"
	"yalk/chat/client"
	"yalk/chat/models/events"

	"nhooyr.io/websocket"
)

func (s *serverImpl) SendChat(baseEvent *events.BaseEvent, chatId uint) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Find the clients associated with the chatID
	clients, err := s.getClientsByChatID(chatId)
	if err != nil {
		return err
	}

	messageBytes, err := s.serializer.Serialize(baseEvent)
	if err != nil {
		return fmt.Errorf("error serializing baseEvent: %v", err)
	}

	// Send the message to all clients in the chat
	for _, client := range clients {
		if err := client.SendMessageWithTimeout(websocket.MessageText, messageBytes); err != nil {
			return err
		}
	}
	return nil
}

func (s *serverImpl) getClientsByChatID(chatID uint) ([]client.Client, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var clientsInChat []client.Client

	users, err := s.db.GetUsersByChatId(chatID)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if client, exists := s.clients[user.ID]; exists {
			clientsInChat = append(clientsInChat, client)

		}
	}

	return clientsInChat, nil
}
