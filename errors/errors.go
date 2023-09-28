package errors

import (
	"fmt"
	"net/http"
)

// Custom error types
var (
	ErrSessionValidation  = fmt.Errorf("session validation failed")
	ErrSessionDeletion    = fmt.Errorf("session deletion failed")
	ErrWebSocketUpgrade   = fmt.Errorf("websocket upgrade failed")
	ErrUserFetch          = fmt.Errorf("failed to fetch user")
	ErrNewClient          = fmt.Errorf("failed to create new client")
	ErrClientRegistration = fmt.Errorf("failed to register client")
	ErrInvalidMethodGet   = fmt.Errorf("wrong request method, expected get")
	ErrInvalidMethodPost  = fmt.Errorf("wrong request method, expected post")
	ErrCookieMissing      = fmt.Errorf("cookie missing")
)

// TODO: Placeholder, finish implementation
func HandleError(w http.ResponseWriter, r *http.Request, err error) {
	switch err {
	case ErrSessionValidation:
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	case ErrWebSocketUpgrade, ErrUserFetch, ErrNewClient, ErrClientRegistration:
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	default:
		http.Error(w, "Unknown Error", http.StatusInternalServerError)
	}
	fmt.Println("error: ", err)
}
