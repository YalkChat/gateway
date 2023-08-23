package handlers

import (
	"log"
	"yalk/chat/chatmodels"
	"yalk/database/dbmodels"

	"gorm.io/gorm"
)

func HandleAccount(rawEvent *chatmodels.RawEvent, db *gorm.DB) (*dbmodels.Account, error) {
	account := &dbmodels.Account{}

	if err := account.Deserialize(rawEvent.Data); err != nil {
		log.Printf("Error Deserializing Account")
		return nil, err
	}

	if err := account.Create(db); err != nil {
		log.Printf("Error Creating Account: %d\n", account.ID)
		return nil, err
	}

	log.Printf("Account Created: %d\n", account.ID)
	return account, nil
}
