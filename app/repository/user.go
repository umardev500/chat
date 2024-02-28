package repository

import (
	"github.com/umardev500/chat/config"
	"github.com/umardev500/chat/domain"
)

type userRepository struct {
	conn *config.Connection
}

func NewUserRepository(c *config.Connection) domain.UserRepository {
	return &userRepository{
		conn: c,
	}
}
