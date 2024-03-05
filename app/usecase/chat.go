package usecase

import (
	"context"
	"fmt"

	"github.com/umardev500/chat/domain"
	"github.com/umardev500/chat/domain/models"
)

type chatUsecase struct {
	repo domain.ChatRepository
}

func NewChatUsecase(repo domain.ChatRepository) domain.ChatUsecase {
	return &chatUsecase{
		repo: repo,
	}
}

func (c *chatUsecase) FindByUserID(ctx context.Context, find models.ChatFind) {
	chats, err := c.repo.FindByUserID(ctx, find)
	if err != nil {
		return
	}

	fmt.Println(chats)
}
