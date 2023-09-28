package handlers

import (
	"net/http"
	"yalk/app"
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
		errors.HandleError(w, r, errors.ErrSessionValidation)
		return
	}

	if err != nil && err != errors.ErrCookieMissing {
		errors.HandleError(w, r, err)
	}

})
