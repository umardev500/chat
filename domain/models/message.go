package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/chat/constants"
)

type Message struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	UserID    uuid.UUID  `json:"user_id" db:"user_id"`
	RoomID    uuid.UUID  `json:"room_id" db:"room_id"`
	Message   string     `json:"message" db:"message"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type MessageCreate struct {
	ID      uuid.UUID `json:"-" db:"id"`
	UserID  uuid.UUID `json:"user_id" db:"user_id"`
	RoomID  uuid.UUID `json:"room_id" db:"room_id"`
	Message string    `json:"message" db:"message"`
}

type MessageUpdate struct {
	ID      uuid.UUID `json:"-" db:"-"`
	UserID  uuid.UUID `json:"user_id,omitempty" db:"user_id"`
	RoomID  uuid.UUID `json:"room_id,omitempty" db:"room_id"`
	Message string    `json:"message,omitempty" db:"message"`
}

type MessageDelete struct {
	ID         uuid.UUID            `json:"-" db:"id"`
	DeleteType constants.DeleteType `query:"delete_type" json:"delete_type" db:"delete_type" default:"soft"`
}

type MessageFilter struct{}
type MessagePageInfo struct {
	Offset int64 `query:"offset" json:"offset" default:"0"`
	Limit  int64 `query:"limit" json:"limit" default:"12"`
}

type MessageFind struct {
	Filter   MessageFilter
	PageInfo MessagePageInfo
}
