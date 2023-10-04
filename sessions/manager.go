package sessions

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"
	"yalk/encryption"
	"yalk/errors"
)

type sessionManagerImpl struct {
	defaultLenght     time.Duration
	activeSessions    []*Session //TODO: Seems unused, remove
	db                SessionDatabase
	encryptionService encryption.Service
	cookieName        string
}

func NewSessionManager(db SessionDatabase, encryptionService encryption.Service, lenght time.Duration, cookieName string) SessionManager {
	return &sessionManagerImpl{
		db:                db,
		defaultLenght:     lenght,
		cookieName:        cookieName,
		encryptionService: encryptionService,
	}
}

// TODO: Redis will be used for this at some point in development.
// TODO: Some decent in error checking would be nice.
func (sm *sessionManagerImpl) Create(userId uint, lenght time.Duration) (*Session, error) {
	tokenBytes := make([]byte, 16)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return nil, err
	}
	token := hex.EncodeToString(tokenBytes)

	var _lenght time.Time = time.Now().Add(sm.defaultLenght)

	// TODO: What to do with this?
	// if lenght != 0 {
	// 	// return nil, fmt.Errorf("provided lenght is zero")
	// 	_lenght = lenght
	// }

	return sm.db.SaveSession(token, userId, _lenght)
}

func (sm *sessionManagerImpl) Validate(r *http.Request) (*Session, error) {
	// TODO: This requires redis?
	// TODO: Further cookie properties check
	// ? If loading from DB data, this is 'redundant'
	// Server might have restarted, search DB and compare - fuck them they just log in again
	cookie, err := r.Cookie(sm.cookieName)
	if err != nil {
		return nil, errors.ErrCookieMissing // Use the custom error type
	}

	return sm.db.LoadSession(cookie.Value)
}

func (sm *sessionManagerImpl) Delete(token string) error {
	return sm.db.DeleteSession(token)
}

func (sm *sessionManagerImpl) Extend(session *Session, w http.ResponseWriter) error {
	session.ExpiresAt = time.Now().Add(sm.defaultLenght)

	// Serialize the session back to a token
	token, err := sm.tokenService.Serialize(session)
	if err != nil {
		return err
	}

	// Set the new token as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     sm.cookieName,
		Value:    token,
		Expires:  session.ExpiresAt,
		HttpOnly: true,
		Secure:   sm.secureCookie,
	})
	return nil
}
