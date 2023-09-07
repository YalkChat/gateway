package event

import (
	"encoding/json"

	"gorm.io/gorm"
)

type eventImpl struct {
	eventType string
	data      json.RawMessage
	clientID  string
}

func NewEvent(eventType string, data json.RawMessage) Event {
	return &eventImpl{
		eventType: eventType,
		data:      data,
	}
}

func (e *eventImpl) Type() string {
	return e.eventType
}

func (e *eventImpl) Data() json.RawMessage {
	return e.data
}

func (e *eventImpl) ClientID() string {
	return e.clientID
}

// TODO: Check if ther should be actual code in here?
func (e *eventImpl) HandleEvent(db *gorm.DB, event Event) error {
	return nil
}
