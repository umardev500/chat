package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/chat/domain/models"
)

type UserRepository interface {
	// Create creates a new user.
	//
	// Params:
	// 	- ctx: the context
	// 	- payload: the user to create
	Create(context.Context, models.UserCreate) error

	// Delete deletes a user.
	//
	// Params:
	// 	- ctx: the context
	// 	- payload: the user to delete
	Delete(context.Context, models.UserDelete) error

	// Find is a function that finds users from the repository.
	//
	// Params:
	// 	- ctx: the context
	// 	- payload: the user to find
	Find(context.Context, models.UserFind) ([]models.User, error)

	// FindByID is a function that finds a user from the repository.
	//
	// Params:
	// 	- ctx: the context
	// 	- id: the id of the user to find
	FindByID(context.Context, uuid.UUID) (models.User, error)

	// Update is a function that updates a user in the repository.
	//
	// Params:
	// 	- ctx: the context
	// 	- payload: the user to update
	Update(context.Context, models.UserUpdate) error
}
