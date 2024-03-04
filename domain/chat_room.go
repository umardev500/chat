package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/chat/domain/models"
)

type ChatRoomRepository interface {
	// Create creates a new chat room.
	//
	// Params:
	// 	- ctx: the context
	// 	- payload: the chat room to create
	Create(context.Context, models.ChatRoomCreate) error

	// Delete deletes a chat room.
	//
	// Params:
	// 	- ctx: the context
	// 	- payload: the chat room to delete
	Delete(context.Context, models.ChatRoomDelete) error

	// Find is a function that finds chat rooms from the repository.
	//
	// Params:
	// 	- ctx: the context
	// 	- payload: the chat room to find
	Find(context.Context, models.ChatRoomFind) ([]models.ChatRoom, error)

	// FindByID is a function that finds a chat room from the repository.
	//
	// Params:
	// 	- ctx: the context
	// 	- id: the id of the chat room to find
	FindByID(context.Context, uuid.UUID) (models.ChatRoom, error)

	// Update is a function that updates a chat room in the repository.
	//
	// Params:
	// 	- ctx: the context
	// 	- payload: the chat room to update
	Update(context.Context, models.ChatRoomUpdate) error
}
