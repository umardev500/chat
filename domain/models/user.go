package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/chat/constants"
)

type User struct {
	ID         uuid.UUID  `json:"id" db:"id"`
	Username   string     `json:"usernane" db:"username"`
	Email      string     `json:"email" db:"email"`
	Password   string     `json:"password" db:"password"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	DocVersion int64      `json:"doc_version" db:"doc_version"`
}

type UserCreate struct {
	ID       uuid.UUID `json:"-" db:"id"`
	Username string    `json:"username" db:"username"`
	Email    string    `json:"email" db:"email"`
	Password string    `json:"password" db:"password"`
}

type UserUpdate struct {
	ID       uuid.UUID `json:"-" db:"id"`
	Username string    `json:"username" db:"username"`
	Email    string    `json:"email" db:"email"`
	Password string    `json:"password" db:"password"`
}

type UserDelete struct {
	ID         uuid.UUID            `json:"-" db:"id"`
	DeleteType constants.DeleteType `query:"delete_type" json:"delete_type" db:"delete_type" default:"soft"`
}

type UserFilter struct{}
type UserPageInfo struct {
	Offset uint64 `query:"offset" json:"offset" default:"0"`
	Limit  uint64 `query:"limit" json:"limit" default:"12"`
}

type UserFind struct {
	Filter   UserFilter
	PageInfo UserPageInfo
}
