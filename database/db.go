package database

import (
	"fmt"
	"yalk/chat"
	"yalk/config" // ? I'm not entirely sure I should do this thing.

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDb(config *config.Config) (*gorm.DB, error) {

	db, err := openDbConnection(config)
	if err != nil {
		return nil, fmt.Errorf("failed to Open DB Connection")
	}

	if err := createDbTables(db); err != nil {
		return nil, fmt.Errorf("failed to AutoMigrate DB tables")
	}

	return db, nil
}

func openDbConnection(config *config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName, config.DbSslMode)

	db := postgres.Open(psqlInfo)

	gorm, err := gorm.Open(db, &gorm.Config{QueryFields: true})
	if err != nil {
		return nil, err
	}

	return gorm, nil
}

func createDbTables(db *gorm.DB) error {
	if err := db.AutoMigrate(&chat.Account{}, &chat.User{}, &chat.Chat{}, &chat.Message{}, &chat.ServerSettings{}, &chat.Status{}); err != nil {
		return err
	}
	return nil
}
