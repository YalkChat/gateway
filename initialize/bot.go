package initialize

import (
	"yalk/database/dbmodels"

	"gorm.io/gorm"
)

func createBotAccount(db *gorm.DB) (*dbmodels.Account, error) {
	botAccount := &dbmodels.Account{
		Email:    "invalid@example.com",
		Username: "bot",
		Password: "none",
		Verified: false}

	if err := botAccount.Create(db); err != nil {
		return nil, err
	}
	return botAccount, nil
}

func createBotUser(db *gorm.DB, botAccount *dbmodels.Account) (*dbmodels.User, error) {
	botUser := &dbmodels.User{
		DisplayedName: "Bot",
		AvatarUrl:     "/bot.png",
		Account:       botAccount,
		StatusName:    "bot"}

	if err := botUser.Create(db); err != nil {
		return nil, err
	}
	return botUser, nil
}
