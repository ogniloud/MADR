package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/charmbracelet/log"
)

var ErrUserNotFound = fmt.Errorf("user not found")

// Database is a struct that represents database.
type Database struct {
	l  *log.Logger
	db *sql.DB
}

// New returns new database.
func New(l *log.Logger, db *sql.DB) *Database {
	return &Database{
		l:  l,
		db: db,
	}
}

// HasEmailOrUsername checks whether a user with the given username or email exists.
func (d *Database) HasEmailOrUsername(ctx context.Context, username, email string) (bool, error) {
	var count int

	err := d.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM user_credentials WHERE username = $1 OR email = $2", username, email).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("unable to count rows in HasEmailOrUsername: %w", err)
	}

	if count == 0 {
		return false, nil
	}
	return true, nil
}

// InsertUser inserts a new user into the database.
func (d *Database) InsertUser(ctx context.Context, username, salt, hash, email string) error {
	_, err := d.db.ExecContext(ctx, "INSERT INTO user_credentials (username, salt, hash, email) VALUES ($1, $2, $3, $4)", username, salt, hash, email)
	if err != nil {
		return fmt.Errorf("unable to insert user in InsertUser: %w", err)
	}

	return nil
}

// GetSaltAndHash returns salt and hash for the given username.
func (d *Database) GetSaltAndHash(ctx context.Context, username string) (salt, hash string, err error) {
	err = d.db.QueryRowContext(ctx, "SELECT salt, hash FROM user_credentials WHERE username = $1", username).Scan(&salt, &hash)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", "", ErrUserNotFound
		}

		return "", "", fmt.Errorf("unable to get salt and hash in GetSaltAndHash: %w", err)
	}

	return salt, hash, nil
}
