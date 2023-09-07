package events

type Chat struct {
	ID          uint      `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	ChatTypeID  uint      `json:"chatTypeID,omitempty"`
	ChatType    *ChatType `json:"chatType"`
	CreatedByID uint      `json:"createdByID,omitempty"`
	CreatedBy   *User     `json:"createdBy,omitempty"`
	// CreatedAt   time.Time  `json:"createdAt,omitempty"`
	Users    []*User    ` json:"users,omitempty"`
	Messages []*Message `json:"messages"`
}
