package migrations

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/chat/config"
)

func Down(conn *config.Connection) {
	driver, err := postgres.WithInstance(conn.DB, &postgres.Config{})
	if err != nil {
		log.Fatal().Msgf("failed to create postgres driver: %v", err)
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"postgres",
		driver,
	)

	if err != nil {
		log.Fatal().Msgf("failed to create migrate instance: %v", err)
		return
	}

	log.Info().Msg("migrating down...")
	err = m.Down()
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Info().Msg(err.Error())
		} else {
			log.Fatal().Msgf("failed to migrate: %v", err)
		}
	}

	log.Info().Msg("migrations down completed")
}

func Migrate(conn *config.Connection) {
	driver, err := postgres.WithInstance(conn.DB, &postgres.Config{})
	if err != nil {
		log.Fatal().Msgf("failed to create postgres driver: %v", err)
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"postgres",
		driver,
	)

	if err != nil {
		log.Fatal().Msgf("failed to create migrate instance: %v", err)
		return
	}

	log.Info().Msg("migrating...")
	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Info().Msg(err.Error())
		} else {

			log.Fatal().Msgf("failed to migrate: %v", err)
		}
	}
	log.Info().Msg("migrations completed")
}
