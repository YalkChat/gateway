package server

import (
	"yalk/newchat/client"
	"yalk/newchat/models/events"
)

type Server interface {
	RegisterClient(client.Client) error
	UnregisterClient(client.Client) error
	// TODO: Should I use a client for models.Message?
	SendToChat(*events.Message) error
	BroadcastMessage(*events.Message) error
	GetClientByID(string) (client.Client, error)
	HandleEvent(*events.BaseEventWithMetadata) error
}

// Additional type definitions for Client, Message, etc.
