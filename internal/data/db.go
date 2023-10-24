package data

import (
	"fmt"

	"github.com/ogniloud/madr/internal/models"
)

var ErrEmailExists error = fmt.Errorf("user with this email already exists")

type Datalayer struct {
}

func NewDatalayer() *Datalayer {
	return &Datalayer{}
}

var users models.Users

func (d *Datalayer) isEmailExists(email string) bool {
	for _, user := range users {
		if user.Email == email {
			return true
		}
	}
	return false
}

func (d *Datalayer) CreateUser(user models.User) (models.User, error) {
	if d.isEmailExists(user.Email) {
		return models.User{}, ErrEmailExists
	}
	if len(users) == 0 {
		user.ID = 1
	} else {
		user.ID = users[len(users)-1].ID + 1
	}
	users = append(users, user)
	return user, nil
}
