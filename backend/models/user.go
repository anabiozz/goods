package models

// User ..
type User struct {
	Username string
	Email    string
}

// NewUser ...
func NewUser(username string, email string) *User {
	return &User{
		Username: username,
		Email:    email,
	}
}
