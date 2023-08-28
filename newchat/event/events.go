package event

import "encoding/json"

// UserActivityEvent for handling user activities like online, offline, etc.
type UserActivityEvent struct {
	ClientID  string
	Status    string
	EventType string
}

func (e *UserActivityEvent) Type() string {
	return e.EventType
}

// MessageEvent for handling chat messages
type MessageEvent struct {
	EventType string
	// other fields
}

func (e *MessageEvent) Type() string          { return e.EventType }
func (e *MessageEvent) Data() json.RawMessage { /* ... */ }
func (e *MessageEvent) ChatID() string        { /* ... */ }

// Similarly, you can define other specific event types like MessageDeletedEvent, UserOnlineEvent, etc.

// Add more specific event types as needed
