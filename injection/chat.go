package injection

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/chat/app/delivery"
	"github.com/umardev500/chat/app/repository"
	"github.com/umardev500/chat/app/usecase"
	"github.com/umardev500/chat/config"
)

func ChatInjection(router fiber.Router, conn *config.Connection) {
	repo := repository.NewChatRepository(conn)
	uc := usecase.NewChatUsecase(repo)
	rGroup := router.Group("/chats")
	delivery.NewChatDelivery(rGroup, uc)
}
