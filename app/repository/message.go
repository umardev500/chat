package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/umardev500/chat/config"
	"github.com/umardev500/chat/constants"
	"github.com/umardev500/chat/domain"
	"github.com/umardev500/chat/domain/models"
	"github.com/umardev500/chat/utils"
)

type messageRepository struct {
	conn *config.Connection
}

// NewMessageRepository returns a new MessageRepository
func NewMessageRepository(c *config.Connection) domain.MessageRepository {
	return &messageRepository{
		conn: c,
	}
}

// Create is a function that creates a new message.
//
// It takes the context and a payload of type models.MessageCreate as parameters and returns an error.
func (m *messageRepository) Create(ctx context.Context, payload models.MessageCreate) (err error) {
	query := `INSERT INTO messages (id, user_id, room_id, message) VALUES ($1, $2, $3, $4);`

	db := m.conn.TrOrDB(ctx)

	_, err = db.ExecContext(ctx, query, payload.ID, payload.UserID, payload.RoomID, payload.Message)
	return
}

// Delete deletes a message from the repository.
//
// It takes a context and a MessageDelete payload as parameters.
// It returns an error.
func (m *messageRepository) Delete(ctx context.Context, payload models.MessageDelete) (err error) {
	softDelete := payload.DeleteType == constants.SoftDelete
	var query string = fmt.Sprintf(`SELECT delete_record('messages', %s, %t)`, payload.ID, softDelete)

	db := m.conn.TrOrDB(ctx)

	result, err := db.ExecContext(ctx, query, payload.ID)
	if err != nil {
		return
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected < 1 {
		return constants.ErrNotAffected
	}

	return
}

// Find is a function that finds messages from the repository.
//
// It takes a context and a MessageFind payload as parameters.
// It returns a slice of Message and an error.
func (m *messageRepository) Find(ctx context.Context, payload models.MessageFind) (messages []models.Message, err error) {
	query := `
        SELECT *
        FROM messages
        LIMIT $1 OFFSET $2;
    `

	db := m.conn.TrOrDB(ctx)

	info := payload.PageInfo
	rows, err := db.QueryxContext(ctx, query, info.Limit, info.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var message models.Message
		if err := rows.StructScan(&message); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

// FindByID is a function that finds a message from the repository.
//
// It takes a context and a uuid.UUID as parameters.
// It returns a Message and an error.
func (m *messageRepository) FindByID(ctx context.Context, id uuid.UUID) (message models.Message, err error) {
	query := `SELECT * FROM messages WHERE id = $1;`

	db := m.conn.TrOrDB(ctx)

	err = db.QueryRowxContext(ctx, query, id).StructScan(&message)
	return
}

// Update is a function that updates a message in the repository.
//
// It takes a context and a MessageUpdate payload as parameters.
// It returns an error.
func (m *messageRepository) Update(ctx context.Context, payload models.MessageUpdate) (err error) {
	// Build a query string
	condition := "id = $1"
	query, args, err := utils.BuildUpdateQuery("messages", payload, condition, 1)
	if err != nil {
		return
	}

	db := m.conn.TrOrDB(ctx)

	// Append user_id and room_id as arguments
	args = append([]interface{}{payload.ID}, args...)

	// Execute query
	result, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return
	}

	// Get rows affected
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}

	// Check for affected rows if less than 1 it will return not affected error
	if affected < 1 {
		return constants.ErrNotAffected
	}

	return
}
