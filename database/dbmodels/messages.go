package dbmodels

import (
	"encoding/json"
	"log"
	"time"

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
	return "message"
}

func (message *Message) Serialize() ([]byte, error) {
	return json.Marshal(message)
}

func (message *Message) Deserialize(data []byte) error {
	return json.Unmarshal(data, message)
}

func (message *Message) SaveToDb(db *gorm.DB) error {
	tx := db.Create(message)
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	log.Printf("Rows affected: %d", tx.RowsAffected)
	return nil
}
