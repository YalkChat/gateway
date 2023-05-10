package chat

import (
	"encoding/json"
	"fmt"

	"yalk/chat/events"
	"yalk/chat/messages"
	"yalk/logger"
)

// * Handle incoming user payload and process it eventually
// * forwarding in the correct routine channels for other users to receive.
func (server *Server) HandlePayload(jsonEventMessage []byte) (err error) {

	eventMessage, err := events.DecodeEventMessage(jsonEventMessage)
	if err != nil {
		logger.Err("BROAD", "Listener - Error decoding EventMessage")
		return err
	}

	// ! TEMPs
	// eventMessage.ID = fmt.Sprintf("%v", rand.Uint32())

	// * Broadcasting event to correct channel
	switch eventMessage.Type {
	case "channel_message":
		// ? It's own function to share with DM?
		var message *messages.Message
		if err := json.Unmarshal([]byte(message.Content), &message); err != nil {
			logger.Err("HNDL", fmt.Sprintf("Failed to unmarshal channel message content: %v", err))
			return err
		}

		if err := message.SaveToDb(message.ConversationID, server.Db); err != nil {
			return err
		}
		server.Channels.Msg <- eventMessage

	case "direct_message":
		server.Channels.Dm <- eventMessage

	case "user_login":
		server.Channels.Notify <- eventMessage

	case "user_logout":
		server.Channels.Notify <- eventMessage

	case "user_update":
		server.Channels.Notify <- eventMessage

	case "chat_create":
		server.Channels.Notify <- eventMessage

	case "chat_delete":
		server.Channels.Notify <- eventMessage

	case "chat_join":
		server.Channels.Notify <- eventMessage

	default:
		logger.Warn("HNDL", "Payload Handler received an invalid event type")
		return fmt.Errorf("invalid_request")
	}
	return nil
}
