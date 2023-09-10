package event

import (
	"encoding/json"
	"yalk/chat/database"
	"yalk/chat/models/events"
)

type Event interface {
	Type() string
	Data() json.RawMessage
	ClientID() string
	// Other methods as needed
}

// Handler defines the methods that any event handler must implement
// TODO: I must chose whether I want to keep the DB here, or use something else
type Handler interface {
	HandleEvent(*HandlerContext, *events.BaseEvent) error
}

type HandlerContext struct {
	DB                database.DatabaseOperations
	SendMessageToChat func(*events.Message) error
	SendToAll         func()
}
