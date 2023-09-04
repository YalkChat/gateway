package sessions

import (
	"net/http"
	"time"

	"gorm.io/gorm"
)

type SessionDatabase interface {
	SaveSession(*gorm.DB, token, uint, time.Time) (*Session, error)
	LoadSession(*gorm.DB, token) (*Session, error)
	DeleteSession(*gorm.DB, string) error
}

type SessionManager interface {
	Create(*gorm.DB, token, uint, time.Time) (*Session, error)
	Validate(*gorm.DB, *http.Request, string) (*Session, error)
	Delete(*gorm.DB, token) error
}
