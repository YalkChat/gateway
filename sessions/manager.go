package sessions

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type Claims struct {
	jwt.StandardClaims
}

type Token = string
type SessionLenght = time.Time

type SessionManager interface {
	Create(*sql.DB, Token, uint, SessionLenght) *Session
	Delete(*sql.DB, string) error
	Validate(*sql.DB, string) (*Session, error)
	Read(*sql.DB, string) (*Session, error)
}

func New(db *gorm.DB, dl time.Duration) *Manager {
	db.AutoMigrate(&Session{})
	return &Manager{
		defaultLenght:  dl,
		activeSessions: make([]*Session, 0),
	}
}

type Manager struct {
	defaultLenght  time.Duration
	activeSessions []*Session
}

// !!!!!!! JUST USE REDIS INSTEAD!!!!!!!
// TODO: Some decent in error checking would be nice
func (sm *Manager) Create(db *gorm.DB, token Token, id uint, lenght SessionLenght) (*Session, error) {
	var _lenght time.Time = time.Now().Add(sm.defaultLenght)

	if !lenght.IsZero() {
		_lenght = lenght
	}

	session := &Session{
		AccountID: id,
		Token:     token,
		Created:   time.Now(),
		Expiry:    _lenght,
	}

	if tx := db.Create(&session); tx.Error != nil {
		return nil, tx.Error
	}

	sm.activeSessions = append(sm.activeSessions, session)
	log.Println("Saved new session")
	return session, nil
}

func (sm *Manager) Validate(db *gorm.DB, r *http.Request, cookieName string) (*Session, error) {
	// TODO: This requires redis?
	// TODO: Further cookie properties check
	// ? If loading from DB data, this is 'redundant'
	// Server might have restarted, search DB and compare - fuck them they just log in again

	cookie, err := r.Cookie("YLK")
	if err != nil {
		return nil, errors.New("cookie_missing")
	}

	err = cookie.Valid()
	if err != nil {
		return nil, err
	}

	for _, session := range sm.activeSessions {
		if cookie.Value == session.Token {
			log.Println("Found stored active session")
			return session, nil
		}
	}

	session := &Session{}

	if tx := db.First(&session, "token = ?", cookie.Value); tx.Error != nil {
		return nil, tx.Error
	}

	log.Printf("Found session %s", session.Token)

	if session.isExpired() {
		sm.Delete(db, session.Token)

		sm.activeSessions[session.AccountID] = nil

		log.Println("Session expired, removed")
		return nil, errors.New("expired")
	}

	return session, nil
}

func (sm *Manager) Delete(db *gorm.DB, token Token) error {
	var session *Session
	if tx := db.Delete(&session, token); tx.Error != nil {
		return tx.Error
	}
	log.Printf("Deleted token %s", session.Token)
	return nil
}
