package models

import (
	"time"
)

// RoomRestriction is used by pop to map your room_restrictions database table to your go code.
type RoomRestriction struct {
	ID            int       `json:"id" db:"id"`
	RoomID        int       `json:"room_id" db:"room_id"`
	ReservationId int       `json:"reservation_id" db:"reservation_id"`
	RestrictionID int       `json:"restriction_id" db:"restriction_id"`
	StartDate     time.Time `json:"start_date" db:"start_date"`
	EndDate       time.Time `json:"end_date" db:"end_date"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}
