package data

import (
	"context"
	"fmt"

	"github.com/ogniloud/madr/internal/models"
)

// ErrEmailExists is an error returned when a user with the given email already exists.
var ErrEmailExists = fmt.Errorf("user with this email already exists")

// Datalayer is a struct that helps us to interact with the data.
type Datalayer struct {
	// users is a slice of models.User
	// This is our in-memory database. We will replace this with a real database later.
	users models.Users
}

// New returns a new Datalayer struct.
func New() *Datalayer {
	return &Datalayer{}
}

// isEmailExists is a helper function to check if a user with the given email already exists.
// Returns true if the user exists, false otherwise.
func (d *Datalayer) isEmailExists(email string) bool {
	for _, user := range d.users {
		if user.Email == email {
			return true
		}
	}

	return false
}

// isPasswordCorrect is a helper function to check if the given password is correct for the given email.
func (d *Datalayer) isPasswordCorrect(email, password string) bool {
	for _, user := range d.users {
		if user.Email == email && user.Password == password {
			return true
		}
	}

	return false
}

// CreateUser is a function to create a new user.
func (d *Datalayer) CreateUser(_ context.Context, user models.User) error {
	if d.isEmailExists(user.Email) {
		return ErrEmailExists
	}

	if len(d.users) == 0 {
		user.ID = 1
	} else {
		user.ID = d.users[len(d.users)-1].ID + 1
	}

	d.users = append(d.users, user)

	return nil
}

// SignInUser is a function to sign in a user.
// It returns an authorization token if the user is signed in successfully.
func (d *Datalayer) SignInUser(_ context.Context, email, password string) (authorization string, err error) {
	if d.isPasswordCorrect(email, password) {
		authorization = "Bearer blablablaIMATOKENyouAREpoorBASTARD"
		return
	}

	return authorization, models.ErrUnauthorized
}
