package db

import "time"

// Chat represents a chat room
type Chat struct {
	ID        string    `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Name      string    `gorm:"uniqueIndex"`
	Clients   []Client  `gorm:"many2many:chat_clients;"`
}
