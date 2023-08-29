package types

import "encoding/json"

type NewMessageEvent struct {
	clientID string
	data     json.RawMessage
	// eventType string
}

// TODO: This must go in the RawPayload
func (e NewMessageEvent) Type() string {
	return "NewMessage"
}

func (e NewMessageEvent) Data() json.RawMessage {
	return e.data
}

func (e NewMessageEvent) ClientID() string {
	return e.clientID
}
