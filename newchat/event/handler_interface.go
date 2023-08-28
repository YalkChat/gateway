package event

// Handler defines the methods that any event handler must implement
type Handler interface {
	HandleEvent(any) error
}
