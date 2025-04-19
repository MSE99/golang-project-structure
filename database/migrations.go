package database

import (
	"context"
	"database/sql"
	"embed"
	"log"

	"github.com/mse99/golang-project-structure/config"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

func setupGoose() {
	goose.SetDialect("sqlite")
	goose.SetBaseFS(migrationsFS)
	goose.SetSequential(true)
}

func GenerateNewMigration(db *sql.DB, name string) error {
	assertInDev()

	setupGoose()

	return goose.Create(db, "migrations", name, "sql")
}

func MigrationsStatus(ctx context.Context, db *sql.DB) error {
	setupGoose()
	return goose.StatusContext(ctx, db, "migrations")
}

func MigrateToLatest(ctx context.Context, db *sql.DB) error {
	setupGoose()
	return goose.UpContext(ctx, db, "migrations")
}

func MigrateUp(ctx context.Context, db *sql.DB) error {
	setupGoose()
	return goose.UpByOneContext(ctx, db, "migrations")
}

func MigrateDown(ctx context.Context, db *sql.DB) error {
	setupGoose()
	return goose.DownContext(ctx, db, "migrations")
}

func assertInDev() {
	if config.AppEnv != config.EnvDev {
		log.Panic("application must be in env")
	}
}
