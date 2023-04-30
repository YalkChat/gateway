package chat

import (
	"encoding/json"
	"fmt"
	"log"
	"yalk-backend/logger"
)

func decodeEventMessage(jsonEvent []byte) (*EventMessage, error) {
	var event *EventMessage

	if err := json.Unmarshal(jsonEvent, &event); err != nil {
		logger.Err("BROAD", "Listener - Error deserializing JSON")
		return nil, err
	}
	return event, nil
}

// * Handle incoming user payload and process it eventually
// * forwarding in the correct routine channels for other users to receive.
func (server *Server) handlePayload(jsonEventMessage []byte) (err error) {

	message, err := decodeEventMessage(jsonEventMessage)
	if err != nil {
		logger.Err("BROAD", "Listener - Error decoding EventMessage")
		return err
	}

	// * Routing event to server
	switch message.Type {
	case "channel_message":
		server.Channels.Notify <- message

	// case "direct_message":
	// 	server.Channels.Notify <- EventMessage{}

	// case "user_logout":
	// 	server.Channels.Notify <- Payload{Success: true, Event: "user_logout"}

	// case "user_update":
	// 	server.Channels.Notify <- Payload{Success: true, Event: "user_update"}

	// case "chat_create":
	// 	server.Channels.Notify <- Payload{Success: true, Event: "chat_create"}

	// case "chat_delete":
	// 	server.Channels.Notify <- Payload{Success: true, Event: "chat_delete"}

	// case "chat_join":
	// 	server.Channels.Notify <- Payload{Success: true, Event: "chat_join"}

	// case "user_login":
	// 	server.Channels.Notify <- Payload{Success: true, Event: "user_conn"}

	default:
		log.Println("Broadcast received an invalid request")
		message.Success = false
		return fmt.Errorf("invalid_request")
	}
	return nil
}
