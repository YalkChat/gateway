package chats

import (
	"math/rand"
	"time"
	"yalk/chat/events"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

// Messages     map[string]*MessageEvent `gorm:"type:hstore;messages" json:"messages"`

type Chat struct {
	ID        uint
	Name      string    `json:"name"`
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
}

type ChatUsers struct {
	ID     uint
	ChatID uint
}

type ConversationEvent struct {
	events.Event
	Users pq.StringArray `gorm:"type:text[];users" json:"users"`
}

func (c *Chat) TableName() uint {
	c.ID = uint(rand.Uint32())
	return c.ID
}

func (c *Chat) BeforeCreate(db *gorm.DB) error {
	// Updating the general table list
	// chatListEntry := &ChatList{}
	// db.Table("chat_list").Create()
	return nil
}

// func getJoinedUser(db *gorm.DB) []*UserProfile {

// }
