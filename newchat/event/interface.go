package event

import "encoding/json"

type Event interface {
	Type() string
	Data() json.RawMessage
	ClientID() string
	// Other methods as needed
}
