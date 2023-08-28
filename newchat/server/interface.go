package server

import (
	"yalk/newchat/client"
	"yalk/newchat/event"
	"yalk/newchat/message"
)

type Server interface {
	RegisterClient(client client.Client) error
	UnregisterClient(client client.Client) error
	SendToChat(message *message.Message, chatID string) error
	HandleEvent(event event.Event) error
}

// Additional type definitions for Client, Message, etc.
