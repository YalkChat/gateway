package initialize

import (
	"yalk/database/dbmodels"

	"gorm.io/gorm"
)

func saveInitialSettings(db *gorm.DB) error {
	serverSettings := &dbmodels.ServerSettings{IsInitialized: true}
	err := serverSettings.Create(db)
	if err != nil {
		return err
	}
	return nil
}
