package server

import (
	"yalk/chat/models"
)

type ChatServer interface {
	// Register a new client connection
	RegisterClient(client models.Client)

	// Unregister a client connection
	UnregisterClient(client models.Client)

	// Send a message to a specific chat room
	SendToChat(message *models.Message, chatID string)

	// Broadcast a message to all connected clients
	BroadcastMessage(message *models.Message)

	// Handle a specific chat-related command or event
	HandleEvent(event *models.Event)

	// Other chat-related meethods as needed
}
