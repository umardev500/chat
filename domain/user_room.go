package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/chat/domain/models"
)

type UserRoomRepository interface {
	// Create creates a new user room.
	//
	// Params:
	// 	- ctx: the context
	// 	- payload: the user room to create
	Create(context.Context, models.UserRoomCreate) error

	// Delete deletes a user room.
	//
	// Params:
	// 	- ctx: the context
	// 	- payload: the user room to delete
	Delete(context.Context, models.UserRoomDelete) error

	// Find is a function that finds user rooms from the repository.
	//
	// Params:
	// 	- ctx: the context
	// 	- payload: the user room to find
	Find(context.Context, models.UserRoomFind) ([]models.UserRoom, error)

	// FindByID is a function that finds a user room from the repository.
	//
	// Params:
	// 	- ctx: the context
	// 	- userID: the user ID of the user room to find
	// 	- roomID: the room ID of the user room to find
	FindByID(context.Context, uuid.UUID, uuid.UUID) (models.UserRoom, error)
}
