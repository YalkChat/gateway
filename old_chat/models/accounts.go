package models

// import (
// 	"encoding/json"

// 	"gorm.io/gorm"
// )

// type Account struct {
// 	Verified bool `json:"verified"`
// }

// func (account *Account) Type() string {
// 	return "account"
// }

// func (account *Account) Serialize() ([]byte, error) {
// 	return json.Marshal(account)
// }

// func (account *Account) Deserialize(data []byte) error {
// 	return json.Unmarshal(data, account)
// }

// func (account *Account) Create(db *gorm.DB) error {
// 	return db.Create(&account).Error
// }

// func (account *Account) GetInfo(db *gorm.DB) error {
// 	return db.Find(&account).Error
// }
