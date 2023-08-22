package dbmodels

import (
	"encoding/json"

	"gorm.io/gorm"
)

type ChatType struct {
	ID   uint   `json:"id"`
	Type string `json:"type"`
}

type Chat struct {
	ID          uint      `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	ChatTypeID  uint      `json:"chatTypeID,omitempty"`
	ChatType    *ChatType `json:"chatType"`
	CreatedByID uint      `json:"createdByID,omitempty"`
	CreatedBy   *User     `json:"createdBy,omitempty"`
	// CreatedAt   time.Time  `json:"createdAt,omitempty"`
	Users    []*User    `gorm:"many2many:chat_users;" json:"users,omitempty"`
	Messages []*Message `json:"messages"`
}

func (chat *Chat) Type() string {
	return "chat"
}

func (chat *Chat) Serialize() ([]byte, error) {
	return json.Marshal(chat)
}

func (chat *Chat) Deserialize(data []byte) error {
	return json.Unmarshal(data, chat)
}

func (chat *Chat) GetInfo(db *gorm.DB) (*Chat, error) {

	tx := db.Preload("Users").Preload("Messages").Preload("ChatType").Find(&chat)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return chat, nil
}

// TODO: Move to Server method
func (chat *Chat) GetUsers(db *gorm.DB) ([]*User, error) {
	tx := db.Preload("Users").Find(&chat)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return chat.Users, nil
}

func (chat *Chat) Create(db *gorm.DB) error {
	return db.Preload("Users").Preload("ChatType").Create(&chat).Error
}
