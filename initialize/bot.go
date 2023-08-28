package initialize

import (
	"yalk/chat/models"

	"gorm.io/gorm"
)

func createBotAccount(db *gorm.DB) (*models.Account, error) {
	botAccount := &models.Account{
		Email:    "invalid@example.com",
		Username: "bot",
		Password: "none",
		Verified: false}

	if err := botAccount.Create(db); err != nil {
		return nil, err
	}
	return botAccount, nil
}

func createBotUser(db *gorm.DB, botAccount *models.Account) (*models.User, error) {
	botUser := &models.User{
		DisplayedName: "Bot",
		AvatarUrl:     "/bot.png",
		Account:       botAccount,
		StatusName:    "bot"}

	if err := botUser.Create(db); err != nil {
		return nil, err
	}
	return botUser, nil
}
