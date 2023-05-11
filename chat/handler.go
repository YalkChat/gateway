package chat

import (
	"fmt"

	"yalk/logger"
)

// * Handle incoming user payload and process it eventually
// * forwarding in the correct routine channels for other users to receive.
func (server *Server) HandlePayload(jsonEventMessage []byte) (err error) {

	rawEvent, err := DecodeEventMessage(jsonEventMessage)
	if err != nil {
		logger.Err("HNDL", "Error decoding RawEvent")
		return err
	}

	switch rawEvent.Type {
	case "chat_message":
		message, err := newMessage(rawEvent)
		if err != nil {
			logger.Err("HNDL", " - Error creating Chat Message")
			return err
		}

		if err := message.SaveToDb(server.Db); err != nil {
			return err
		}
		server.Channels.Msg <- message

	case "direct_message":

		server.Channels.Dm <- rawEvent

	case "user_login":
		server.Channels.Notify <- rawEvent

	case "user_logout":
		server.Channels.Notify <- rawEvent

	case "user_update":
		server.Channels.Notify <- rawEvent

	case "chat_create":
		server.Channels.Notify <- rawEvent

	case "chat_delete":
		server.Channels.Notify <- rawEvent

	case "chat_join":
		server.Channels.Notify <- rawEvent

	default:
		logger.Warn("HNDL", "Payload Handler received an invalid event type")
		return fmt.Errorf("invalid_request")
	}
	return nil
}

func newMessage(rawEvent *RawEvent) (*Message, error) {
	message := &Message{}

	if err := message.Deserialize(rawEvent.Data); err != nil {
		logger.Err("HNDL", " - Error Deserializing Chat Message")
		return nil, err
	}
	return message, nil
}
