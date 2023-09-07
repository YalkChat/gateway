package events

import "encoding/json"

type BaseEvent struct {
	Opcode   string          `json:"opcode"`
	Data     json.RawMessage `json:"data"`
	ClientID string          `json:"clientID"`
	Type     string          `json:"type"`
}
