package messages

import (
	"fmt"
	"time"

	"yalk/chat/chats"
	"yalk/chat/users"
	"yalk/logger"

	"gorm.io/gorm"
)

type Message struct {
	ID             uint       `json:"id,omitempty"`
	UserID         uint       `json:"user_id,omitempty"`
	User           users.User `json:"user,omitempty"`
	From           string     `json:"from,omitempty"`
	ConversationID uint       `json:"conversationId,omitempty"` // convention to use it as Foreign Key
	Conversation   chats.Chat `json:"conversation,omitempty"`   // message belongs to conversation
	Type           string     `json:"type,omitempty"`
	Timestamp      time.Time  `json:"timestamp,omitempty"`
	Content        string     `json:"content,omitempty"`
}

func (message *Message) SaveToDb(chatId uint, db *gorm.DB) error {
	tx := db.Table(fmt.Sprint(chatId)).Create(message)
	if tx.Error != nil {
		return tx.Error
	}
	logger.Info("MSG", fmt.Sprintf("Rows affected: %d", tx.RowsAffected))
	return nil
}
