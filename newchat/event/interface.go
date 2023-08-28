package event

import "encoding/json"

type Event interface {
	Type() string
	Data() json.RawMessage
	ClientID() string
	// Other methods as needed
}

// Define a new interface for chat events
type ChatEvent interface {
	Type() string
	Data() json.RawMessage
	ChatID() string
}
