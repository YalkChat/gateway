package server

import (
	"yalk/newchat/client"
	"yalk/newchat/models"
)

type Server interface {
	RegisterClient(client client.Client) error
	UnregisterClient(client client.Client) error
	// TODO: Should I use a client for models.Message?
	SendToChat(message *models.Message) error
	HandleEvent(event *models.BaseEventWithMetadata) error
}

// Additional type definitions for Client, Message, etc.
