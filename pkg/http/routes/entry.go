package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/mse99/golang-project-structure/pkg/http/controllers"
)

func MountRoutes(db *sql.DB, app *fiber.App) {
	router := app.Group("/api")
	userController := &controllers.UserController{DB: db}
	router.Get("/users", userController.Index)
}
