package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mse99/golang-project-structure/config"
	"github.com/mse99/golang-project-structure/database"
	"github.com/mse99/golang-project-structure/pkg/http/routes"
	"github.com/urfave/cli/v3"
)

func createMainCliCommand() *cli.Command {
	return &cli.Command{
		Name:        "example",
		Description: "This is an example golang project",
		Commands: []*cli.Command{
			createVersionCommand(),
			createDbCommands(),
			createServeCommand(),
		},
	}
}

func createVersionCommand() *cli.Command {
	return &cli.Command{
		Name:        "version",
		Description: "Prints the current version",
		Action: func(ctx context.Context, c *cli.Command) error {
			fmt.Println("VERSION 1.0")
			return nil
		},
	}
}

func createServeCommand() *cli.Command {
	return &cli.Command{
		Name:        "serve",
		Description: "starts an HTTP server",
		Action: func(ctx context.Context, c *cli.Command) error {
			db, err := database.Connect(ctx, config.DatabaseURL)
			if err != nil {
				log.Panic(err)
			}
			defer db.Close()

			app := fiber.New()

			go func() {
				err := app.Listen(":3000")
				if err != nil {
					log.Panic(err)
				}
			}()

			routes.MountRoutes(db, app)

			<-ctx.Done()

			shutdownErr := app.ShutdownWithTimeout(time.Second * 10)
			if shutdownErr != nil {
				log.Panic(shutdownErr)
			}

			return nil
		},
	}
}
