package handlers

import (
	"encoding/json"
	"yalk/chat/chatmodels"
	"yalk/database/dbmodels"

	"gorm.io/gorm"
)

func MakeInitialPayload(db *gorm.DB, user *dbmodels.User) ([]byte, error) {

	var chats *[]dbmodels.Chat
	tx := db.Joins("left join chat_users on chat_users.chat_id=chats.id").
		Where("chat_users.user_id = ?", user.ID).
		Preload("Messages", func(db *gorm.DB) *gorm.DB {
			return db.Order("messages.timestamp ASC")
		}).
		Preload("Messages.User").
		Preload("ChatType").
		Find(&chats)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var serverAccounts *[]dbmodels.Account
	if user.IsAdmin {
		tx = db.Find(&serverAccounts)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	var users *[]dbmodels.User
	tx = db.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	initialPayload := struct {
		User     *dbmodels.User      `json:"user"`
		Chats    *[]dbmodels.Chat    `json:"chats"`
		Accounts *[]dbmodels.Account `json:"accounts"`
		Users    *[]dbmodels.User    `json:"users"`
	}{user, chats, serverAccounts, users}

	jsonPayload, err := json.Marshal(&initialPayload)
	if err != nil {
		return nil, err
	}

	newRawEvent := &chatmodels.RawEvent{Type: "initial", Data: jsonPayload}

	jsonEvent, err := newRawEvent.Serialize()
	if err != nil {
		return nil, err
	}

	return jsonEvent, nil
}
