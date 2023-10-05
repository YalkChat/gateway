package handlers

import (
	"net/http"
	"yalk/app"
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

})
