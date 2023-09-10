package db

type ChatType struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}
