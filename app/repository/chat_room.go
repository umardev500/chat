package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/umardev500/chat/config"
	"github.com/umardev500/chat/constants"
	"github.com/umardev500/chat/domain"
	"github.com/umardev500/chat/domain/models"
)

type chatRoomRepository struct {
	conn *config.Connection
}

// NewChatRoomRepository returns a new ChatRoomRepository
func NewChatRoomRepository(c *config.Connection) domain.ChatRoomRepository {
	return &chatRoomRepository{
		conn: c,
	}
}

// Create is a function that creates a new chat room.
//
// It takes the context and a payload of type models.ChatRoomCreate as parameters and returns an error.
func (c *chatRoomRepository) Create(ctx context.Context, payload models.ChatRoomCreate) (err error) {
	query := `INSERT INTO chat_rooms (id, name, is_group) VALUES ($1, $2, $3);`

	db := c.conn.TrOrDB(ctx)

	_, err = db.ExecContext(ctx, query, payload.ID, payload.Name, payload.IsGroup)
	return
}

// Delete deletes a chat room from the repository.
//
// It takes a context and a ChatRoomDelete payload as parameters.
// It returns an error.
func (c *chatRoomRepository) Delete(ctx context.Context, payload models.ChatRoomDelete) (err error) {
	softDelete := payload.DeleteType == constants.SoftDelete
	var query string = fmt.Sprintf(`SELECT delete_record('chat_rooms', %s, %t)`, payload.ID, softDelete)

	db := c.conn.TrOrDB(ctx)

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

// Find is a function that finds chat rooms from the repository.
//
// It takes a context and a ChatRoomFind payload as parameters.
// It returns a slice of ChatRoom and an error.
func (c *chatRoomRepository) Find(ctx context.Context, payload models.ChatRoomFind) (chatRooms []models.ChatRoom, err error) {
	// Construct your query based on the payload
	// For simplicity, let's assume no filter is applied
	query := `
		SELECT *
		FROM chat_rooms
		LIMIT $1 OFFSET $2;
	`

	db := c.conn.TrOrDB(ctx)

	info := payload.PageInfo
	rows, err := db.QueryxContext(ctx, query, info.Limit, info.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var chatRoom models.ChatRoom
		if err := rows.StructScan(&chatRoom); err != nil {
			return nil, err
		}
		chatRooms = append(chatRooms, chatRoom)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return chatRooms, nil
}

// FindByID is a function that finds a chat room from the repository.
//
// It takes a context and a uuid.UUID as parameters.
// It returns a ChatRoom and an error.
func (c *chatRoomRepository) FindByID(ctx context.Context, id uuid.UUID) (chatRoom models.ChatRoom, err error) {
	query := `SELECT * FROM chat_rooms WHERE id = $1;`

	db := c.conn.TrOrDB(ctx)

	err = db.QueryRowxContext(ctx, query, id).StructScan(&chatRoom)
	return
}

// Update is a function that updates a chat room in the repository.
//
// It takes a context and a ChatRoomUpdate payload as parameters.
// It returns an error.
func (c *chatRoomRepository) Update(ctx context.Context, payload models.ChatRoomUpdate) (err error) {
	// Build your update query based on the payload
	query := `
		UPDATE chat_rooms
		SET name = $2, is_group = $3
		WHERE id = $1;
	`

	db := c.conn.TrOrDB(ctx)

	_, err = db.ExecContext(ctx, query, payload.ID, payload.Name, payload.IsGroup)
	return
}
