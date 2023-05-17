package chat

import (
	"time"

	"gorm.io/gorm"
)

type ChatType struct {
	ID   uint   `json:"id"`
	Type string `json:"type"`
}

type Chat struct {
	ID          uint       `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	ChatTypeID  uint       `json:"chatTypeID,omitempty"`
	ChatType    *ChatType  `json:"chatType"`
	CreatedByID uint       `json:"createdByID,omitempty"`
	CreatedBy   *User      `json:"createdBy,omitempty"`
	CreatedAt   time.Time  `json:"createdAt,omitempty"`
	Users       []*User    `gorm:"many2many:chat_users;" json:"users,omitempty"`
	Messages    []*Message `json:"messages,omitempty"`
}

func (chat *Chat) GetInfo(db *gorm.DB) (*Chat, error) {

	tx := db.Preload("Users").Preload("Messages").Preload("ChatType").Find(&chat)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return chat, nil
}
