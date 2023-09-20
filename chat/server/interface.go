package server

import (
	"yalk/chat/client"
	"yalk/chat/models/events"
)

type Server interface {
	RegisterClient(client.Client) error
	UnregisterClient(client.Client) error
	SendChat(*events.BaseEvent, uint) error
	SendAll(*events.BaseEvent) error
	GetClientByID(uint) (client.Client, error)
	HandleEvent(*events.BaseEventWithMetadata) error
	GetUserByID(uint) (*events.User, error)
}

// Additional type definitions for Client, Message, etc.
// TODO: IMPORTANT. Server is the only things to be used, it's db
// .. methods for now are private and server provides an abstraction
