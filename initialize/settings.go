package initialize

import (
	"yalk/chat"

	"gorm.io/gorm"
)

func saveInitialSettings(db *gorm.DB) error {
	serverSettings := &chat.ServerSettings{IsInitialized: true}
	err := serverSettings.Create(db)
	if err != nil {
		return err
	}
	return nil
}
