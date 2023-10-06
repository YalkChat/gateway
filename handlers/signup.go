package handlers

import (
	"encoding/json"
	"net/http"
	"yalk/app"
	"yalk/chat/models/db"
	"yalk/errors"

	"github.com/AleRosmo/cattp"
)

var SignupHandler = cattp.HandlerFunc[app.HandlerContext](func(w http.ResponseWriter, r *http.Request, ctx app.HandlerContext) {
	defer r.Body.Close()

	srv, sessionsManager, config := getContextComponents(ctx)

	if r.Method != http.MethodPost {
		errors.HandleError(w, r, errors.ErrInvalidMethodPost)
		return
	}

	var newUser db.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		errors.HandleError(w, r, errors.ErrInvalidJson)
	}

	userID, err := srv.CreateUser(&newUser)
	if err != nil {
		errors.HandleError(w, r, errors.ErrUserCreation)
	}

	sessionLenght, err := convertSessionLenght(config.SessionLenght)
	if err != nil {
		errors.HandleError(w, r, errors.ErrInternalServerError)
		return
	}

	session, err := sessionsManager.Create(userID, sessionLenght)
	if err != nil {
		errors.HandleError(w, r, errors.ErrSessionCreation)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    session.Token,
		Expires:  session.ExpiresAt,
		HttpOnly: true,
	})

})
