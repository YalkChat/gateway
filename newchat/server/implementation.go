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

func (s *serverImpl) BroadcastMessage(message *models.Message) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// TODO: with the observations on top of this file evaluated if
	// .. the serialization back to Json could happen in an helper function
	// .. called by the handlers before sending the message, or any better way

	// Serialize the message back to JSON bytes
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	// Create a new BaseEvent
	// TODO: Check where to get the client ID from, if the message is good or anything better
	// TODO: Check also where to get eventType from
	baseEvent, err := newBaseEvent("1", messageBytes, message.ClientID, "NewMessage")
	if err != nil {
		return fmt.Errorf("error broadcasting message: %v", err)
	}

	baseEventJSON, err := json.Marshal(baseEvent)
	if err != nil {
		return fmt.Errorf("failed to serialize base event: %v", err)
	}

	for _, client := range s.clients {
		if err := client.SendMessageWithTimeout(websocket.MessageText, baseEventJSON); err != nil {
			// Handle the error based on your application's needs
			fmt.Printf("failed to send message to client %s: %v\n", client.ID(), err)
		}
	}
	return nil
}

// TODO: Implement error checking if args are empty
func newBaseEvent(opcode string, data json.RawMessage, clientID string, eventType string) (*models.BaseEvent, error) {
	baseEvent := &models.BaseEvent{
		Opcode:   opcode,
		Data:     data,
		ClientID: clientID,
		Type:     eventType,
	}
	// There must be a better way to do this error check
	if opcode == "" {
		return nil, fmt.Errorf("opcode empty")
	}
	return baseEvent, nil
}
