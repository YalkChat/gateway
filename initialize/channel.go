package initialize

import (
	"yalk/chat"

	"gorm.io/gorm"
)

// TODO: ChatType here has something wrong, I'm not sure why but it's wrong.
func createMainChannel(db *gorm.DB, adminUser *chat.User, chatType *chat.ChatType) (*chat.Chat, error) {
	mainChat := &chat.Chat{
		Name:      "Main",
		ChatType:  chatType,
		CreatedBy: adminUser,
		Users:     []*chat.User{adminUser}}

	tx := db.Create(mainChat)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return mainChat, nil
}
