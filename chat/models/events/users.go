package events

import (
	"time"
)

type User struct {
	ID            uint      `json:"userId,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
	DeletedAt     time.Time `json:"deletedAt,omitempty"`
	Email         string    `json:"email,omitempty"`
	DisplayedName string    `json:"displayName,omitempty"`
	AvatarUrl     string    `json:"avatarUrl,omitempty"`
	StatusID      string    `json:"statusName,omitempty"`
	Status        *Status   `json:"status,omitempty"`
	CustomStatus  string    `json:"custom_status,omitempty"`
	LastLogin     time.Time `json:"lastLogin,omitempty"`
	LastOffline   time.Time `json:"lastOffline,omitempty"`
	IsAdmin       bool      `json:"isAdmin,omitempty"`
	Chats         []*Chat   `json:"chats,omitempty"`
	// IsOnline      bool      `json:"isOnline,omitempty"` // TODO: If we use status do we need it?
}
