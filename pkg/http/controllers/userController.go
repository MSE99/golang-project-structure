package controllers

import (
	"database/sql"
	"strconv"
	"sync"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/mse99/golang-project-structure/pkg/models"
	"github.com/mse99/golang-project-structure/pkg/views"
)

type UserController struct {
	DB *sql.DB
}

var (
	count     int = 0
	countLock sync.Mutex
)

func (ctl *UserController) Index(c *fiber.Ctx) error {
	countLock.Lock()
	defer countLock.Unlock()

	count++

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

	t := views.Home(users, strconv.Itoa(count))
	responder := templ.Handler(t)

	return adaptor.HTTPHandler(responder)(c)
}
