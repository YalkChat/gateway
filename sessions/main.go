package sessions

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type Credentials struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	jwt.StandardClaims
}

type Token = string
type UserID = uint
type SessionLenght = time.Time

func New(db *gorm.DB, dl time.Duration) *Manager {
	db.AutoMigrate(&Credentials{}, &Session{})
	return &Manager{
		defaultLenght:  dl,
		activeSessions: make([]*Session, 0),
	}
}

type Session struct {
	ID      uint
	UserID  UserID
	Token   string `gorm:"primaryKey"`
	Created time.Time
	Expiry  time.Time
}

func (s *Session) SetClientCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    "YLK", // TODO: Put SessionManager name
		Value:   s.Token,
		Expires: s.Expiry,
		// SameSite: http.SameSiteNoneMode,
		// Secure:   true, //! SET AGAIN WHEN USING HTTPS
		HttpOnly: true,
		Path:     "/",
	})
}

func (s *Session) isExpired() bool {
	return s.Expiry.Before(time.Now())
}
