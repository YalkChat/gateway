package events

type Status struct {
	// gorm.Model
	Name  string `json:"name"`
	Color string `json:"color"`
}
