package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func NewPostgresConnection() (db *sqlx.DB) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	start := time.Now()
	db, err := sqlx.ConnectContext(ctx, "postgres", dsn)
	if err != nil {
		log.Fatal().Msgf("error connecting to postgres: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal().Msgf("error pinging postgres: %v", err)
	}

	duration := time.Since(start)
	msg := fmt.Sprintf("Connected to Postgres \033[32m🎉 (\U000023F3 %s)\033[0m", duration)
	log.Info().Msg(msg)

	return
}
