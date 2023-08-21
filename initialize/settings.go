package initialize

import (
	"yalk/chat/models"

	"gorm.io/gorm"
)

func saveInitialSettings(db *gorm.DB) error {
	serverSettings := &models.ServerSettings{IsInitialized: true}
	err := serverSettings.Create(db)
	if err != nil {
		return err
	}
	return nil
}
