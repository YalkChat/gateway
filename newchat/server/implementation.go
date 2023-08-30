// chat/server/implementation.go
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

	"gorm.io/gorm"
	"nhooyr.io/websocket"
)

type serverImpl struct {
	clients  map[string]client.Client
	mu       sync.Mutex
	handlers map[string]event.Handler
	db       *gorm.DB
}

func NewServer(db *gorm.DB) Server {
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

	clients, err := database.GetClients(s.db, chatID)
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
	handler, exists := s.handlers[eventType]
	if !exists {
		return fmt.Errorf("no handler registered for event type %s", eventType)
	}

	ctx := &event.HandlerContext{DB: s.db, SendToChat: s.SendToChat}

	// Pass the event to the appropriate handler
	return handler.HandleEvent(ctx, baseEvent)
}

func (s *serverImpl) SendToChat(message models.Message) error {
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

// // Placeholder for the message processing logic
// func processMessage(message message.Message) error {
// 	// TODO: Implement the logic for processing the incoming message
// 	// This could involve saving the message to a database, validation, etc.
// 	return nil
// }

// func (s *serverImpl) ListenForClientEvents(c client.Client) {
// 	for {
// 		// Read a message from the WebSocket connection
// 		// This is a simplified example; we'll use a more complex message format

// 		var incomingMessage event.Event
// 		err := c.ReadMessage(&incomingMessage) // Assume ReadMessage is a method on your Client interface
// 		if err != nil {
// 			log.Printf("Error reading message from client %s: %v", c.ID(), err)
// 		}

// 		// Dispatch the message to the appropriate handler
// 		if handler, exists := s.eventHandlers[incomingMessage.Type()]; exists {
// 			err := handler.HandleEvent(incomingMessage)
// 			if err != nil {
// 				log.Printf("Error handling event: %v", err)
// 			}
// 		} else {
// 			log.Printf("No handler found for event type %s", incomingMessage.Type())
// 		}

// 	}
// }

// func (s *serverImpl) BroadcastMessage(message models.Message) error {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()

// 	// Fetch the list of clienst in the same chat room as the message sender
// 	chatID := message.ChatID
// 	clientsInChat, err := s.getClientsByChatID(chatID)
// 	if err != nil {
// 		return err
// 	}

// 	// Send the message to each client
// 	for _, client := range clientsInChat {
// 		if err := client.SendMessage(websocket.MessageText, message); err != nil {
// 			// Log error but continue sending to other clients
// 			log.Printf("Error sending message to client %s: %v", client.ID(), err)
// 		}
// 	}
// 	return nil
// }

// func (s *serverImpl) BroadcastEvent(event event.Event) error {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()

// 	// Determine the clients interested in this event
// 	// ...

// 	// Sedn the event to each interested client

// 	for _, client := range clients {
// 		if err := client.SendEvent(event, event.ClientID()); err != nil {
// 			// Handle error, e.g., log, remove client, etc.
// 		}
// 	}
// 	return nil
// }
