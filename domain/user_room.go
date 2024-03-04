package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/chat/domain/models"
)

type UserRoomRepository interface {
	// Create adds a user to a room.
	//
	// Params:
	// 	- ctx: the context
	// 	- payload: the user room association to create
	Create(context.Context, models.UserRoomCreate) error

	// Delete removes a user from a room.
	//
	// Params:
	// 	- ctx: the context
	// 	- payload: the user room association to delete
	Delete(context.Context, models.UserRoomDelete) error

	// FindRoomsByUserID returns the list of rooms associated with a user.
	//
	// Params:
	// 	- ctx: the context
	// 	- userID: the ID of the user
	FindRoomsByUserID(context.Context, uuid.UUID) ([]uuid.UUID, error)

	// FindUsersByRoomID returns the list of users associated with a room.
	//
	// Params:
	// 	- ctx: the context
	// 	- roomID: the ID of the room
	FindUsersByRoomID(context.Context, uuid.UUID) ([]uuid.UUID, error)
}
