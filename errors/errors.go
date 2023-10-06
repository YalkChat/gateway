package errors

import (
	"fmt"
	"net/http"
)

// Custom error types
var (
	ErrSessionCreation     = fmt.Errorf("session creation failed")
	ErrSessionValidation   = fmt.Errorf("session validation failed")
	ErrSessionDeletion     = fmt.Errorf("session deletion failed")
	ErrWebSocketUpgrade    = fmt.Errorf("websocket upgrade failed")
	ErrUserFetch           = fmt.Errorf("failed to fetch user")
	ErrNewClient           = fmt.Errorf("failed to create new client")
	ErrClientRegistration  = fmt.Errorf("failed to register client")
	ErrInvalidMethodGet    = fmt.Errorf("wrong request method, expected get")
	ErrInvalidMethodPost   = fmt.Errorf("wrong request method, expected post")
	ErrCookieMissing       = fmt.Errorf("cookie missing")
	ErrAuthInvalid         = fmt.Errorf("invalid authentication")
	ErrInvalidJson         = fmt.Errorf("invalid JSON")
	ErrInternalServerError = fmt.Errorf("internal server error")
	ErrValidSessionExists  = fmt.Errorf("valid session exists")  //TODO: this is not really an error
	ErrUserCreation        = fmt.Errorf("failed to create user") //TODO: this is not really an error
)

// TODO: Placeholder, finish implementation
func HandleError(w http.ResponseWriter, r *http.Request, err error) {
	switch err {
	case ErrSessionValidation:
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	case ErrWebSocketUpgrade, ErrUserFetch, ErrNewClient, ErrClientRegistration:
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	case ErrAuthInvalid:
		http.Error(w, "Invalid authorizaton", http.StatusUnauthorized)
	case ErrValidSessionExists:
		http.Redirect(w, r, "/chat", http.StatusFound)

	default:
		http.Error(w, "Unknown Error", http.StatusInternalServerError)
	}
	fmt.Println("error: ", err)
}
