// TODO: this file must be reviewed to where it should go in the file structure
package app

import (
	"gateway/config"
	"gateway/sessions"

	"github.com/AleRosmo/engine/server"
)

type HandlerContext interface {
	ChatServer() server.Server
	SessionsManager() sessions.SessionManager
	Config() *config.Config
}
