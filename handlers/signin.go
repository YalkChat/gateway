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

	// TODO: if the session is valid should redirect to main page?
	session, err := sessionsManager.Validate(r)
	if err != nil {
		handleSessionValidationError(w, r, err)
		return
	}

	var userLogin *events.UserLogin
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&userLogin)
	if err != nil {
		errors.HandleError(w, r, errors.ErrInvalidJson) //TODO: Add in errors package
		return
	}
	userID, err := authenticateUser(*userLogin)
	if err != nil {
		errors.HandleError(w, r, errors.ErrAuthInvalid) // TODO: Add in errors package
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

// Change return to user event if I need more info basides the user id
func authenticateUser(userLogin events.UserLogin) (string, error) {
	dbUser, err := s.db.GetUserByUsername(userLogin.Username)
	if err != nil {
		return "", err
	}

	// Validate the password
	if !validatePassword(user.Password, loginPayload.Password) {
		return "", errors.New(errors.ErrInvalidPassword)
	}

	return user.ID, nil
}
