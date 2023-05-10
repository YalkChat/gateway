package chat

import (
	"fmt"
	"time"
	"yalk-backend/logger"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key" json:"id,omitempty"`
	From      string    `gorm:"from" json:"from,omitempty"`
	To        string    `gorm:"to" json:"to,omitempty"`
	Type      string    `gorm:"type" json:"type,omitempty"`
	Timestamp time.Time `gorm:"timestamp" json:"timestamp,omitempty"`
	Content   string    `gorm:"content" json:"content,omitempty"`
}

func (message *Message) saveToDb(chatId string, db *gorm.DB) error {
	tx := db.Table(chatId).Create(message)
	if tx.Error != nil {
		return tx.Error
	}
	logger.Info("MSG", fmt.Sprintf("Rows affected: %d", tx.RowsAffected))
	return nil
}
