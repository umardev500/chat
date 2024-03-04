package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/chat/domain/models"
)

type MessageRepository interface {
	// Create creates a new message.
	//
	// Params:
	// 	- ctx: the context
	// 	- payload: the message to create
	Create(context.Context, models.MessageCreate) error

	// Delete deletes a message.
	//
	// Params:
	// 	- ctx: the context
	// 	- payload: the message to delete
	Delete(context.Context, models.MessageDelete) error

	// Find is a function that finds messages from the repository.
	//
	// Params:
	// 	- ctx: the context
	// 	- payload: the message to find
	Find(context.Context, models.MessageFind) ([]models.Message, error)

	// FindByID is a function that finds a message from the repository.
	//
	// Params:
	// 	- ctx: the context
	// 	- id: the id of the message to find
	FindByID(context.Context, uuid.UUID) (models.Message, error)

	// Update is a function that updates a message in the repository.
	//
	// Params:
	// 	- ctx: the context
	// 	- payload: the message to update
	Update(context.Context, models.MessageUpdate) error
}
