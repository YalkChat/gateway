// TODO: this file must be reviewed to where it should go in the file structure
package context

import (
	"yalk/newchat/server"
	"yalk/sessions"
)

type HandlerContext interface {
	ChatServer() server.Server
	SessionManager() sessions.SessionManager
}
