package events

type UserCreationEvent struct {
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	StatusID  string `json:"status_id,omitempty"`
}
