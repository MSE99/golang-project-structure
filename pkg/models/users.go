package models

import (
	"context"
	"database/sql"
	"log"

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

		users = append(users, u)
	}

	return users, nil
}

func InsertUser(ctx context.Context, tx *sql.Tx, u User) error {
	q, args, err := goqu.Insert("users").Cols("username", "password").Vals([]any{u.Username, u.Password}).ToSQL()
	if err != nil {
		return err
	}

	log.Println(q, args)

	_, err = tx.ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
