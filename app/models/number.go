package models

import (
	"github.com/google/uuid"
)

// Number struct
type Number struct {
	id            uuid.UUID `json:"id" validate:"uuid"`
	phone_number  string    `json:"Ph_num" validate:"required,lte=15"`
	tg_id         string    `json:"tg_id" validate:"uuid"`
	fio           string    `json:"fio" validate:"required,lte=255"`
	number_status int       `json:"number_status" validate:"required,len=1"`
}
