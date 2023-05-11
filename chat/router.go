package chat

import (
	"encoding/json"
	"yalk/logger"

	"golang.org/x/exp/slices"
)

// Echoing to client is default behavior for error checking.

// Broadcast to all
// TODO: Error checking
func (server *Server) SendMessage(userIds []uint, payload []byte) {
	for userId, client := range server.Clients {
		if slices.Contains(userIds, userId) {
			client.Msgs <- payload
		}
	}
}

// Send to one or multiple connected clients
// func (server *Server[T]) SendMessageToAll(event *Event) {
// 	payload, err := EncodeEventMessage(event)

// 	if err != nil {
// 		logger.Err("ROUTER", "Error encoding payload")
// 	}

// 	// TODO: New logic
// 	for _, id := range event.Receivers {
// 		wsClient := server.Clients[id]
// 		if wsClient != nil {

// 			wsClient.Msgs <- payload
// 		}
// 	}
// }

func (server *Server) Router() {
	for {
		select {
		// case event := <-server.Channels.Notify:
		// 	logger.Info("ROUTER", "Router: Notify received")
		// 	server.SendMessage(event)
		// case event := <-server.Channels.Login:
		// 	logger.Info("ROUTER", "Router: Login received")
		// 	server.SendMessage(event)

		// case event := <-server.Channels.Logout:
		// 	logger.Info("ROUTER", "Router: Logout received")
		// 	server.SendMessage(event)

		case message := <-server.Channels.Msg:
			logger.Info("ROUTER", "Router: Broadcast message received")

			userIds, err := GetChatUserIds(message.ConversationID, server.Db)
			if err != nil {
				logger.Err("ROUTER", "Shurizzle si e' rotto il cazzo")
				return
			}

			serializedData, err := message.Serialize()
			if err != nil {
				logger.Err("ROUTER", "Error serializing")
			}

			newRawEvent := RawEvent{Type: message.Type(), Data: serializedData}

			jsonEvent, err := json.Marshal(newRawEvent)
			if err != nil {
				logger.Err("ROUTER", "Error serializing RawEvent")
			}

			server.SendMessage(userIds, jsonEvent)

			// case event := <-server.Channels.Dm:
			// 	logger.Info("ROUTER", "Router: Dm received")

			// server.SendMessageToAll(event)

		}
	}
}
