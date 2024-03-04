package main

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/chat/config"
	"github.com/umardev500/chat/database/migrations"
	"github.com/umardev500/chat/database/seeders"
)

func init() {
	log.Logger = log.With().Caller().Logger()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	if err := godotenv.Load(); err != nil {
		log.Fatal().Msgf("error loading .env file: %v", err)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := config.NewPostgresConnection()
	conn := config.NewConnection(db)

	migrations.Down(conn)
	migrations.Migrate(conn)

	seeder := seeders.NewSeeder(conn)
	seeder.Populate(ctx)
}
