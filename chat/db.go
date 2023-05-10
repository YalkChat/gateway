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

	err = createUserProfileTable(gorm)
	if err != nil {
		logger.Warn("DB", fmt.Sprintf("Failed to migrate user profile table: %v", err))
		// return nil, err
	}

	err = createChatTable(gorm)
	if err != nil {
		logger.Warn("DB", fmt.Sprintf("Failed to migrate chat table: %v", err))
		// return nil, err
	}

	return gorm, nil
}

func createUserProfileTable(gorm *gorm.DB) error {
	err := gorm.AutoMigrate(&UserProfile{})
	if err != nil {
		return err
	}
	return nil
}

func createChatTable(gorm *gorm.DB) error {
	// testMsg := make(map[string]*Message)
	// testMsg["test1"] = &Message{ID: "test1", From: "test", To: "MAIN", Type: "channel_pub", Data: "Eh basta porco dio"}

	// var chat = &Chat{ID: uuid.New(), Type: "channel_pub", Name: "NameTest", Users: []string{"test"}, Messages: testMap, Creator: "test", CreationDate: time.Now()}
	// chat :=
	// chat.ID = uuid.New()
	// chat.ID = uuid.NewString()

	err := gorm.AutoMigrate(&Chat{})
	if err != nil {
		return err
	}
	return nil
}
