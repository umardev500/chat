package domain

import (
	"context"

	"github.com/umardev500/chat/domain/models"
)

type ChatRepository interface {
	Find(context.Context, models.ChatFind) ([]models.Chat, error)
}
