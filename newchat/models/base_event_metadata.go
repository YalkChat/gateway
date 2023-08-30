package models

type BaseEventWithMetadata struct {
	Event  *BaseEvent
	UserID string
}
