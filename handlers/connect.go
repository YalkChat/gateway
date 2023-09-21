package handlers

import (
	"fmt"
	"net/http"
	"yalk/app"
	"yalk/chat/client"
	"yalk/chat/server"
	"yalk/config"

	"github.com/AleRosmo/cattp"
	"nhooyr.io/websocket"
)

func registerNewClient(server server.Server, conn *websocket.Conn, userID uint, config *config.Config) (client.Client, error) {
	client := client.NewClient(userID, conn, config.ClientTimeout) // TODO: time placeholder
	if err := server.RegisterClient(client); err != nil {

		return nil, ErrClientRegistration
	}
	return client, nil
}

// TODO: Handle Error must be finished
var ConnectionHandler = cattp.HandlerFunc[app.HandlerContext](func(w http.ResponseWriter, r *http.Request, ctx app.HandlerContext) {
	defer r.Body.Close()
	server := ctx.ChatServer()
	sessionsManager := ctx.SessionsManager()
	config := ctx.Config()

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

	userID := user.ID

	client, err := registerNewClient(server, conn, userID, config)
	if err != nil {
		handleError(w, r, ErrUserFetch)
		return
	}

	fmt.Printf("registered new client: %d", client.ID())
})
