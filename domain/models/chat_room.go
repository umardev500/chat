package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/chat/constants"
)

type ChatRoom struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	Name      string     `json:"name,omitempty" db:"name"`
	IsGroup   bool       `json:"is_group" db:"is_group"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	Version   int        `json:"version" db:"version"`
}

type ChatRoomCreate struct {
	ID      uuid.UUID `json:"-" db:"id"`
	Name    string    `json:"name,omitempty" db:"name"`
	IsGroup bool      `json:"is_group" db:"is_group"`
}

type ChatRoomUpdate struct {
	ID      uuid.UUID `json:"-" db:"-"`
	Name    string    `json:"name,omitempty" db:"name"`
	IsGroup bool      `json:"is_group" db:"is_group"`
}

type ChatRoomDelete struct {
	ID         uuid.UUID            `json:"-" db:"id"`
	DeleteType constants.DeleteType `query:"delete_type" json:"delete_type" db:"delete_type" default:"soft"`
}

type ChatRoomFilter struct{}
type ChatRoomPageInfo struct {
	Offset int64 `query:"offset" json:"offset" default:"0"`
	Limit  int64 `query:"limit" json:"limit" default:"12"`
}

type ChatRoomFind struct {
	Filter   ChatRoomFilter
	PageInfo ChatRoomPageInfo
}
