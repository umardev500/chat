package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/chat/constants"
)

type UserRoom struct {
	UserID    uuid.UUID  `json:"user_id" db:"user_id"`
	RoomID    uuid.UUID  `json:"room_id" db:"room_id"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type UserRoomCreate struct {
	UserID uuid.UUID `json:"user_id" db:"user_id"`
	RoomID uuid.UUID `json:"room_id" db:"room_id"`
}

type UserRoomDelete struct {
	UserID     uuid.UUID            `json:"user_id" db:"user_id"`
	RoomID     uuid.UUID            `json:"room_id" db:"room_id"`
	DeleteType constants.DeleteType `query:"delete_type" json:"delete_type" db:"delete_type" default:"soft"`
}

type UserRoomFind struct {
	Filter   UserRoomFilter
	PageInfo UserRoomPageInfo
}

type UserRoomFilter struct{}

type UserRoomPageInfo struct {
	Offset int64 `query:"offset" json:"offset" default:"0"`
	Limit  int64 `query:"limit" json:"limit" default:"12"`
}
