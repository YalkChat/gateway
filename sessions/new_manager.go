package sessions

import (
	"fmt"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type sessionManagerImpl struct {
	defaultLenght  time.Duration
	activeSessions []*Session
	db             SessionDatabase
}

func NewSessionManager(db SessionDatabase) SessionManager {
	return &sessionManagerImpl{db: db}
}

// TODO: Redis will be used for this at some point in development.
// TODO: Some decent in error checking would be nice.
func (sm *sessionManagerImpl) Create(db *gorm.DB, token token, id uint, lenght time.Time) (*Session, error) {
	var _lenght time.Time = time.Now().Add(sm.defaultLenght)

	if !lenght.IsZero() {
		// return nil, fmt.Errorf("provided lenght is zero")
		_lenght = lenght
	}

	return sm.db.SaveSession(db, token, id, _lenght)
}

func (sm *sessionManagerImpl) Validate(db *gorm.DB, r *http.Request, cookieName string) (*Session, error) {
	// TODO: This requires redis?
	// TODO: Further cookie properties check
	// ? If loading from DB data, this is 'redundant'
	// Server might have restarted, search DB and compare - fuck them they just log in again
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return nil, fmt.Errorf("missing cookie")
	}

	return sm.db.LoadSession(db, cookie.Value)
}

func (sm *sessionManagerImpl) Delete(db *gorm.DB, token token) error {
	return sm.db.DeleteSession(db, token)
}
