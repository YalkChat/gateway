package handlers

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

var SigninHandle = cattp.HandlerFunc[*chat.Server](func(w http.ResponseWriter, r *http.Request, context *chat.Server) {
	defer r.Body.Close()

	dbSession, err := context.SessionsManager.Validate(context.Db, r, "YLK")

	if err != nil && err.Error() != "cookie_missing" {
		// logger.Err("SESS", "Validation failed")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if err == nil {
		// TODO: Extend session upon device validation
		log.Println("Session found - redirecting to app")
		dbSession.SetClientCookie(w)
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var login *chat.Account
	// payload, err := io.ReadAll(r.Body)
	err = json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		// logger.Err("SESS", fmt.Sprintf("Failed to decode login request: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbCredentials := &chat.Account{}

	if tx := context.Db.First(&dbCredentials, "email = ?", login.Email); tx.Error != nil {
		log.Println("Can't get credentials from DB, wrong Email/Username")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if login.Email == "" {
		log.Println("Empty Email")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if login.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Empty password")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbCredentials.Password), []byte(login.Password))
	if err != nil {
		log.Println("Incorrect password")
		w.WriteHeader(http.StatusBadRequest)
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
		// logger.Err("SESS", fmt.Sprintf("Error signing token: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	session, err := context.SessionsManager.Create(context.Db, token, dbCredentials.ID, time.Time{})
	if err != nil {
		// logger.Err("SESS", fmt.Sprintf("Error creating session: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	session.SetClientCookie(w) // TODO: Reimplement for JWT and WebSession
	if err != nil {
		// logger.Err("SESS", fmt.Sprintf("Error setting client cookie: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%s succesfully logged in.", dbCredentials.Email)
	// w.Write(payload)
	// http.Redirect(w, r, "/chat", http.StatusFound)
	w.WriteHeader(http.StatusOK)

})