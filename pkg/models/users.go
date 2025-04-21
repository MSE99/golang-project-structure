package models

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoadAllUsers(ctx context.Context, tx *sql.Tx) ([]User, error) {
	q, _, err := goqu.Select("username", "password").From("users").ToSQL()
	if err != nil {
		return nil, err
	}

	rows, err := tx.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		var u User

		err := rows.Scan(&u.Username, &u.Password)
		if err != nil {
			continue
		}
	}

	return users, nil
}
