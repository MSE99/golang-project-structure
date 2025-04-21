package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/mse99/golang-project-structure/config"
	"github.com/mse99/golang-project-structure/database"
	"github.com/mse99/golang-project-structure/pkg/models"
	"github.com/urfave/cli/v3"
)

func createDbCommands() *cli.Command {
	return &cli.Command{
		Name:        "db",
		Description: "Commands to control the database",
		Commands: []*cli.Command{
			createDbStatusCommand(),
			createDbGenMigrationCommand(),
			createDbMigrateCommand(),
			createDbMigrateUpCommand(),
			createDbMigrateDownCommand(),
			createDbListUsersCommand(),
		},
	}
}

func createDbStatusCommand() *cli.Command {
	return &cli.Command{
		Name:        "status",
		Description: "checks the migration status against the current database",
		Action: func(ctx context.Context, c *cli.Command) error {
			conn, err := database.Connect(ctx, config.DatabaseURL)
			if err != nil {
				fmt.Println(err)
				return err
			}
			defer conn.Close()

			err = database.MigrationsStatus(ctx, conn)
			if err != nil {
				log.Panic(err)
			}

			fmt.Println("connected to the database successfully")

			return nil
		},
	}
}

func createDbGenMigrationCommand() *cli.Command {
	return &cli.Command{
		Name:        "gen:migration",
		Description: "generates a new migration",
		Action: func(ctx context.Context, c *cli.Command) error {
			name := c.Args().First()
			if name == "" {
				log.Fatal("name is required!")
			}

			db, err := database.Connect(ctx, config.DatabaseURL)
			if err != nil {
				log.Panic(err)
			}
			defer db.Close()

			err = database.GenerateNewMigration(db, name)
			if err != nil {
				log.Panic(err)
			}
			return nil
		},
	}
}

func createDbMigrateCommand() *cli.Command {
	return &cli.Command{
		Name:        "migrate",
		Description: "Migrates the database to the latest version",
		Action: func(ctx context.Context, c *cli.Command) error {
			db, err := database.Connect(ctx, config.DatabaseURL)
			if err != nil {
				log.Panic(err)
			}
			defer db.Close()

			err = database.MigrateToLatest(ctx, db)
			if err != nil {
				log.Panic(err)
			}
			return nil
		},
	}
}

func createDbMigrateUpCommand() *cli.Command {
	return &cli.Command{
		Name:        "migrate:up",
		Description: "Migrates the database to the latest version",
		Action: func(ctx context.Context, c *cli.Command) error {
			db, err := database.Connect(ctx, config.DatabaseURL)
			if err != nil {
				log.Panic(err)
			}
			defer db.Close()

			err = database.MigrateUp(ctx, db)
			if err != nil {
				log.Panic(err)
			}
			return nil
		},
	}
}

func createDbMigrateDownCommand() *cli.Command {
	return &cli.Command{
		Name:        "migrate:down",
		Description: "Migrates the database to the latest version",
		Action: func(ctx context.Context, c *cli.Command) error {
			db, err := database.Connect(ctx, config.DatabaseURL)
			if err != nil {
				log.Panic(err)
			}
			defer db.Close()

			err = database.MigrateDown(ctx, db)
			if err != nil {
				log.Panic(err)
			}
			return nil
		},
	}
}

func createDbListUsersCommand() *cli.Command {
	type user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	return &cli.Command{
		Name:        "list:users",
		Description: "lists the current users in the DB",
		Action: func(ctx context.Context, c *cli.Command) error {
			db, err := database.Connect(ctx, config.DatabaseURL)
			if err != nil {
				log.Panic(err)
			}
			defer db.Close()

			tx, err := db.BeginTx(ctx, &sql.TxOptions{})
			if err != nil {
				return nil
			}
			defer tx.Commit()

			users, err := models.LoadAllUsers(ctx, tx)
			if err != nil {
				return nil
			}

			log.Println("loaded all users")

			for _, u := range users {
				log.Println(u.Username)
			}

			return nil
		},
	}

}
