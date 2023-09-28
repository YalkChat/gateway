package sessions

import (
	"fmt"
	"net/http"
	"time"
)

type sessionManagerImpl struct {
	defaultLenght  time.Duration
	activeSessions []*Session //TODO: Seems unused, remove
	db             SessionDatabase
	cookieName     string
}

func NewSessionManager(db SessionDatabase, lenght time.Duration, cookieName string) SessionManager {
	return &sessionManagerImpl{db: db, defaultLenght: lenght, cookieName: cookieName}
}

// TODO: Redis will be used for this at some point in development.
// TODO: Some decent in error checking would be nice.
func (sm *sessionManagerImpl) Create(token string, id uint, lenght time.Time) (*Session, error) {
	var _lenght time.Time = time.Now().Add(sm.defaultLenght)

	if !lenght.IsZero() {
		// return nil, fmt.Errorf("provided lenght is zero")
		_lenght = lenght
	}

	return sm.db.SaveSession(token, id, _lenght)
}

func (sm *sessionManagerImpl) Validate(r *http.Request) (*Session, error) {
	// TODO: This requires redis?
	// TODO: Further cookie properties check
	// ? If loading from DB data, this is 'redundant'
	// Server might have restarted, search DB and compare - fuck them they just log in again
	cookie, err := r.Cookie(sm.cookieName)
	if err != nil {
		return nil, fmt.Errorf("missing cookie")
	}

	return sm.db.LoadSession(cookie.Value)
}

func (sm *sessionManagerImpl) Delete(token string) error {
	return sm.db.DeleteSession(token)
}
