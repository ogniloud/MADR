package usercred

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/ogniloud/madr/internal/db"
	"github.com/ogniloud/madr/internal/flashcards/models"
)

var ErrUserNotFound = fmt.Errorf("user not found")

// UserCredentials is a struct that represents database.
type UserCredentials struct {
	conn *db.PSQLDatabase
}

// GetUserInfo returns username, email and error by given userId.
func (d *UserCredentials) GetUserInfo(ctx context.Context, userId int) (string, string, error) {
	var username string
	var email string

	err := d.conn.QueryRow(ctx, "SELECT username, email FROM user_credentials WHERE user_id = $1", userId).Scan(&username, &email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", "", ErrUserNotFound
		}

		return "", "", fmt.Errorf("unable to get username and email from db in GetUserInfo: %w", err)
	}

	return username, email, nil
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
func (d *UserCredentials) InsertUser(ctx context.Context, username, salt, hash, email string) (int, error) {
	row := d.conn.QueryRow(ctx, "INSERT INTO user_credentials (username, salt, hash, email) VALUES ($1, $2, $3, $4) RETURNING user_id", username, salt, hash, email)
	var id int

	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("unable to insert user in InsertUser: %w", err)
	}

	_, err := d.conn.Exec(ctx, "INSERT INTO user_info (user_id, max_box) VALUES ($1, $2)", id, 4)
	if err != nil {
		return 0, fmt.Errorf("unable to insert user info: %w", err)
	}

	return id, nil
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

// GetUserID returns user id for the given username.
func (d *UserCredentials) GetUserID(ctx context.Context, username string) (int64, error) {
	var userID int64

	err := d.conn.QueryRow(ctx, "SELECT user_id FROM user_credentials WHERE username = $1", username).Scan(&userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrUserNotFound
		}

		return 0, fmt.Errorf("unable to get user id in GetUserID: %w", err)
	}

	return userID, nil
}

func (d *UserCredentials) ImportGoldenWords(ctx context.Context, userId int) error {
	tx, err := d.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		d.conn.Logger().Errorf("unable to start transaction in ImportGoldenWords: %v", err)
		return fmt.Errorf("unable to start transaction: %w", err)
	}

	for i := range 4 {
		row := tx.QueryRow(ctx, `
INSERT INTO deck_config(user_id, name)
SELECT $1, dc.name from deck_config dc
	WHERE dc.deck_id=$2
RETURNING deck_id;
`, userId, i)

		var deckId models.DeckId
		if err := row.Scan(&deckId); err != nil {
			d.conn.Logger().Errorf("unable to insert deck_config in ImportGoldenWords: %v", err)
			return fmt.Errorf("unable to insert deck config: %w", err)
		}

		_, err = tx.Exec(ctx, `
INSERT INTO flashcard(word, backside, deck_id, answer, multiple_backside)
SELECT f.word, f.backside, $1, f.answer, '[]'::jsonb || f.backside FROM flashcard f
    WHERE f.deck_id=$2;
`, deckId, i)

		_, err = tx.Exec(ctx, `
INSERT INTO user_leitner(user_id, card_id, box, cool_down)
SELECT $1, flashcard.card_id, 0, now() FROM flashcard
WHERE flashcard.deck_id=$2;
`, userId, deckId)
		if err != nil {
			d.conn.Logger().Errorf("unable to insert user_leitner in ImportGoldenWords: %v", err)
			return fmt.Errorf("unable to insert user leitner in ImportGoldenWords: %w", err)
		}

		_, err = tx.Exec(ctx, `
INSERT INTO copied_by(copier_id, deck_id, time_copied)
VALUES ($1, $2, now());
`, userId, i)
		if err != nil {
			d.conn.Logger().Errorf("unable to insert copied by in ImportGoldenWords: %v", err)
			return fmt.Errorf("unable to insert copied by in ImportGoldenWords: %w", err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		d.conn.Logger().Errorf("unable to commit transaction in ImportGoldenWords: %v", err)
		return fmt.Errorf("unable to commit transaction: %w", err)
	}

	return nil
}

func (d *UserCredentials) GetAllIds(ctx context.Context) ([]int, error) {
	rows, err := d.conn.Query(ctx, "SELECT user_id FROM user_credentials")
	if err != nil {
		d.conn.Logger().Errorf("unable to get all user credentials: %v", err)
		return nil, fmt.Errorf("unable to get all user credentials: %w", err)
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			d.conn.Logger().Errorf("unable to get all user credentials: %v", err)
			return nil, fmt.Errorf("unable to get all user credentials: %w", err)
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (d *UserCredentials) ImportGoldenWordsForOld(ctx context.Context) error {
	ids, err := d.GetAllIds(ctx)
	if err != nil {
		return fmt.Errorf("unable to get all user credentials: %w", err)
	}

	var errs error
	for _, userId := range ids {
		if userId == 0 {
			continue
		}
		row := d.conn.QueryRow(ctx, `
SELECT count(*) FROM copied_by 
    WHERE copier_id = $1 AND deck_id=0;
`, userId)
		var count int
		if err := row.Scan(&count); err != nil {
			errs = errors.Join(errs, err)
			d.conn.Logger().Errorf("unable to insert copied by in ImportGoldenWordsForOld: %v", err)
		}

		if count > 0 {
			continue
		}

		if err := d.ImportGoldenWords(ctx, userId); err != nil {
			errs = errors.Join(errs, err)
		}
	}

	return errs
}
