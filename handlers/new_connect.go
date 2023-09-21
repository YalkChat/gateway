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

// TODO: Placeholder, finish implementation
func handleError(w http.ResponseWriter, err error) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)

}

// TODO: Handle Error must be finished
var ConnectionHandler = cattp.HandlerFunc[app.HandlerContext](func(w http.ResponseWriter, r *http.Request, ctx app.HandlerContext) {
	server := ctx.ChatServer()
	sessionsManager := ctx.SessionsManager()

	session, err := sessionsManager.Validate(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Upgrades to WebSocket
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	user, err := server.GetUserByID(session.UserID)
	if err != nil {
		handleError(w, err)
		return
	}

	// TODO: More correct to use session.UserID or user.ID()?
	userID := user.ID

	client := client.NewClient(userID, conn, time.Second*5) // TODO: time placeholder
	if err != nil {
		handleError(w, err)
	}

	err = server.RegisterClient(client)

	if err != nil {
		handleError(w, err)
		return
	}

	fmt.Printf("registered new client: %d", client.ID())

})
