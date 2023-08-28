package client

import (
	"yalk/newchat/event"
	"yalk/newchat/message"
)

type Client interface {
	// ID returns the unique identifier for the client
	ID() string

	// SendMessage sends a message to the client
	SendMessage(message *message.Message) error

	// SendEvent sends an event to the client
	SendEvent(event event.Event) error

	ReadEvent() (event.Event, error)

	// Other methods as needed, such as receiving messages, handling events, etc.
}
