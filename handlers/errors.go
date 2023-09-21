package handlers

import "fmt"

// Custom error types
var (
	ErrSessionValidation  = fmt.Errorf("session validation failed")
	ErrWebSocketUpgrade   = fmt.Errorf("websocket upgrade failed")
	ErrUserFetch          = fmt.Errorf("failed to fetch user")
	ErrNewClient          = fmt.Errorf("failed to create new client")
	ErrClientRegistration = fmt.Errorf("failed to register client")
)
