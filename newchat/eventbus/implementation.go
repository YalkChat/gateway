package eventbus

import (
	"sync"
	"yalk/newchat/event"
)

type eventBusImpl struct {
	subscribers map[string][]func(event.Event) error
	mu          sync.Mutex
}

func NewEventBus() EventBus {
	return &eventBusImpl{
		subscribers: make(map[string][]func(event.Event) error),
	}
}

func (eb *eventBusImpl) Subscribe(eventType string, handler func(event.Event) error) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.subscribers[eventType] = append(eb.subscribers[eventType], handler)
}

func (eb *eventBusImpl) Unsubscribe(eventType string, handler func(event.Event) error) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	delete(eb.subscribers, eventType)
}

// TODO: Check if we can be more specific
func (eb *eventBusImpl) Publish(eventType string, event event.Event) error {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	if handlers, found := eb.subscribers[eventType]; found {
		for _, handler := range handlers {
			if err := handler(event); err != nil {
				return err
			}
		}
	}

	return nil
}

// New method to clear all handlers for a specific event type
func (eb *eventBusImpl) ClearHandlers(eventType string) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	delete(eb.subscribers, eventType)
}
