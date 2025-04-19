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

func GenerateNewMigration(db *sql.DB, name string) error {
	assertInDev()

	goose.SetDialect("sqlite")
	goose.SetBaseFS(migrationsFS)
	goose.SetSequential(true)

	return goose.Create(db, "./database/migrations", name, "sql")
}

func MigrationsStatus(ctx context.Context, db *sql.DB) error {
	return goose.StatusContext(ctx, db, "./database/migrations")
}

func MigrateToLatest(ctx context.Context, db *sql.DB) error {
	goose.SetDialect("sqlite")

	return goose.UpContext(ctx, db, "./database/migrations")
}

func assertInDev() {
	if config.AppEnv != config.EnvDev {
		log.Panic("application must be in env")
	}
}
