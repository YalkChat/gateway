package handlers

import (
	"log"
	"yalk/chat/models"

	"gorm.io/gorm"
)

func HandleUser(rawEvent *models.RawEvent, db *gorm.DB) (*models.User, error) {
	var newUser = &models.User{ID: rawEvent.UserID}
	// var status = &Status{}
	if err := newUser.GetInfo(db); err != nil {
		log.Printf("Error getting user info ID: %d\n", newUser.ID)
		return nil, err
	}
	switch rawEvent.Action {
	case "change_status":

		// TODO: Change to status event type
		statusPayload := &models.User{}
		if err := statusPayload.Deserialize(rawEvent.Data); err != nil {
			log.Printf("Error Deserializing models.User")
			return nil, err
		}

		newUser.Status = &models.Status{Name: statusPayload.StatusName}

		if err := newUser.SaveToDb(db); err != nil {
			log.Printf("Error saving to DB models.User: %d\n", newUser.ID)
			return nil, err
		}
	}
	return newUser, nil
}

// TODO: Figure out where to keep models related functions and handlers (like in this case)
// XXX: This function was commented out when moving it for some reason
func handleUserCreate(rawEvent *models.RawEvent, db *gorm.DB, account *models.Account) (*models.User, error) {
	user := &models.User{Account: account}

	if err := user.Deserialize(rawEvent.Data); err != nil {
		log.Printf("Error Deserializing User")
		return nil, err
	}

	if err := user.Create(db); err != nil {
		log.Printf("Error Creating User: %d\n", user.ID)
		return nil, err
	}

	log.Printf("User Created: %d\n", user.ID)
	return user, nil
}
