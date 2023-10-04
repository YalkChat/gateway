package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
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
	existingSession, err := sessionsManager.Validate(r)
	if err != nil {
		errors.HandleError(w, r, err)
		return
	}

	if existingSession != nil {
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
	userID, err := srv.AuthenticateUser(*userLogin)
	if err != nil {
		errors.HandleError(w, r, errors.ErrAuthInvalid) // TODO: Add in errors package
	}

	sessionLenghtInt, err := strconv.Atoi(config.SessionLenght)
	if err != nil {
		errors.HandleError(w, r, errors.ErrInternalServerError)
	}
	sessionLenght := time.Minute * time.Duration(sessionLenghtInt)

	// Create a new session
	session, err := sessionsManager.Create(userID, sessionLenght)
	if err != nil {
		errors.HandleError(w, r, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    session.Token,
		Expires:  session.ExpiresAt,
		HttpOnly: true,
	})

	response := map[string]string{"status": "success", "message": "Succesfully signed in"}

	serializedResponse, err := json.Marshal(response)
	if err != nil {
		errors.HandleError(w, r, errors.ErrInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(serializedResponse)

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
