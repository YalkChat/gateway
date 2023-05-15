package chat

import (
	"math/rand"
	"time"

	"gorm.io/gorm"
)

// Messages     map[string]*MessageEvent `gorm:"type:hstore;messages" json:"messages"`

type Chat struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
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

type ChatUser struct {
	ID     uint
	ChatID uint
	Chat   Chat
	UserID uint
	User   User
}

func GetChatUserIds(chatId uint, db *gorm.DB) ([]uint, error) {
	var chatUsers []*ChatUser

	result := db.Where("chat_id = ?", chatId).Find(&chatUsers)

	if result.Error != nil {
		return nil, result.Error
	}

	var userIds []uint

	for _, user := range chatUsers {
		if user.ChatID == chatId {
			userIds = append(userIds, user.UserID)
		}
	}

	return userIds, nil
}