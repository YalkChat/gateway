package chat

import (
	"fmt"

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

	return gorm, nil
}

func CreateDbTables(db *gorm.DB) error {
	if err := db.AutoMigrate(&Account{}, &User{}, &Chat{}, &Message{}, &ServerSettings{}); err != nil {
		return err
	}
	return nil
}
