package db

import (
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func RunMigrations(client *sql.DB) {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Fatal().Err(err).Msg("failed to set dialect")
	}

	if err := goose.Up(client, "migrations"); err != nil {
		log.Fatal().Err(err).Msg("failed to run migrations")
	}
}
