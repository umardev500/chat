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

type userRoomRepository struct {
	conn *config.Connection
}

// NewUserRoomRepository returns a new UserRoomRepository
func NewUserRoomRepository(c *config.Connection) domain.UserRoomRepository {
	return &userRoomRepository{
		conn: c,
	}
}

func (u *userRoomRepository) Create(ctx context.Context, payload models.UserRoomCreate) (err error) {
	query := `INSERT INTO user_rooms (id, user_id, room_id) VALUES ($1, $2, $3);`

	db := u.conn.TrOrDB(ctx)

	_, err = db.ExecContext(ctx, query, payload.ID, payload.UserID, payload.RoomID)
	return
}

func (u *userRoomRepository) Delete(ctx context.Context, payload models.UserRoomDelete) (err error) {
	softDelete := payload.DeleteType == constants.SoftDelete
	var query string = fmt.Sprintf(`SELECT delete_record('user_rooms', %s, %t)`, payload.UserID, softDelete)

	db := u.conn.TrOrDB(ctx)

	result, err := db.ExecContext(ctx, query, payload.UserID)
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

func (u *userRoomRepository) Find(ctx context.Context, payload models.UserRoomFind) (userRooms []models.UserRoom, err error) {
	// Construct your query based on the payload
	// For simplicity, let's assume no filter is applied
	query := `
		SELECT *
		FROM user_rooms
		LIMIT $1 OFFSET $2;
	`

	db := u.conn.TrOrDB(ctx)

	info := payload.PageInfo
	rows, err := db.QueryxContext(ctx, query, info.Limit, info.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var userRoom models.UserRoom
		if err := rows.StructScan(&userRoom); err != nil {
			return nil, err
		}
		userRooms = append(userRooms, userRoom)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return userRooms, nil
}

func (u *userRoomRepository) FindByID(ctx context.Context, userID, roomID uuid.UUID) (userRoom models.UserRoom, err error) {
	query := `SELECT * FROM user_rooms WHERE user_id = $1 AND room_id = $2;`

	db := u.conn.TrOrDB(ctx)

	err = db.QueryRowxContext(ctx, query, userID, roomID).StructScan(&userRoom)
	return
}
