package initialize

import (
	"yalk/chat"

	"gorm.io/gorm"
)

func createBotAccount(db *gorm.DB) (*chat.Account, error) {
	botAccount := &chat.Account{
		Email:    "invalid@example.com",
		Username: "bot",
		Password: "none",
		Verified: false}

	if err := botAccount.Create(db); err != nil {
		return nil, err
	}
	return botAccount, nil
}

func createBotUser(db *gorm.DB, botAccount *chat.Account) (*chat.User, error) {
	botUser := &chat.User{
		DisplayedName: "Bot",
		AvatarUrl:     "/bot.png",
		Account:       botAccount,
		StatusName:    "bot"}

	if err := botUser.Create(db); err != nil {
		return nil, err
	}
	return botUser, nil
}
