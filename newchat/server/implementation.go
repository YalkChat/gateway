// chat/server/implementation.go
package server

import (
	"fmt"
	"log"
	"sync"

	"yalk/database"
	"yalk/newchat/client"
	"yalk/newchat/event"
	"yalk/newchat/event/handlers"
	"yalk/newchat/eventbus"
	"yalk/newchat/message"

	"gorm.io/gorm"
)

type serverImpl struct {
	clients       map[string]client.Client
	mu            sync.Mutex
	eventHandlers map[string]event.Handler
	eventBus      eventbus.EventBus
	db            *gorm.DB
}

func NewServer(db *gorm.DB, eb eventbus.EventBus) Server {
	s := &serverImpl{
		clients:       make(map[string]client.Client),
		eventHandlers: make(map[string]event.Handler),
		eventBus:      eb,
		db:            db,
	}

	s.eventBus.Subscribe("MessageCreate", handlers.HandleMessageCreate(db))
	s.eventBus.Subscribe("ChatMessage", handlers.HandleChatMessageEvent(db))
	s.eventBus.Subscribe("UserStatus", handlers.HandleUserStatusEvent(db))
	return s
}

func (s *serverImpl) RegisterClient(client client.Client) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if client already exists
	if _, exists := s.clients[client.ID()]; exists {
		return fmt.Errorf("client with ID %s already registered", client.ID())
	}

	// Add client to internal tracking
	s.clients[client.ID()] = client

	// Start the receiver and sender Go routines for this client
	go s.StartReceiver(client, someChannel)
	go s.StartSender(client, someChannel)

	// Notify other components or clients as needed

	return nil
}

// chat/server/implementation.go
func (s *serverImpl) UnregisterClient(client client.Client) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if client exists
	if _, exists := s.clients[client.ID()]; !exists {
		return fmt.Errorf("client with ID %s not found", client.ID())
	}

	// Remove client from internal tracking
	delete(s.clients, client.ID())

	// Notify other components or clients as needed

	return nil
}

// chat/server/implementation.go
func (s *serverImpl) SendToChat(message *message.Message, chatID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Find the clients associated with the chatID
	clients, err := s.getClientsByChatID(chatID)
	if err != nil {
		return err
	}

	// Send the message to all clients in the chat
	for _, client := range clients {
		if err := client.SendMessage(message); err != nil {
			// Handle or log the error if needed
		}
	}

	// Notify other components or clients as needed

	return nil
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
		if client, exists := s.clients[client.ID]; exists {
			clientsInChat = append(clientsInChat, client)

		}
	}

	return clientsInChat, nil
}

// TODO: Revisit for specialized event handling
func (s *serverImpl) HandleEvent(event event.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Look up the event handler for the given event type
	handler, exists := s.eventHandlers[event.Type()]
	if !exists {
		return fmt.Errorf("no handler registered for event type %s", event.Type())
	}

	// Pass the event to the appropriate handler
	return handler.HandleEvent(event)
}

// Register an event handler
func (s *serverImpl) RegisterEventHandler(eventType string, handler handlers.Event) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.eventHandlers[eventType] = handler
}

func (s *serverImpl) StartReceiver(client client.Client, eventChannel chan event.Event, quit chan struct{}) {
	for {
		select {
		case <-quit:
			// Handle cleanup if needed
			return
		default:
			receivedEvent, err := client.ReadEvent()
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

func (s *serverImpl) StartSender(c client.Client, outgoingEvents chan event.Event) {
	for event := range outgoingEvents {
		if err := c.SendEvent(event); err != nil {
			// handle or log the error
		}
	}
}
