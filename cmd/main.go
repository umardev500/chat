package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/chat/config"
	"github.com/umardev500/chat/database/migrations"
	"github.com/umardev500/chat/routes"
)

func init() {
	// log.Logger = log.With().Caller().Logger()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	if err := godotenv.Load(); err != nil {
		log.Fatal().Msgf("error loading .env file: %v", err)
	}
}

type Application struct {
	conn *config.Connection
}

func NewApplication(conn *config.Connection) *Application {
	return &Application{
		conn: conn,
	}
}

func (a *Application) Start(ctx context.Context) (err error) {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: false,
	})

	// Register router
	routes.NewRouter(app, a.conn).Register()

	ch := make(chan error, 1)
	go func() {
		port := os.Getenv("PORT")
		addr := fmt.Sprintf(":%s", port)
		log.Info().Msgf("Starting server on port %s", port)
		err = app.Listen(addr)
		if err != nil {
			log.Fatal().Msgf("Failed to start the server %v", err)
			ch <- err
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		log.Info().Msgf("Gracefully shutting down...")
		app.Shutdown()
	}

	return
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	db := config.NewPostgresConnection()
	conn := config.NewConnection(db)

	migrations.Migrate(conn)

	app := NewApplication(conn)
	err := app.Start(ctx)
	if err != nil {
		log.Error().Msgf("error starting application: %v", err)
	}
}
