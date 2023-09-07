package events

type Account struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"` // ! NOPE
	Verified bool   `json:"verified"`
}
