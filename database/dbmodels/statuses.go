package dbmodels

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Status struct {
	// gorm.Model
	Name  string `gorm:"primaryKey" json:"name"`
	Color string `json:"color"`
}

func (status *Status) Type() string {
	return "chat_message"
}

func (status *Status) Serialize() ([]byte, error) {
	return json.Marshal(status)
}

func (status *Status) Deserialize(data []byte) error {
	return json.Unmarshal(data, status)
}

func (s *Status) GetInfo(db *gorm.DB, statusName string) error {
	return db.First(&s, statusName).Error
}
