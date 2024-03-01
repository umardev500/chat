package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/chat/domain/models"
)

type UserDetailRepository interface {
	// Create creates a new user detail
	//
	// Params:
	// 	- ctx: the context
	// 	- payload: the user detail to create
	Create(context.Context, models.UserDetailCreate) error

	// Delete deletes a user detail
	//
	// Params:
	// 	- ctx: the context
	// 	- payload: the user detail to delete
	Delete(context.Context, models.UserDelete) error

	// Find finds a user detail
	//
	// Params:
	// 	- ctx: the context
	// 	- payload: the user detail to find
	Find(context.Context, models.UserFind) ([]models.UserDetail, error)

	// FindByID finds a user detail
	//
	// Params:
	// 	- ctx: the context
	// 	- id: the id of the user detail to find
	FindByID(context.Context, uuid.UUID) (models.UserDetail, error)

	// Update updates a user detail
	//
	// Params:
	// 	- ctx: the context
	// 	- payload: the user detail to update
	Update(context.Context, models.UserDetailUpdate) error
}
