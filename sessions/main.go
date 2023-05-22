package sessions

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type Account struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (account *Account) Type() string {
	return "account"
}

func (account *Account) Serialize() ([]byte, error) {
	return json.Marshal(account)
}

func (account *Account) Deserialize(data []byte) error {
	return json.Unmarshal(data, account)
}

func (c *Account) Create(db *gorm.DB) error {
	return db.Create(&c).Error
}

type Claims struct {
	jwt.StandardClaims
}

type Token = string
type UserID = uint
type SessionLenght = time.Time

func New(db *gorm.DB, dl time.Duration) *Manager {
	db.AutoMigrate(&Account{}, &Session{})
	return &Manager{
		defaultLenght:  dl,
		activeSessions: make([]*Session, 0),
	}
}

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
