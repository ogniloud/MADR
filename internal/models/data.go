package models

// User is a struct that defines the user model.
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Users is a struct that defines a slice of User.
type Users []User

// AuthorizationToken is a type that defines the authorization token.
type AuthorizationToken string
