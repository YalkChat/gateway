package handlers

import (
	"encoding/json"
	"net/http"
	"yalk/app"
	"yalk/chat/models/events"
	"yalk/errors"

	"github.com/AleRosmo/cattp"
)

var SigninHandler = cattp.HandlerFunc[app.HandlerContext](func(w http.ResponseWriter, r *http.Request, ctx app.HandlerContext) {
	defer r.Body.Close()

	if r.Method != http.MethodPost {
		errors.HandleError(w, r, errors.ErrInvalidMethodPost)
		return
	}

	srv, sessionsManager, config := getContextComponents(ctx)

	session, err := sessionsManager.Validate(r)
	if err != nil {
		handleSessionValidationError(w, r, err)
		return
	}

	var userLogin *events.UserLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userLogin)
	if err != nil {
		errors.HandleError(w, r, errors.ErrInvalidJson)
		return
	}

})

func handleSessionValidationError(w http.ResponseWriter, r *http.Request, err error) {
	if err == errors.ErrCookieMissing {
		errors.HandleError(w, r, errors.ErrSessionValidation)
	} else {

		errors.HandleError(w, r, err)
	}
	return
}
