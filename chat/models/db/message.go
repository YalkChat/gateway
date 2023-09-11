package db

import (
	"time"
)

// Message represents a chat message
// TODO: Review GORM tags as ChatGPT doesn't know now I can skip some
// TODO: Rewrite with gorm.Model
type Message struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	ChatID    uint      `gorm:"index"`
	UserID    uint      `gorm:"index"`
	Content   string    `gorm:"type:text"`
}
