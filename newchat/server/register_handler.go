package server

import "yalk/newchat/event"

// Register an event handler
func (s *serverImpl) RegisterEventHandler(eventType string, handler event.Handler) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers[eventType] = handler
}
