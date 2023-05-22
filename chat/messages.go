package chat

import (
	"encoding/json"
	"fmt"
	"time"

	"yalk/logger"

	"gorm.io/gorm"
)

type Message struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"userId"`
	User      *User     `json:"user"`
	ChatID    uint      `json:"chatId"` // convention to use it as Foreign Key
	Chat      *Chat     `json:"chat"`   // message belongs to conversation
	Timestamp time.Time `json:"timestamp"`
	Content   string    `json:"content"`
	// MessageType string    `json:"type,omitempty"`
}

func (message *Message) Type() string {
	return "chat_message"
}

func (message *Message) Serialize() ([]byte, error) {
	return json.Marshal(message)
}

func (message *Message) Deserialize(data []byte) error {
	return json.Unmarshal(data, message)
}

func (message *Message) SaveToDb(db *gorm.DB) error {
	message.UserID = 1 // ! Debug
	tx := db.Create(message)
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	logger.Info("MSG", fmt.Sprintf("Rows affected: %d", tx.RowsAffected))
	return nil
}
