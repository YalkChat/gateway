package server

import (
	"log"
	"yalk/chat/models"
)

type MessageSender interface {
	Send(message *models.Message)
}

// func (server *Server) SendToAdmins(message *Message, payload []byte) {}

func (server *Server) SendToChat(message *models.Message, payload []byte) {
	// TODO: Move to server method
	chat := &models.Chat{ID: message.ChatID}
	user, err := chat.GetUsers(server.Db)
	if err != nil {
		log.Printf("Error getting chat users")
		return
	}
	for _, user := range user {
		if client, ok := server.Clients[user.ID]; ok {
			client.Msgs <- payload
		}
	}
}

func (server *Server) SendBack(id uint, payload []byte) {
	if client, ok := server.Clients[id]; ok {
		client.Msgs <- payload
	}
}

func (server *Server) SendAll(payload []byte) {
	for _, client := range server.Clients {
		client.Msgs <- payload
	}
}
