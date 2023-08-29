package event

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Event interface {
	Type() string
	Data() json.RawMessage
	ClientID() string
	// Other methods as needed
}

type HandlerContext struct {
	DB         *gorm.DB
	SendToChat func(string, Event) error
}

// Handler defines the methods that any event handler must implement
// TODO: I must chose whether I want to keep the DB here, or use something else
type Handler interface {
	HandleEvent(*HandlerContext, Event) error
}
