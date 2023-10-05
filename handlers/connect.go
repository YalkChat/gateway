package handlers

import (
	"fmt"
	"net/http"
	"yalk/app"
	"yalk/chat/client"
	"yalk/chat/server"
	"yalk/config"
	"yalk/errors"

	"github.com/AleRosmo/cattp"
	"nhooyr.io/websocket"
)

func registerNewClient(server server.Server, conn *websocket.Conn, userID uint, config *config.Config) (client.Client, error) {
	client := client.NewClient(userID, conn, config.ClientTimeout) // TODO: time placeholder
	if err := server.RegisterClient(client); err != nil {

		return nil, errors.ErrClientRegistration
	}
	return client, nil
}

// TODO: Handle Error must be finished
// TODO: Break down in smaller functions
var ConnectionHandler = cattp.HandlerFunc[app.HandlerContext](func(w http.ResponseWriter, r *http.Request, ctx app.HandlerContext) {
	defer r.Body.Close()

	srv, sessionsManager, config := getContextComponents(ctx)

	if r.Method != http.MethodGet {
		errors.HandleError(w, r, errors.ErrInvalidMethodGet)
		return
	}

	session, err := sessionsManager.Validate(r)
	if err != nil {
		errors.HandleError(w, r, errors.ErrSessionValidation)
		return
	}
	// Upgrades to WebSocket
	// TODO: I still am not convinced of using Config(), besides the Config struct
	conn, err := srv.UpgradeHttpRequest(w, r, ctx.Config())
	if err != nil {
		errors.HandleError(w, r, errors.ErrWebSocketUpgrade)
		return
	}

	user, err := srv.GetUserByID(session.UserID)
	if err != nil {
		errors.HandleError(w, r, errors.ErrUserFetch)
		return
	}

	userID := user.ID

	client, err := registerNewClient(srv, conn, userID, config)
	if err != nil {
		errors.HandleError(w, r, errors.ErrUserFetch)
		return
	}

	fmt.Printf("registered new client: %d", client.ID())

	// TODO: Missing initial payload
})
