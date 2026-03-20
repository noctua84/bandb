package models

import (
	"time"
)

// RoomRestriction is used by pop to map your room_restrictions database table to your go code.
type RoomRestriction struct {
	ID        int       `json:"id" db:"id"`
	StartDate time.Time `json:"start_date" db:"start_date"`
	EndDate   time.Time `json:"end_date" db:"end_date"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
