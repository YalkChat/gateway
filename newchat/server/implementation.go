// TODO: Consider the methods to send messages and notificatins: How should they be and what is the default structure to keep?
// TODO: We can have a structure like this with specific types on args and these names (SendToChat, SendNotification)..
// ..or more generic like SendToClient, SendToMany, Broadcast, and so on or even just use SendClient accepting an array of clients along..
// ..with the others above.
// TODO: Maybe not exported methods (which I guess are helper functions?) might go in one single utils.go file?

package server

import (
	"encoding/json"
	"fmt"
	"sync"
	"yalk/database"
	"yalk/newchat/client"
	"yalk/newchat/event"
	"yalk/newchat/event/handlers"
	"yalk/newchat/models"

	"nhooyr.io/websocket"
)

type serverImpl struct {
	clients  map[string]client.Client
	mu       sync.Mutex
	handlers map[string]event.Handler
	db       database.DatabaseOperations
}

func NewServer(db database.DatabaseOperations) Server {
	s := &serverImpl{
		clients:  make(map[string]client.Client),
		handlers: make(map[string]event.Handler),
		db:       db,
	}

	s.RegisterEventHandler("ChatMessage", handlers.NewMessageHandler{})
	return s
}

func (s *serverImpl) getClientsByChatID(chatID string) ([]client.Client, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var clientsInChat []client.Client

	clients, err := s.db.GetUsers(chatID)
	if err != nil {
		return nil, err
	}

	for _, client := range clients {
		if client, exists := s.clients[client]; exists {
			clientsInChat = append(clientsInChat, client)

		}
	}

	return clientsInChat, nil
}

// TODO: Revisit for specialized event handling, and make type with UserID metadata
func (s *serverImpl) HandleEvent(eventWithMetadata *models.BaseEventWithMetadata) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	baseEvent := eventWithMetadata.Event
	eventType := baseEvent.Type

	// Look up the event handler for the given event type
	handler, err := s.getHandler(eventType)
	if err != nil {
		return err
	}
	ctx := &event.HandlerContext{DB: s.db, SendToChat: s.SendToChat}

	// Pass the event to the appropriate handler
	return handler.HandleEvent(ctx, baseEvent)
}

func (s *serverImpl) getHandler(eventType string) (event.Handler, error) {
	handler, exists := s.handlers[eventType]
	if !exists {
		return nil, fmt.Errorf("no handler registered for event type %s", eventType)
	}
	return handler, nil
}

func (s *serverImpl) SendToChat(message *models.Message) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Serialize the message back to JSON bytes
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	// Find the clients associated with the chatID
	clients, err := s.getClientsByChatID(message.ChatID)
	if err != nil {
		return err
	}

	// Send the message to all clients in the chat
	for _, client := range clients {
		if err := client.SendMessageWithTimeout(websocket.MessageText, messageBytes); err != nil {
			return err
		}
	}
	return nil
}

// 	// Notify other components or clients as needed

// 	return nil
// }

// func (s *serverImpl) BroadcastEvent(event event.Event) error {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()

// 	// Sedn the event to each interested client

// 	for _, client := range clients {
// 		if err := client.SendEvent(event, event.ClientID()); err != nil {
// 			// Handle error, e.g., log, remove client, etc.
// 		}
// 	}
// 	return nil
// }
