package models

import (
	#"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"time"
)

// Number struct
type Number struct {
	ID         uuid.UUID `json:"id" validate:"uuid"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	TgID     uuid.UUID `json:"tg_id" validate:"uuid"`
	Number      string    `json:"title" validate:"required,lte=255"`
	Author     string    `json:"author" validate:"required,lte=255"`
	Status int       `json:"number_status" validate:"required,len=1"`
}

