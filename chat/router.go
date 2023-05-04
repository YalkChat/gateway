package chat

import (
	"encoding/json"
	"yalk-backend/logger"
)

func encodePayload(event *EventMessage) ([]byte, error) {
	eventMessageJson, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}
	return eventMessageJson, nil
}

// Echoing to client is default behavior for error checking.

// Broadcast to all
// TODO: Error checking
func (server *Server) SendMessage(event *EventMessage) {
	for userId, client := range server.Clients {
		if userId != event.Sender {
			payload, err := encodePayload(event)
			if err != nil {
				logger.Err("ROUTER", "Error encoding payload")
			}
			client.Msgs <- payload
		}
	}
}

// Send to one or multiple connected clients
func (server *Server) SendToClients(event *EventMessage) {
	for _, id := range event.Receivers {
		wsClient := server.Clients[id]
		if wsClient != nil {
			payload, err := encodePayload(event)

			if err != nil {
				logger.Err("ROUTER", "Error encoding payload")
			}
			wsClient.Msgs <- payload
		}
	}
}

func (server *Server) Router() {
	for {
		select {
		case event := <-server.Channels.Notify:
			logger.Info("ROUTER", "Router: Notify received")
			server.SendMessage(event)
		case event := <-server.Channels.Login:
			logger.Info("ROUTER", "Router: Login received")
			server.SendMessage(event)

		case event := <-server.Channels.Logout:
			logger.Info("ROUTER", "Router: Logout received")
			server.SendMessage(event)

		case event := <-server.Channels.Msg:
			logger.Info("ROUTER", "Router: Broadcast message received")
			server.SendMessage(event)

		case event := <-server.Channels.Dm:
			logger.Info("ROUTER", "Router: Dm received")
			server.SendToClients(event)

		}
	}
}
