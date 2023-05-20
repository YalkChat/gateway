package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"yalk/cattp"
	"yalk/chat"
	"yalk/sessions"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var signinHandle = cattp.HandlerFunc[*chat.Server](func(w http.ResponseWriter, r *http.Request, context *chat.Server) {
	defer r.Body.Close()

	dbSession, err := context.SessionsManager.Validate(context.Db, r, "YLK")

	if err == nil {
		// TODO: Extend session upon device validation
		log.Println("Session found - redirecting to app")
		dbSession.SetClientCookie(w)
		w.WriteHeader(http.StatusAccepted)
		return
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var login *sessions.Credentials
	// payload, err := io.ReadAll(r.Body)
	err = json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		panic(err)
	}

	// if err != nil {
	// 	log.Println("Can't get credentials from DB, wrong Email/Username")
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	dbCredentials := &sessions.Credentials{}

	if tx := context.Db.First(&dbCredentials, "email = ?", login.Email); tx.Error != nil {
		return
	}

	if login.Password == "" {
		log.Println("Empty password")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbCredentials.Password), []byte(login.Password))
	if err != nil {
		log.Println("Incorrect password")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// TODO: Move to method in session manager
	jwtKey := []byte("palle")
	claims := &sessions.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 720).Unix(),
		},
	}
	authToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := authToken.SignedString(jwtKey)
	if err != nil {
		return
	}

	session, err := context.SessionsManager.Create(context.Db, token, dbCredentials.ID, time.Time{})
	if err != nil {
		return
	}

	session.SetClientCookie(w) // TODO: Reimplement for JWT and WebSession
	if err != nil {
		log.Println("Error marshaling JWT Token")
		return
	}

	// TODO: Post response for WebSock?
	// payload, err := json.Marshal("/chat/1")
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	log.Printf("%s succesfully logged in.", dbCredentials.Email)
	// w.Write(payload)
	// http.Redirect(w, r, "/chat", http.StatusFound)
	w.WriteHeader(http.StatusOK)

})

var validateHandle = cattp.HandlerFunc[*chat.Server](func(w http.ResponseWriter, r *http.Request, context *chat.Server) {
	defer r.Body.Close()
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	session, err := context.SessionsManager.Validate(context.Db, r, "YLK")
	if err != nil {
		// TODO: Extend session upon device validation
		log.Println("Invalid session")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// session.SetClientCookie(w)
	// TODO: Post response for WebSock?
	w.Header().Add("Content-Type", "application/json")
	log.Printf("[%d][ID %v] Validated Session\n", http.StatusOK, session.UserID)
	w.WriteHeader(http.StatusOK)
})

var signoutHandle = cattp.HandlerFunc[*chat.Server](func(w http.ResponseWriter, r *http.Request, context *chat.Server) {
	defer r.Body.Close()

	cookie, err := r.Cookie("YLK")
	if err != nil {
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:   "YLK",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	err = context.SessionsManager.Delete(context.Db, cookie.Value) // TODO: Even just a property on the SessionManager is ok
	if err != nil {
		log.Println("Error deleting session", err)
	}
	log.Println("Signed out")
	w.WriteHeader(http.StatusOK)
})
