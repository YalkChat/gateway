package initialize

import (
	"yalk/database/dbmodels"

	"gorm.io/gorm"
)

// TODO: ChatType here has something wrong, I'm not sure why but it's wrong.
func createMainChannel(db *gorm.DB, adminUser *dbmodels.User, chatType *dbmodels.ChatType) (*dbmodels.Chat, error) {
	mainChat := &dbmodels.Chat{
		Name:      "Main",
		ChatType:  chatType,
		CreatedBy: adminUser,
		Users:     []*dbmodels.User{adminUser}}

	tx := db.Create(mainChat)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return mainChat, nil
}
