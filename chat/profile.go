package chat

import (
	"fmt"
	"time"
	"yalk-backend/logger"

	"gorm.io/gorm"
)

type UserProfile struct {
	ID            int       `gorm:"primaryKey" json:"user_id"`
	Username      string    `gorm:"username" json:"username"`
	Email         string    `gorm:"email" json:"email"`
	DisplayedName string    `gorm:"displayed_name" json:"display_name"`
	IsOnline      bool      `gorm:"isOnline" json:"isOnline"`
	LastLogin     time.Time `gorm:"last_login" json:"lastLogin"`
	LastOffline   time.Time `gorm:"last_offline" json:"lastOffline"`
	IsAdmin       bool      `gorm:"isAdmin" json:"isAdmin"`
}

func getUserProfile(userId string, db *gorm.DB) *UserProfile {
	var userProfile *UserProfile
	tx := db.Find(&UserProfile{}).Where("id = ?", userId).Scan(&userProfile)

	if tx.Error != nil {
		// TODO: Extend session upon device validation
		logger.Err("PROFILE", fmt.Sprintf("Error getting User Profile ID: %s\n", userId))
		return nil
	}

	return userProfile
}
