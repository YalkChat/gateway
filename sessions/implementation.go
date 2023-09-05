package sessions

import (
	"time"

	"gorm.io/gorm"
)

type SessionDB struct{}

func NewDatabase() SessionDatabase {
	return &SessionDB{}
}

func (sdb *SessionDB) SaveSession(db *gorm.DB, token string, id uint, lenght time.Time) (*Session, error) {
	session := &Session{
		Token:     token,
		AccountID: id,
		ExpiresAt: lenght,
	}
	result := db.Create(session)
	if result.Error != nil {
		return nil, result.Error
	}
	return session, nil
}

func (sdb *SessionDB) LoadSession(db *gorm.DB, token string) (*Session, error) {
	var session Session
	result := db.Where("token = ?", token).First(&session)
	if result.Error != nil {
		return nil, result.Error
	}
	return &session, nil
}

func (sdb *SessionDB) DeleteSession(db *gorm.DB, token token) error {
	result := db.Where("token = ?", token).Delete(&Session{})
	return result.Error
}
