package events

// MessageEvent for handling chat messages
type MessageEvent struct {
	ChatID    string
	ClientID  string
	Content   string
	EventType string
}

func (e *MessageEvent) Type() string {
	return e.EventType
}

// UserActivityEvent for handling user activities like online, offline, etc.
type UserActivityEvent struct {
	ClientID  string
	Status    string
	EventType string
}

func (e *UserActivityEvent) Type() string {
	return e.EventType
}

// Add more specific event types as needed
