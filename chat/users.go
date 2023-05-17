package chat

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint      `json:"userId"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	DisplayedName string    `json:"displayName"`
	AvatarUrl     string    `json:"avatarUrl"`
	IsOnline      bool      `json:"isOnline"`
	LastLogin     time.Time `json:"lastLogin"`
	LastOffline   time.Time `json:"lastOffline"`
	IsAdmin       bool      `json:"isAdmin"`
	Chats         []*Chat   `gorm:"many2many:chat_users;" json:"chats"`
}

// * We both return and assign to the user to allow method chaining.
func (user *User) Create(db *gorm.DB) (*User, error) {

	tx := db.Create(&user)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (user *User) GetInfo(db *gorm.DB) (*User, error) {

	tx := db.Preload("Chats").Preload("Chats.ChatType").Find(&user)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return user, nil
}

func (user *User) GetJoinedChats(db *gorm.DB) ([]*Chat, error) {
	var chats []*Chat

	tx := db.Preload("Chats").Find(&chats)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return chats, nil
}

func (user *User) CheckValid() (*User, error) {
	if user.ID == 0 {
		return nil, errors.New("no user ID")
	}
	return user, nil
}
