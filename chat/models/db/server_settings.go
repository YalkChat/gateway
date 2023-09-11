package db

import "gorm.io/gorm"

type ServerSettings struct {
	gorm.Model
	IsInitialized bool `gorm:"is_initialized"`
}
