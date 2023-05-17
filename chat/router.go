package chat

import (
	"encoding/json"
	"yalk/logger"
)

// Echoing to client is default behavior for error checking.

// Broadcast to all
// TODO: Error checking
func (server *Server) SendToChat(message *Message, payload []byte) {
	for _, user := range message.Chat.Users {
		if client, ok := server.Clients[user.ID]; ok {
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
		case message := <-server.Channels.Msg:
			logger.Info("ROUTER", "Router: Broadcast message received")

			serializedData, err := message.Serialize()
			if err != nil {
				logger.Err("ROUTER", "Error serializing")
			}

			newRawEvent := RawEvent{Type: message.Type(), Data: serializedData}

			jsonEvent, err := json.Marshal(newRawEvent)
			if err != nil {
				logger.Err("ROUTER", "Error serializing RawEvent")
			}

			server.SendToChat(message, jsonEvent)

			// case event := <-server.Channels.Dm:
			// 	logger.Info("ROUTER", "Router: Dm received")

			// server.SendMessageToAll(event)

		}
	}
}
