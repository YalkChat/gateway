package chat

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (account *Account) Type() string {
	return "account"
}

func (account *Account) Serialize() ([]byte, error) {
	return json.Marshal(account)
}

func (account *Account) Deserialize(data []byte) error {
	return json.Unmarshal(data, account)
}

func (c *Account) Create(db *gorm.DB) error {
	return db.Create(&c).Error
}
