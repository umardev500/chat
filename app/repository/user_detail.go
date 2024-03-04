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

type userDetailRepository struct {
	conn *config.Connection
}

func NewUserDetailRepository(conn *config.Connection) domain.UserDetailRepository {
	return &userDetailRepository{
		conn: conn,
	}
}

// Hello
func (u *userDetailRepository) Create(ctx context.Context, payload models.UserDetailCreate) (err error) {
	query := `
		INSERT INTO user_details (id, first_name, last_name, date_of_birth, bio, photo)
		VALUES ($1, $2, $3, $4, $5, $6);
	`
	db := u.conn.TrOrDB(ctx)
	_, err = db.ExecContext(
		ctx,
		query,
		payload.ID,
		payload.FirstName,
		payload.LastName,
		payload.DateOfBirth,
		payload.Bio,
		payload.Photo,
	)

	return
}

func (u *userDetailRepository) Delete(ctx context.Context, payload models.UserDelete) (err error) {
	id := payload.ID
	softDelete := payload.DeleteType == constants.SoftDelete
	var query string = fmt.Sprintf(`SELECT delete_record('users', %s, %t)`, payload.ID, softDelete)

	db := u.conn.TrOrDB(ctx)
	result, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if affected < 1 {
		return constants.ErrNotAffected
	}

	return
}

func (u *userDetailRepository) Find(ctx context.Context, payload models.UserFind) (userDetails []models.UserDetail, err error) {
	query := `
		SELECT ud.* FROM user_details ud
	`

	db := u.conn.TrOrDB(ctx)
	cur, err := db.QueryxContext(ctx, query)
	if err != nil {
		return
	}

	userDetails = make([]models.UserDetail, 0)

	for cur.Next() {
		var each models.UserDetail
		if err = cur.StructScan(&each); err != nil {
			return
		}
		userDetails = append(userDetails, each)
	}

	return
}

func (u *userDetailRepository) FindByID(ctx context.Context, id uuid.UUID) (userDetail models.UserDetail, err error) {
	query := `SELECT ud.* FROM user_details ud WHERE id = $1;`

	db := u.conn.TrOrDB(ctx)
	err = db.QueryRowxContext(ctx, query, id).StructScan(&userDetail)

	return
}

func (u *userDetailRepository) Update(ctx context.Context, payload models.UserDetailUpdate) (err error) {
	condition := "id = $1"
	query, args, err := utils.BuildUpdateQuery("user_details", payload, condition, 1)
	if err != nil {
		return
	}

	args = append([]interface{}{payload.ID}, args...)
	db := u.conn.TrOrDB(ctx)

	result, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return
	}

	if affected < 1 {
		return constants.ErrNotAffected
	}

	return
}
