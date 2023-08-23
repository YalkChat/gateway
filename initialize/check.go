package initialize

import (
	"yalk/chat/models"

	"gorm.io/gorm"
)

func checkIsInitialized(db *gorm.DB) bool {
	var serverSettings *models.ServerSettings
	tx := db.Select("is_initialized").First(&serverSettings, "is_initialized = true")
	return tx.Error == nil
}
