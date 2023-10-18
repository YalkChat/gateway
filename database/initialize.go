package database

import (
	"fmt"

	"gateway/config" // XXX: I'm not entirely sure I should do this thing.

	"github.com/AleRosmo/engine/models/db"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDb(config *config.Config) (*gorm.DB, error) {

	db, err := OpenDbConnection(config)
	if err != nil {
		return nil, fmt.Errorf("failed to Open DB Connection")
	}

	if err := createDbTables(db); err != nil {
		return nil, fmt.Errorf("failed to AutoMigrate DB tables")
	}

	return db, nil
}

// XXX: This needs to be removed, ChatGateway should not interact directly with DB
func OpenDbConnection(config *config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName, config.DbSslMode)

	dbConn := postgres.Open(psqlInfo)

	gorm, err := gorm.Open(dbConn, &gorm.Config{QueryFields: true})
	if err != nil {
		return nil, err
	}

	return gorm, nil
}

func createDbTables(dbConn *gorm.DB) error {
	if err := dbConn.AutoMigrate(&db.User{}, &db.Chat{}, &db.Message{}, &db.ServerSettings{}, &db.Status{}); err != nil {
		return err
	}
	return nil
}
