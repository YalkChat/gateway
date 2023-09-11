package initialize

import (
	"yalk/chat/models/events"

	"gorm.io/gorm"
)

// TODO: ChatType here has something wrong, I'm not sure why but it's wrong.
// TODO: Missing method in DatabaseOperations
func createMainChannel(db *gorm.DB, adminUser *events.User, chatType *events.ChatType) (*events.Chat, error) {
	mainChat := &db.Chat{
		Name:      "Main",
		ChatType:  chatType,
		CreatedBy: adminUser,
		Users:     []*events.User{adminUser}}

	return mainChat, nil
}
