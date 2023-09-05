package sessions

import (
	"net/http"
	"time"
)

type SessionDatabase interface {
	SaveSession(token, uint, time.Time) (*Session, error)
	LoadSession(token) (*Session, error)
	DeleteSession(string) error
}

type SessionManager interface {
	Create(token, uint, time.Time) (*Session, error)
	Validate(*http.Request) (*Session, error)
	Delete(token) error
}
