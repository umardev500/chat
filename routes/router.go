package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/chat/config"
	"github.com/umardev500/chat/injection"
)

type Router struct {
	app  *fiber.App
	conn *config.Connection
}

func NewRouter(app *fiber.App, conn *config.Connection) *Router {
	return &Router{
		app:  app,
		conn: conn,
	}
}

func (r *Router) Register() {
	api := r.app.Group("/api")
	injection.ChatInjection(api, r.conn)
}
