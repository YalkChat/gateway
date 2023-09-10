package server

import (
	"fmt"
	"yalk/chat/client"
)

func (s *serverImpl) RegisterClient(client client.Client) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if client already exists
	if _, exists := s.clients[client.ID()]; exists {
		return fmt.Errorf("client with ID %s already registered", client.ID())
	}

	// Add client to internal tracking
	s.clients[client.ID()] = client

	quit := make(chan struct{})

	// Start the receiver and sender Go routines for this client
	go s.StartReceiver(client, quit)
	// go s.StartSender(client, quit)

	// Notify other components or clients as needed

	return nil
}