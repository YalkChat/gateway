package server

import (
	"yalk/newchat/client"
	"yalk/newchat/event"
	"yalk/newchat/models"
)

type Server interface {
	RegisterClient(client client.Client) error
	UnregisterClient(client client.Client) error
	SendToChat(message *models.Message, chatID string) error
	HandleEvent(event event.Event) error
}

// Additional type definitions for Client, Message, etc.
