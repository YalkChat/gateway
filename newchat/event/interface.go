package event

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Event interface {
	Type() string
	HandleEvent(*gorm.DB, Event) error
	Data() json.RawMessage
	ClientID() string
	// Other methods as needed
}

// TODO: One must go, or we must find a way to keep both separated

// Handler defines the methods that any event handler must implement
// TODO: I must chose whether I want to keep the DB here, or use something else
// type Handler interface {
// 	HandleEvent(*gorm.DB, Event) error
// }
