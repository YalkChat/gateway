package handlers

import (
	"fmt"
	"net/http"
	"time"
	"yalk/app"
	"yalk/chat/client"

	"github.com/AleRosmo/cattp"
	"nhooyr.io/websocket"
)

// Custom error types
var (
	ErrSessionValidation  = fmt.Errorf("session validation failed")
	ErrWebSocketUpgrade   = fmt.Errorf("websocket upgrade failed")
	ErrUserFetch          = fmt.Errorf("failed to fetch user")
	ErrNewClient          = fmt.Errorf("failed to create new client")
	ErrClientRegistration = fmt.Errorf("failed to register client")
)

// TODO: Placeholder, finish implementation
func handleError(w http.ResponseWriter, r *http.Request, err error) {
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

// TODO: Handle Error must be finished
var ConnectionHandler = cattp.HandlerFunc[app.HandlerContext](func(w http.ResponseWriter, r *http.Request, ctx app.HandlerContext) {
	server := ctx.ChatServer()
	sessionsManager := ctx.SessionsManager()

	session, err := sessionsManager.Validate(r)
	if err != nil {
		handleError(w, r, ErrSessionValidation)
		return
	}

	// Upgrades to WebSocket
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		handleError(w, r, ErrWebSocketUpgrade)
		return
	}

	user, err := server.GetUserByID(session.UserID)
	if err != nil {
		handleError(w, r, ErrUserFetch)
		return
	}

	// TODO: More correct to use session.UserID or user.ID()?
	userID := user.ID

	client := client.NewClient(userID, conn, time.Second*5) // TODO: time placeholder
	if err != nil {
		handleError(w, r, ErrNewClient)
		return
	}

	err = server.RegisterClient(client)
	if err != nil {
		handleError(w, r, ErrClientRegistration)
		return
	}

	fmt.Printf("registered new client: %d", client.ID())

})
