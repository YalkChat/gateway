package sessions

import (
	"net/http"
	"time"
)

type UserID uint

type Session struct {
	ID      uint `gorm:"primaryKey"`
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
