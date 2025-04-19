package main

import (
	"context"
	"fmt"

	"github.com/mse99/golang-project-structure/config"
	"github.com/mse99/golang-project-structure/database"
	"github.com/urfave/cli/v3"
)

func createMainCliCommand() *cli.Command {
	return &cli.Command{
		Name:        "example",
		Description: "This is an example golang project",
		Commands: []*cli.Command{
			createVersionCommand(),
			createDbCommands(),
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

func createDbCommands() *cli.Command {
	return &cli.Command{
		Name:        "db",
		Description: "Commands to control the database",
		Commands: []*cli.Command{
			{
				Name:        "test",
				Description: "Tests the connection to the database",
				Action: func(ctx context.Context, c *cli.Command) error {
					_, err := database.Connect(ctx, config.DatabaseURL)
					if err != nil {
						fmt.Println(err)
						return err
					}

					fmt.Println("connected to the database successfully")

					return nil
				},
			},
		},
	}
}
