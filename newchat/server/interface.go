package server

import (
	"yalk/newchat/client"
	"yalk/newchat/models"
)

type Server interface {
	RegisterClient(client.Client) error
	UnregisterClient(client.Client) error
	// TODO: Should I use a client for models.Message?
	SendToChat(*models.Message) error
	BroadcastMessage(*models.Message) error
	HandleEvent(*models.BaseEventWithMetadata) error
}

// Additional type definitions for Client, Message, etc.
