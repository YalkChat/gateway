package chat

import (
	"fmt"
	"time"
	"yalk/logger"

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

	if err := createUserTable(gorm); err != nil {
		logger.Warn("DB", fmt.Sprintf("Failed to migrate user profile table: %v", err))
	}

	if err := createChatTable(gorm); err != nil {
		logger.Warn("DB", fmt.Sprintf("Failed to migrate chat table: %v", err))
	}

	return gorm, nil
}

func createUserTable(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	testUser := &User{
		Username:      "TestUser",
		Email:         "test@test.com",
		DisplayedName: "TestDisplayedName",
		AvatarUrl:     "/testpp.png",
		IsOnline:      false,
		LastLogin:     time.Now(),
		LastOffline:   time.Now(),
		IsAdmin:       false,
	}
	if err := testUser.Create(db); err != nil {

		return err

	}
	return nil
}

func createChatTable(gorm *gorm.DB) error {
	// testMsg := make(map[string]*Message)
	// testMsg["test1"] = &Message{ID: "test1", From: "test", To: 1, Type: "channel_pub", Data: "Eh basta porco dio"}

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
