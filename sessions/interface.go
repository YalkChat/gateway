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
	Create(userID uint, defaultLenght time.Time) (*Session, error)
	Validate(*http.Request) (*Session, error)
	Delete(string) error
}
