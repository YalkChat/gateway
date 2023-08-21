package handlers

import (
	"encoding/json"
	"yalk/chat/events"
	"yalk/chat/models"

	"gorm.io/gorm"
)

func makeInitialPayload(db *gorm.DB, user *models.User) ([]byte, error) {

	var chats *[]models.Chat
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

	var serverAccounts *[]models.Account
	if user.IsAdmin {
		tx = db.Find(&serverAccounts)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	var users *[]models.User
	tx = db.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	initialPayload := struct {
		User     *models.User      `json:"user"`
		Chats    *[]models.Chat    `json:"chats"`
		Accounts *[]models.Account `json:"accounts"`
		Users    *[]models.User    `json:"users"`
	}{user, chats, serverAccounts, users}

	jsonPayload, err := json.Marshal(&initialPayload)
	if err != nil {
		return nil, err
	}

	newRawEvent := &events.RawEvent{Type: "initial", Data: jsonPayload}

	jsonEvent, err := newRawEvent.Serialize()
	if err != nil {
		return nil, err
	}

	return jsonEvent, nil
}
