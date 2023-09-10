package initialize

import (
	"yalk/chat/models/events"

	"gorm.io/gorm"
)

func createBotAccount(db *gorm.DB) (*events.User, error) {
	botAccount := &events.User{
		Email:    "invalid@example.com",
		Username: "bot",
		Password: "none",
		Verified: false}

	if err := botAccount.Create(db); err != nil {
		return nil, err
	}
	return botAccount, nil
}

func createBotUser(db *gorm.DB, botAccount *events.User) (*events.User, error) {
	botUser := &events.User{
		DisplayedName: "Bot",
		AvatarUrl:     "/bot.png",
		StatusID:      "bot"}

	if err := botUser.Create(db); err != nil {
		return nil, err
	}
	return botUser, nil
}
