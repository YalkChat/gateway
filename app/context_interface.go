// TODO: this file must be reviewed to where it should go in the file structure
package app

import (
	"yalk/chat/server"
	"yalk/config"
	"yalk/sessions"
)

type HandlerContext interface {
	ChatServer() server.Server
	SessionsManager() sessions.SessionManager
	Config() *config.Config
}
