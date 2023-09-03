package sessions

import "database/sql"

type SessionDatabase interface {
	SaveSession(*sql.DB, Token, uint, SessionLenght) *Session
	LoadSession(*sql.DB, string) error
	DeleteSession(*sql.DB, string) (*Session, error)
}
