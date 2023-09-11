package server

import (
	"yalk/chat/client"
	"yalk/chat/models/events"
)

type Server interface {
	RegisterClient(client.Client) error
	UnregisterClient(client.Client) error
	SendToChat(*events.BaseEvent, uint) error
	BroadcastMessage(*events.BaseEvent) error
	GetClientByID(uint) (client.Client, error)
	HandleEvent(*events.BaseEventWithMetadata) error
}

// Additional type definitions for Client, Message, etc.
