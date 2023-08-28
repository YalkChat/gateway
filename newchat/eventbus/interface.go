package eventbus

import "yalk/newchat/event"

// TODO: Check if we can be more specific rather than using "any", maybe with event.Event
type EventBus interface {
	Subscribe(eventType string, handler func(event.Event) error)
	Unsubscribe(eventType string, handler func(event.Event) error)
	Publish(eventType string, event event.Event) error
}
