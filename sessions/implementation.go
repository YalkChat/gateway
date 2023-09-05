package sessions

import (
	"time"

	"gorm.io/gorm"
)

type SessionDB struct {
	db *gorm.DB
}

func NewDatabase(db *gorm.DB) SessionDatabase {
	return &SessionDB{db: db}
}

func (sdb *SessionDB) SaveSession(token string, id uint, lenght time.Time) (*Session, error) {
	session := &Session{
		Token:     token,
		AccountID: id,
		ExpiresAt: lenght,
	}
	result := sdb.db.Create(session)
	if result.Error != nil {
		return nil, result.Error
	}
	return session, nil
}

func (sdb *SessionDB) LoadSession(token string) (*Session, error) {
	var session Session
	result := sdb.db.Where("token = ?", token).First(&session)
	if result.Error != nil {
		return nil, result.Error
	}
	return &session, nil
}

func (sdb *SessionDB) DeleteSession(token token) error {
	result := sdb.db.Where("token = ?", token).Delete(&Session{})
	return result.Error
}
