package sessions

import (
	"net/http"
	"time"
)

type SessionDatabase interface {
	SaveSession(token string, id uint, lenght time.Time) (*Session, error)
	LoadSession(string) (*Session, error)
	DeleteSession(string) error
}

type SessionManager interface {
	Create(userID uint, defaultLenght time.Duration) (*Session, error)
	Validate(r *http.Request) (*Session, error)
	Extend(session *Session, w http.ResponseWriter) error
	Delete(token string) error
}
