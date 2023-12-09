package models

// User is a struct that defines the user model.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Users is a struct that defines a slice of User.
type Users []User

// AuthorizationToken is a type that defines the authorization token.
type AuthorizationToken string

// UserInfo is a struct that defines the user info model.
type UserInfo struct {
	ID       int
	Username string
	Email    string
}
