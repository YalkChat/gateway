package handlers

import (
	"strconv"
	"time"
	"yalk/app"
	"yalk/chat/server"
	"yalk/config"
	"yalk/sessions"
)

func getContextComponents(ctx app.HandlerContext) (server.Server, sessions.SessionManager, *config.Config) {
	return ctx.ChatServer(), ctx.SessionsManager(), ctx.Config()
}

func convertSessionLenght(sessionLenghtString string) (time.Duration, error) {
	sessionLenghtInt, err := strconv.Atoi(sessionLenghtString)
	if err != nil {
		return 0, err

	}
	sessionLenght := time.Minute * time.Duration(sessionLenghtInt)
	return sessionLenght, nil
}
