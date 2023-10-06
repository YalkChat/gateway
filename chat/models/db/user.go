package db

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string
	Password     string
	Username     string
	AvatarUrl    string
	StatusID     string `gorm:"foreignKey:Name"`
	Status       *Status
	CustomStatus string
	LastLogin    time.Time
	LastOffline  time.Time
	IsAdmin      bool
	Chats        []*Chat `gorm:"many2many:chat_users;"`
}
