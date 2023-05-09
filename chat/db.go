package chat

import (
	"fmt"
	"yalk-backend/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgConf struct {
	IP       string
	Port     string
	User     string
	Password string
	DB       string
	SslMode  string
}

func OpenDbConnection(conf *PgConf) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", conf.IP, conf.Port, conf.User, conf.Password, conf.DB, conf.SslMode)
	db := postgres.Open(psqlInfo)
	gorm, err := gorm.Open(db, &gorm.Config{QueryFields: true})
	if err != nil {
		return nil, err
	}

	err = migrateUserProfileTable(gorm)
	if err != nil {
		logger.Warn("DB", fmt.Sprintf("Failed to migrate user profile table: %v", err))
		// return nil, err
	}

	return gorm, nil
}

func migrateUserProfileTable(gorm *gorm.DB) error {
	err := gorm.AutoMigrate(&UserProfile{})
	if err != nil {
		return err
	}
	return nil
}
