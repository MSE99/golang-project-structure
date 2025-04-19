package main

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func createMainCliCommand() *cli.Command {
	return &cli.Command{
		Name:        "example",
		Description: "This is an example golang project",
		Commands: []*cli.Command{
			createVersionCommand(),
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
