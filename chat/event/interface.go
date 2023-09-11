package event

import (
	"encoding/json"
	"yalk/chat/database"
	"yalk/chat/models/events"
	"yalk/chat/serialization"
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
	DB         database.DatabaseOperations
	Serializer serialization.SerializationStrategy
	SendToChat func(*events.BaseEvent, uint) error
	SendAll    func(*events.BaseEvent) error
}
