package handlers

import (
	"log"
	"net/http"
	"yalk/app"

	"github.com/AleRosmo/cattp"
)

func clearCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   "YLK",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}

var SignoutHandle = cattp.HandlerFunc[app.HandlerContext](func(w http.ResponseWriter, r *http.Request, ctx app.HandlerContext) {
	defer r.Body.Close()

	sessionManager := ctx.SessionsManager()

	cookie, err := r.Cookie("YLK")
	if err != nil {
		handleError(w, r, http.ErrNoCookie)
		return
	}
	clearCookie(w)

	if err = sessionManager.Delete(cookie.Value); err != nil {
		handleError(w, r, ErrSessionDeletion)
		return
	}
	log.Println("Signed out")
	w.WriteHeader(http.StatusOK)

})
