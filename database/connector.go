package database

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Connect(ctx context.Context, dsn string) (*sql.DB, error) {
	conn, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	if err := conn.PingContext(ctx); err != nil {
		return nil, err
	}

	return conn, nil
}
