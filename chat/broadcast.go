package chat

import (
	"encoding/json"
	"fmt"
	"log"
	"yalk-backend/logger"
)

// * Handle incoming user payload and process it eventually forwarding in the correct routine channels for other users to receive.
func handlePayload(msg []byte, origin string, ctx *EventContext) (err error) {
	var _req any
	var _payload map[string]any
	var data Payload
	var event string

	err = json.Unmarshal(msg, &_req)
	if err != nil {
		logger.Err("BROAD", "Listener - Error deserializing JSON")
		return err
	}

	switch p := _req.(type) {
	case map[string]any:
		_payload = p
	default:
		logger.Err("BROAD", "Listener - can't decode payload")
	}

	switch v := _payload["event"].(type) {
	case string:
		event = v
	default:
		log.Println("BROAD", "Listener - can't decode event")
	}

	data = Payload{
		Success: true,
		Origin:  origin,
		Event:   event,
	}

	// * Routing event to server
	switch event {
	case "chat_message":
		ctx.Server.Channels.Notify <- Payload{Success: true, Event: "chat_message"}

	case "user_logout":
		ctx.Server.Channels.Notify <- Payload{Success: true, Event: "user_logout"}

	case "user_update":
		ctx.Server.Channels.Notify <- Payload{Success: true, Event: "user_update"}

	case "chat_create":
		ctx.Server.Channels.Notify <- Payload{Success: true, Event: "chat_create"}

	case "chat_delete":
		ctx.Server.Channels.Notify <- Payload{Success: true, Event: "chat_delete"}

	case "chat_join":
		ctx.Server.Channels.Notify <- Payload{Success: true, Event: "chat_join"}

	case "user_login":
		ctx.Server.Channels.Notify <- Payload{Success: true, Event: "user_conn"}

	default:
		log.Println("Broadcast received an invalid request")
		data.Success = false
		return fmt.Errorf("invalid_request")
	}
	return nil
}
