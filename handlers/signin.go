package handlers

import (
	"net/http"
	"yalk/app"

	"github.com/AleRosmo/cattp"
)

var SigninHandler = cattp.HandlerFunc[app.HandlerContext](func(w http.ResponseWriter, r *http.Request, ctx app.HandlerContext) {
	defer r.Body.Close()

	if r.Method != http.MethodPost {
		handleError(w, r, ErrInvalidMethodPost)
		return
	}

	srv, sessionsManager, config := getContextComponents(ctx)

	session, err := sessionsManager.Validate(r)
	if err != nil {
		handleError(w, r, ErrSessionValidation)
		return
	}

	if err != nil && err != ErrCookieMissing {
		handleError(w, r, err)
	}

})
