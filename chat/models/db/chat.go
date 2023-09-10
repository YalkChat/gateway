package db

import "gorm.io/gorm"

// Chat represents a chat room
type Chat struct {
	gorm.Model
	ChatTypeID uint
	ChatType   *ChatType
	Name       string  `gorm:"uniqueIndex"`
	Users      []*User `gorm:"many2many:chat_clients;"`
	Messages   []*Message
}
