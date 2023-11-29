package usercred

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ogniloud/madr/internal/db"
)

var ErrUserNotFound = fmt.Errorf("user not found")

// UserCredentials is a struct that represents database.
type UserCredentials struct {
	conn *db.PSQLDatabase
}

// New returns new database.
func New(conn *db.PSQLDatabase) *UserCredentials {
	return &UserCredentials{
		conn: conn,
	}
}

// HasEmailOrUsername checks whether a user with the given username or email exists.
func (d *UserCredentials) HasEmailOrUsername(ctx context.Context, username, email string) (bool, error) {
	var count int

	err := d.conn.QueryRow(ctx, "SELECT COUNT(*) FROM user_credentials WHERE username = $1 OR email = $2", username, email).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("unable to count rows in HasEmailOrUsername: %w", err)
	}

	if count == 0 {
		return false, nil
	}
	return true, nil
}

// InsertUser inserts a new user into the database.
func (d *UserCredentials) InsertUser(ctx context.Context, username, salt, hash, email string) error {
	_, err := d.conn.Exec(ctx, "INSERT INTO user_credentials (username, salt, hash, email) VALUES ($1, $2, $3, $4)", username, salt, hash, email)
	if err != nil {
		return fmt.Errorf("unable to insert user in InsertUser: %w", err)
	}

	return nil
}

// GetSaltAndHash returns salt and hash for the given username.
func (d *UserCredentials) GetSaltAndHash(ctx context.Context, username string) (salt, hash string, err error) {
	err = d.conn.QueryRow(ctx, "SELECT salt, hash FROM user_credentials WHERE username = $1", username).Scan(&salt, &hash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", "", ErrUserNotFound
		}

		return "", "", fmt.Errorf("unable to get salt and hash in GetSaltAndHash: %w", err)
	}

	return salt, hash, nil
}
