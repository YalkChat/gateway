package chat

import (
	"errors"
	"fmt"
	"time"
	"yalk/logger"

	"gorm.io/gorm"
)

type User struct {
	ID            uint      `gorm:"primaryKey" json:"userId"`
	Username      string    `gorm:"username" json:"username"`
	Email         string    `gorm:"email" json:"email"`
	DisplayedName string    `gorm:"displayedName" json:"displayName"`
	AvatarUrl     string    `gorm:"avatarUrl" json:"avatarUrl"`
	IsOnline      bool      `gorm:"isOnline" json:"isOnline"`
	LastLogin     time.Time `gorm:"lastLogin" json:"lastLogin"`
	LastOffline   time.Time `gorm:"lastOffline" json:"lastOffline"`
	IsAdmin       bool      `gorm:"isAdmin" json:"isAdmin"`
}

func (user *User) Create(db *gorm.DB) error {
	tx := db.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (user *User) Get(db *gorm.DB) error {
	if user.ID == 0 {
		return errors.New("no user ID")
	}

	tx := db.Find(&user)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func GetUserProfile(userId uint, db *gorm.DB) *User {
	var userProfile *User
	tx := db.Find(&User{}).Where("id = ?", userId).Scan(&userProfile)

	if tx.Error != nil {
		// TODO: Extend session upon device validation
		logger.Err("PROFILE", fmt.Sprintf("Error getting User Profile ID: %d\n", userId))
		return nil
	}

	return userProfile
}
