package chat

import (
	"encoding/json"
	"fmt"
	"time"

	"yalk/logger"

	"gorm.io/gorm"
)

type Message struct {
	ID             uint      `json:"id,omitempty"`
	UserID         uint      `json:"user_id,omitempty"`
	User           User      `json:"user,omitempty"`
	From           string    `json:"from,omitempty"`
	ConversationID uint      `json:"conversationId,omitempty"` // convention to use it as Foreign Key
	Conversation   Chat      `json:"conversation,omitempty"`   // message belongs to conversation
	MessageType    string    `json:"type,omitempty"`
	Timestamp      time.Time `json:"timestamp,omitempty"`
	Content        string    `json:"content,omitempty"`
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
	tx := db.Table(fmt.Sprint(message.ConversationID)).Create(message)
	if tx.Error != nil {
		return tx.Error
	}
	logger.Info("MSG", fmt.Sprintf("Rows affected: %d", tx.RowsAffected))
	return nil
}
