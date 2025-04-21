package controllers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/mse99/golang-project-structure/pkg/models"
)

type UserController struct {
	DB *sql.DB
}

func (ctl *UserController) Index(c *fiber.Ctx) error {
	tx, err := ctl.DB.BeginTx(c.UserContext(), nil)
	if err != nil {
		return c.Status(500).Send([]byte("Something bad happened"))
	}
	defer tx.Rollback()

	users, err := models.LoadAllUsers(c.UserContext(), tx)
	if err != nil {
		return c.Status(500).Send([]byte("Something bad happened"))
	}

	for idx := range users {
		users[idx].Password = ""
	}

	return c.Status(200).JSON(users)
}
