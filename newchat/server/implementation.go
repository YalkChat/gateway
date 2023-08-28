// chat/server/implementation.go
package server

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"yalk/database"
	"yalk/newchat/client"
	"yalk/newchat/event"
	"yalk/newchat/event/handlers"
	"yalk/newchat/event/types"
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

// Helper function to get clients IDs from the database by chatID

// func (s *serverImpl) HandleIncomingMessage(message message.Message) error {
// 	// Process the message, e.g., save to database, validate, etc.
// 	// This could be a call to another component responsible for messages processing

// 	for {
// 		for _, c := range s.clients {
// 			message, err := c.ReadMessage()
// 			if err != nil {
// 				// handle error
// 				continue
// 			}

// 			// Dispatch the message to the appropriate event handler
// 			if handler, exists := s.eventHandler[message.Type()]; exists {
// 				handler.HandleEvent(message)
// 			}
// 		}
// 	}
// 	// Broadcast the message to other clients in the same chat
// 	if err := s.BroadcastMessage(message); err != nil {
// 		return err
// 	}
// 	return nil
// }

// Placeholder for the message processing logic
func processMessage(message message.Message) error {
	// TODO: Implement the logic for processing the incoming message
	// This could involve saving the message to a database, validation, etc.
	return nil
}

func (s *serverImpl) BroadcastMessage(message message.Message) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Fetch the list of clienst in the same chat room as the message sender
	chatID := message.ChatID()
	clientsInChat, err := s.getClientsByChatID(chatID)
	if err != nil {
		return err
	}

	// Send the message to each client
	for _, client := range clientsInChat {
		if err := client.SendMessage(&message); err != nil {
			// Log error but continue sending to other clients
			log.Printf("Error sending message to client %s: %v", client.ID(), err)
		}
	}
	return nil
}

func (s *serverImpl) HandleEvent(event event.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Look up the event handler for the given event type
	handler, exists := s.eventHandlers[event.Type()]
	if !exists {
		return fmt.Errorf("no handler registered for event type %s", event.Type())
	}

	// Based on the event type, unmarshal into the appropriate struct
	switch event.Type() {
	case "MessageCreate":
		var msgData types.MessageCreateEvent

		if err := json.Unmarshal(event.Data(), &msgData); err != nil {
			return err
		}
		// Now msgData is a fully populated MessageCreateEventData struct
		// Pass it to the handler
		return handler.HandleEvent(msgData)
	// Add case for other event types
	default:
		return fmt.Errorf("unknown event type %s", event.Type())

	}

}

func (s *serverImpl) BroadcastEvent(event event.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Determine the clients interested in this event
	// ...

	// Sedn the event to each interested client

	for _, client := range clients {
		if err := client.SendEvent(event, event.ClientID()); err != nil {
			// Handle error, e.g., log, remove client, etc.
		}
	}
	return nil
}

// Register an event handler
func (s *serverImpl) RegisterEventHandler(eventType string, handler handlers.Event) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.eventHandlers[eventType] = handler
}

func (s *serverImpl) ListenForClientEvents(c client.Client) {
	for {
		// Read a message from the WebSocket connection
		// This is a simplified example; we'll use a more complex message format

		var incomingMessage event.Event
		err := c.ReadMessage(&incomingMessage) // Assume ReadMessage is a method on your Client interface
		if err != nil {
			log.Printf("Error reading message from client %s: %v", c.ID(), err)
		}

		// Dispatch the message to the appropriate handler
		if handler, exists := s.eventHandlers[incomingMessage.Type()]; exists {
			err := handler.HandleEvent(incomingMessage)
			if err != nil {
				log.Printf("Error handling event: %v", err)
			}
		} else {
			log.Printf("No handler found for event type %s", incomingMessage.Type())
		}

	}
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
