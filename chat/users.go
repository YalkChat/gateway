package chat

import (
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint      `json:"userId"`
	AccountID     uint      `json:"accountId"`
	Account       *Account  `json:"account"`
	DisplayedName string    `json:"displayName"`
	AvatarUrl     string    `json:"avatarUrl"`
	IsOnline      bool      `json:"isOnline"`
	LastLogin     time.Time `json:"lastLogin"`
	LastOffline   time.Time `json:"lastOffline"`
	IsAdmin       bool      `json:"isAdmin"`
	Chats         []*Chat   `gorm:"many2many:chat_users;" json:"chats"`
}

func (user *User) Type() string {
	return "chat_message"
}

func (user *User) Serialize() ([]byte, error) {
	return json.Marshal(user)
}

func (user *User) Deserialize(data []byte) error {
	return json.Unmarshal(data, user)
}

// * We both return and assign to the user to allow method chaining.
func (user *User) Create(db *gorm.DB) error {
	return db.Create(&user).Error
}

func (user *User) GetInfo(db *gorm.DB) error {
	return db.Preload("Chats").Preload("Chats.ChatType").Preload("Account").Find(&user).Error
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
