package sessions

import (
	"net/http"
	"time"
)

type Session struct {
	ID        uint `gorm:"primaryKey"`
	AccountID uint
	Token     string `gorm:"primaryKey"`
	Created   time.Time
	ExpiresAt time.Time
}

func (s *Session) SetClientCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    "YLK", // TODO: Put SessionManager name
		Value:   s.Token,
		Expires: s.ExpiresAt,
		// SameSite: http.SameSiteNoneMode,
		// Secure:   true, //! SET AGAIN WHEN USING HTTPS
		// HttpOnly: true,
		Path: "/",
	})
}

func (s *Session) isExpired() bool {
	return s.ExpiresAt.Before(time.Now())
}
