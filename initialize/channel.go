package initialize

import (
	"yalk/chat/models"

	"gorm.io/gorm"
)

// TODO: ChatType here has something wrong, I'm not sure why but it's wrong.
func createMainChannel(db *gorm.DB, adminUser *models.User, chatType *models.ChatType) (*models.Chat, error) {
	mainChat := &models.Chat{
		Name:      "Main",
		ChatType:  chatType,
		CreatedBy: adminUser,
		Users:     []*models.User{adminUser}}

	tx := db.Create(mainChat)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return mainChat, nil
}
