package delivery

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/chat/domain"
	"github.com/umardev500/chat/domain/models"
)

type chatDelivery struct {
	uc domain.ChatUsecase
}

func NewChatDelivery(router fiber.Router, uc domain.ChatUsecase) {
	handler := &chatDelivery{
		uc: uc,
	}

	router.Get("/", handler.GetChat)
}

func (cd *chatDelivery) GetChat(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	userID, err := uuid.Parse("223e4567-e89b-12d3-a456-426614174001")
	if err != nil {
		return c.SendStatus(500)
	}

	filter := models.ChatFilter{}
	filter.UserID = &userID
	find := models.ChatFind{
		Filter: filter,
	}
	cd.uc.FindByUserID(ctx, find)

	return nil
}
