package chat

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Chat struct {
	ID           string              `gorm:"type:text;primary_key" json:"id"`
	MessageType  string              `gorm:"messageType" json:"type"`
	Name         string              `gorm:"name" json:"name"`
	Users        pq.StringArray      `gorm:"type:text[];users" json:"users"`
	Messages     map[string]*Message `gorm:"type:hstore;messages" json:"messages"`
	Creator      string              `gorm:"creator" json:"creator"`
	CreationDate time.Time           `gorm:"creationDate" json:"creationDate"`
}

func (Chat) TableName() string {
	return uuid.NewString()
}

// func createNewChat() {

// }
