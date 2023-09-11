package events

type UserUpdateEvent struct {
	User     User   `json:"user"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}
