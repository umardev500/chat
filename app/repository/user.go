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

type userRepository struct {
	conn *config.Connection
}

// NewUserRepository returns a new UserRepository
func NewUserRepository(c *config.Connection) domain.UserRepository {
	return &userRepository{
		conn: c,
	}
}

// Create is a function that creates a new user.
//
// It takes the context and a payload of type models.UserCreate as parameters and returns an error.
func (u *userRepository) Create(ctx context.Context, payload models.UserCreate) (err error) {
	query := `INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4);`

	db := u.conn.TrOrDB(ctx)

	_, err = db.ExecContext(ctx, query, payload.ID, payload.Username, payload.Email, payload.Password)
	return
}

// Delete deletes a user from the repository.
//
// It takes a context and a UserDelete payload as parameters.
// It returns an error.
func (u *userRepository) Delete(ctx context.Context, payload models.UserDelete) (err error) {
	softDelete := payload.DeleteType == constants.SoftDelete
	var query string = fmt.Sprintf(`SELECT delete_record('users', %s, %t)`, payload.ID, softDelete)

	db := u.conn.TrOrDB(ctx)

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

// Find is a function that finds users from the repository.
//
// It takes a context and a UserFind payload as parameters.
// It returns a slice of User and an error.
func (u *userRepository) Find(ctx context.Context, payload models.UserFind) (users []models.User, err error) {
	whereClause := "1=1"

	query := fmt.Sprintf(`
		SELECT u.*
		FROM users u
		WHERE %s
		LIMIT $1
		OFFSET $2
	`, whereClause)

	users = make([]models.User, 0)

	db := u.conn.TrOrDB(ctx)
	info := payload.PageInfo
	cur, err := db.QueryxContext(ctx, query, info.Limit, info.Offset)
	if err != nil {
		return
	}

	for cur.Next() {
		var each models.User
		if err = cur.StructScan(&each); err != nil {
			return
		}
		users = append(users, each)
	}

	return
}

// FindByID is a function that finds a user from the repository.
//
// It takes a context and a uuid.UUID as parameters.
// It returns a User and an error.
func (u *userRepository) FindByID(ctx context.Context, id uuid.UUID) (user models.User, err error) {
	query := `SELECT u.* FROM users u WHERE id = $1;`

	db := u.conn.TrOrDB(ctx)

	err = db.QueryRowxContext(ctx, query, id).StructScan(&user)
	return
}

// Update is a function that updates a user in the repository.
//
// It takes a context and a UserUpdate payload as parameters.
// It returns an error.
func (u *userRepository) Update(ctx context.Context, payload models.UserUpdate) (err error) {
	// Build a query string
	condition := "id = $1"
	query, args, err := utils.BuildUpdateQuery("users", payload, condition, 1)
	if err != nil {
		return
	}

	db := u.conn.TrOrDB(ctx)

	// Append id as first argument
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
