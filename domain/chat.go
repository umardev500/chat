package domain

import (
	"context"

	"github.com/umardev500/chat/domain/models"
)

type ChatUsecase interface {
	FindByUserID(context.Context, models.ChatFind) models.Response
}

type ChatRepository interface {
	FindByUserID(context.Context, models.ChatFind) ([]models.Chat, error)
}
