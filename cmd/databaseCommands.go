package main

import (
	"context"
	"fmt"

	"github.com/mse99/golang-project-structure/config"
	"github.com/mse99/golang-project-structure/database"
	"github.com/urfave/cli/v3"
)

func createDbCommands() *cli.Command {
	return &cli.Command{
		Name:        "db",
		Description: "Commands to control the database",
		Commands: []*cli.Command{
			createDbTestCommand(),
		},
	}
}

func createDbTestCommand() *cli.Command {
	return &cli.Command{
		Name:        "test",
		Description: "Tests the connection to the database",
		Action: func(ctx context.Context, c *cli.Command) error {
			conn, err := database.Connect(ctx, config.DatabaseURL)
			if err != nil {
				fmt.Println(err)
				return err
			}
			defer conn.Close()

			fmt.Println("connected to the database successfully")

			return nil
		},
	}
}
