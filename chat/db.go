package chat

import (
	"fmt"
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
		logger.Warn("DB", fmt.Sprintf("Failed to migrate users profile table: %v", err))
	}

	if err := createChatTable(gorm); err != nil {
		logger.Warn("DB", fmt.Sprintf("Failed to migrate chats table: %v", err))
	}
	if err := createMessagesTable(gorm); err != nil {
		logger.Warn("DB", fmt.Sprintf("Failed to migrate messages table: %v", err))
	}
	if err := createChatUsersTable(gorm); err != nil {
		logger.Warn("DB", fmt.Sprintf("Failed to migrate chat users table: %v", err))
	}

	return gorm, nil
}

func createUserTable(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	// testUser := &User{
	// 	Username:      "TestUser",
	// 	Email:         "test@test.com",
	// 	DisplayedName: "TestDisplayedName",
	// 	AvatarUrl:     "/testpp.png",
	// 	IsOnline:      false,
	// 	LastLogin:     time.Now(),
	// 	LastOffline:   time.Now(),
	// 	IsAdmin:       false,
	// }
	// if err := testUser.Create(db); err != nil {

	// 	return err

	// }
	return nil
}

func createChatTable(gorm *gorm.DB) error {
	err := gorm.AutoMigrate(&Chat{})
	if err != nil {
		return err
	}
	return nil
}

func createMessagesTable(db *gorm.DB) error {
	err := db.AutoMigrate(&Message{})
	if err != nil {
		return err
	}
	return nil
}

func createChatUsersTable(db *gorm.DB) error {
	err := db.AutoMigrate(&ChatUser{})
	if err != nil {
		return err
	}

	// user := &ChatUser{
	// 	ChatID: 1,
	// 	UserID: 1,
	// }

	// tx := db.Create(user)
	// if tx.Error != nil {
	// 	return tx.Error
	// }
	return nil
}
