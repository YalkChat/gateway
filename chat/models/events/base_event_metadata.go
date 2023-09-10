package events

type BaseEventWithMetadata struct {
	Event  *BaseEvent
	UserID uint
}
