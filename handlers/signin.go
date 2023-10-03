package handlers

import (
	"encoding/json"
	"net/http"
	"yalk/app"
	"yalk/chat/models/events"
	"yalk/errors"

	"github.com/AleRosmo/cattp"
	"golang.org/x/crypto/bcrypt"
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
		errors.HandleError(w, r, err)
		return
	}

	if session != nil {
		// Redirect to main chat page
		http.Redirect(w, r, "/chat", http.StatusFound)
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
	// TODO: Move this logic to errros package
	if err == errors.ErrCookieMissing {
		errors.HandleError(w, r, err)
	} else {

		errors.HandleError(w, r, err)
	}
	return
}

// Change return to user event if I need more info basides the user id
func authenticateUser(userLogin events.UserLogin) (string, error) {
	dbUser, err := srv.GetUserByUsername(userLogin.Username)
	if err != nil {
		return "", err
	}

	// Validate the password
	if !validatePassword(dbUser.Password, userLogin.Password) {
		return "", errors.ErrAuthInvalid
	}

	return dbUser.ID, nil
}

func validatePassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
