package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"yalk/app"
	"yalk/chat/models/events"
	"yalk/chat/server"
	"yalk/config"
	"yalk/errors"
	"yalk/sessions"

	"github.com/AleRosmo/cattp"
)

var SigninHandler = cattp.HandlerFunc[app.HandlerContext](func(w http.ResponseWriter, r *http.Request, ctx app.HandlerContext) {
	defer r.Body.Close()

	if r.Method != http.MethodPost {
		errors.HandleError(w, r, errors.ErrInvalidMethodPost)
		return
	}

	srv, sessionsManager, config := getContextComponents(ctx)

	err := validateSession(w, r, sessionsManager)
	if err != nil {
		errors.HandleError(w, r, err)
	}

	userID, err := authenticateUser(w, r, srv)
	if err != nil {
		errors.HandleError(w, r, err)
	}

	// Create a new session
	err = createSession(userID, w, r, sessionsManager, config)
	if err != nil {
		errors.HandleError(w, r, err)
	}

	err = sendResponse(w, r)
	if err != nil {
		errors.HandleError(w, r, err)
	}

})

func validateSession(w http.ResponseWriter, r *http.Request, sessionsManager sessions.SessionManager) error {
	// TODO: if the session is valid should redirect to main page?
	existingSession, err := sessionsManager.Validate(r)
	if err != nil {
		return errors.ErrSessionValidation
	}

	if existingSession != nil {
		// Redirect to main chat page
		// TODO: Is it correct to do the redirection here?
		return errors.ErrValidSessionExists
	}
	return nil
}

func authenticateUser(w http.ResponseWriter, r *http.Request, srv server.Server) (uint, error) {
	var userLogin *events.UserLogin

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&userLogin)
	if err != nil {
		return 0, errors.ErrInvalidJson
	}

	userID, err := srv.AuthenticateUser(*userLogin)
	if err != nil {
		return 0, errors.ErrAuthInvalid // TODO: Add in errors package
	}

	return userID, nil
}

func createSession(userID uint, w http.ResponseWriter, r *http.Request, sessionsManager sessions.SessionManager, config *config.Config) error {
	sessionLenghtInt, err := strconv.Atoi(config.SessionLenght)
	if err != nil {
		return errors.ErrInternalServerError
	}
	sessionLenght := time.Minute * time.Duration(sessionLenghtInt)

	session, err := sessionsManager.Create(userID, sessionLenght)
	if err != nil {
		return errors.ErrSessionCreation
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    session.Token,
		Expires:  session.ExpiresAt,
		HttpOnly: true,
	})
	return nil
}

func sendResponse(w http.ResponseWriter, r *http.Request) error {
	response := map[string]string{"status": "success", "message": "Succesfully signed in"}

	serializedResponse, err := json.Marshal(response)
	if err != nil {
		return errors.ErrInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(serializedResponse)
	return nil
}
