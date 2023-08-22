package dbmodels

import "gorm.io/gorm"

type ServerSettings struct {
	gorm.Model
	IsInitialized bool `json:"is_initialized"`
}

func (s *ServerSettings) Create(db *gorm.DB) error {
	return db.Create(s).Error
}

func (s *ServerSettings) Update(db *gorm.DB) error {
	return db.Save(s).Error
}
