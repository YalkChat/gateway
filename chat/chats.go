package chat

import (
	"time"

	"gorm.io/gorm"
)

type Chat struct {
	ID          uint       `json:"id"`
	Name        string     `json:"name"`
	CreatedByID uint       `json:"createdByID"`
	CreatedBy   *User      `json:"createdBy"`
	CreatedAt   time.Time  `json:"createdAt"`
	Users       []*User    `gorm:"many2many:chat_users;" json:"users"`
	Messages    []*Message `json:"messages"`
}

func (chat *Chat) GetInfo(db *gorm.DB) (*Chat, error) {

	tx := db.Preload("Users").Preload("Messages").Find(chat)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return chat, nil
}
