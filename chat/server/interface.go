package server

import (
	"yalk/chat/client"
	"yalk/chat/models/events"
)

type Server interface {
	RegisterClient(client.Client) error
	UnregisterClient(client.Client) error
	// TODO: Should I use a client for models.Message?
	SendToChat(*events.Message) error
	BroadcastMessage(*events.Message) error
	GetClientByID(uint) (client.Client, error)
	HandleEvent(*events.BaseEventWithMetadata) error
}

// Additional type definitions for Client, Message, etc.
