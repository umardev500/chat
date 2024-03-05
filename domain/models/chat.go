package models

import (
	"time"

	"github.com/google/uuid"
)

type Chat struct {
	ID             uuid.UUID  `json:"id" db:"id"`
	Name           string     `json:"name" db:"name"`
	HiglightedText *string    `json:"highlighted_text,omitempty" db:"highlighted_text"`
	IsMe           bool       `json:"is_me" db:"is_me"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type ChatFilter struct {
	StartTime *time.Time `query:"start_time" json:"start_time"`
	UserID    *uuid.UUID
}

type ChatPageInfo struct {
	Offset int64 `query:"offset" json:"offset" default:"0"`
	Limit  int64 `query:"limit" json:"limit" default:"12"`
}

type ChatFind struct {
	Filter   ChatFilter   `json:"filter"`
	PageInfo ChatPageInfo `json:"page_info"`
}
