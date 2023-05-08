package chat

import (
	"fmt"
	"yalk-backend/logger"

	"math/rand"
)

// * Handle incoming user payload and process it eventually
// * forwarding in the correct routine channels for other users to receive.
func (server *Server) HandlePayload(jsonEventMessage []byte) (err error) {

	message, err := decodeEventMessage(jsonEventMessage)
	if err != nil {
		logger.Err("BROAD", "Listener - Error decoding EventMessage")
		return err
	}

	// ! TEMP
	message.ID = fmt.Sprintf("%v", rand.Int())

	// * Broadcasting event to correct channel
	switch message.Type {
	case "channel_message":
		server.Channels.Msg <- message

	case "direct_message":
		server.Channels.Dm <- message

	case "user_login":
		server.Channels.Notify <- message

	case "user_logout":
		server.Channels.Notify <- message

	case "user_update":
		server.Channels.Notify <- message

	case "chat_create":
		server.Channels.Notify <- message

	case "chat_delete":
		server.Channels.Notify <- message

	case "chat_join":
		server.Channels.Notify <- message

	default:
		logger.Warn("HNDL", "Payload Handler received an invalid event type")
		message.Success = false
		return fmt.Errorf("invalid_request")
	}
	return nil
}
