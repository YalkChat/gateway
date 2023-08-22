package initialize

import (
	"yalk/database/dbmodels"

	"gorm.io/gorm"
)

func checkIsInitialized(db *gorm.DB) bool {
	var serverSettings *dbmodels.ServerSettings
	tx := db.Select("is_initialized").First(&serverSettings, "is_initialized = true")
	return tx.Error == nil
}
