package server

import "yalk/newchat/event"

// Register an event handler
func (s *serverImpl) RegisterEventHandler(eventType string, handler event.Event) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.eventHandlers[eventType] = handler
}
