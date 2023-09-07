package events

import (
	"time"
)

type User struct {
	ID            uint      `json:"userId"`
	AccountID     uint      `json:"accountId"`
	Account       *Account  `json:"account"`
	DisplayedName string    `json:"displayName"`
	AvatarUrl     string    `json:"avatarUrl"`
	IsOnline      bool      `json:"isOnline"`
	StatusName    string    `json:"statusName"`
	Status        *Status   `json:"status"`
	LastLogin     time.Time `json:"lastLogin"`
	LastOffline   time.Time `json:"lastOffline"`
	IsAdmin       bool      `json:"isAdmin"`
	Chats         []*Chat   `json:"chats"`
}
