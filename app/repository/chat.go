package repository

import (
	"context"

	"github.com/umardev500/chat/config"
	"github.com/umardev500/chat/domain"
	"github.com/umardev500/chat/domain/models"
)

type chatRepository struct {
	conn *config.Connection
}

func NewChatRepository(conn *config.Connection) domain.ChatRepository {
	return &chatRepository{
		conn: conn,
	}
}

func (c *chatRepository) FindByUserID(ctx context.Context, find models.ChatFind) (chats []models.Chat, err error) {
	filter := find.Filter
	userID := filter.UserID

	query := `
	SELECT 
		cr.id, 
		CASE
			WHEN cr.is_group = true THEN cr.name
			ELSE u_other.username
		END AS name,
		cr.updated_at,
		m.message AS highlighted_text,
		CASE
			WHEN m.user_id = ur.user_id THEN true
			ELSE false
		END AS is_me, -- the last message was sent by me
		m.created_at AS created_at
	FROM 
		chat_rooms cr
	JOIN 
		user_rooms ur ON cr.id = ur.room_id
	LEFT JOIN 
		user_rooms ur_other ON cr.id = ur_other.room_id AND ur.user_id != ur_other.user_id
	LEFT JOIN 
		users u_other ON ur_other.user_id = u_other.id
	LEFT JOIN LATERAL (
		SELECT 
			user_id,
			"message", 
			created_at
		FROM 
			messages 
		WHERE 
			room_id = cr.id  -- Ensure the message is for the current chat room
		ORDER BY 
			created_at DESC 
		LIMIT 1
	) m ON true
	WHERE 
		ur.user_id = $1
	AND cr.created_at > '2024-01-05T14:17:14+07:00';
	`

	chats = make([]models.Chat, 0)
	db := c.conn.TrOrDB(ctx)
	cur, err := db.QueryxContext(ctx, query, userID)
	if err != nil {
		return
	}

	for cur.Next() {
		var each models.Chat
		if err = cur.StructScan(&each); err != nil {
			return
		}
		chats = append(chats, each)
	}

	return
}
