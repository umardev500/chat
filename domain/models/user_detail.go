package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/chat/constants"
)

type UserDetail struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	FirstName   string     `json:"first_name" db:"first_name"`
	LastName    *string    `json:"last_name,omitempty" db:"last_name"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty" db:"date_of_birth"`
	Bio         *string    `json:"bio,omitempty" db:"bio"`
	Photo       *string    `json:"photo,omitempty" db:"photo"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	DocVersion  int64      `json:"doc_version" db:"doc_version"`
}

type UserDetailCreate struct {
	ID          uuid.UUID  `json:"-" db:"id"`
	FirstName   string     `json:"first_name" db:"first_name"`
	LastName    *string    `json:"last_name,omitempty" db:"last_name"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty" db:"date_of_birth"`
	Bio         *string    `json:"bio,omitempty" db:"bio"`
	Photo       *string    `json:"photo,omitempty" db:"photo"`
}

type UserDetailUpdate struct {
	ID          uuid.UUID  `json:"-" db:"-"`
	FirstName   string     `json:"first_name,omitempty" db:"first_name"`
	LastName    *string    `json:"last_name,omitempty" db:"last_name"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty" db:"date_of_birth"`
	Bio         *string    `json:"bio,omitempty" db:"bio"`
	Photo       *string    `json:"photo,omitempty" db:"photo"`
}

type UserDetailDelete struct {
	ID         uuid.UUID            `json:"-" db:"id"`
	DeleteType constants.DeleteType `query:"delete_type" json:"delete_type" db:"delete_type" default:"soft"`
}

type UserDetailFilter struct{}
type UserDetailPageInfo struct {
	Offset int64 `query:"offset" json:"offset" default:"0"`
	Limit  int64 `query:"limit" json:"limit" default:"12"`
}

type UserDetailFind struct {
	Filter   UserDetailFilter   `json:"filter"`
	PageInfo UserDetailPageInfo `json:"page_info"`
}
