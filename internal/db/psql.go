package db

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/jackc/pgx/v5"
)

// PSQLDatabase accesses to PostgresSQL with other fields.
type PSQLDatabase struct {
	*pgx.Conn
	logger *log.Logger
}

// NewPSQLDatabase creates a new connection to db and tries to connect to it.
func NewPSQLDatabase(ctx context.Context, connString string, logger *log.Logger) (*PSQLDatabase, error) {
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}

	return &PSQLDatabase{
		Conn:   conn,
		logger: logger,
	}, nil
}

func (db *PSQLDatabase) Logger() *log.Logger {
	return db.logger
}
