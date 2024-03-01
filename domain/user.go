package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/chat/domain/models"
)

type UserRepository interface {
	Create(context.Context, models.UserCreate) error
	Delete(context.Context, models.UserDelete) error
	Find(context.Context, models.UserFind) ([]models.User, error)
	FindByID(context.Context, uuid.UUID) (models.User, error)
	Update(context.Context, models.UserUpdate) error
}
