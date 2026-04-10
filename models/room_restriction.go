package models

import (
	"time"
)

// RoomRestriction is used by pop to map your room_restrictions database table to your go code.
type RoomRestriction struct {
	ID            int       `json:"id" db:"id"`
	RoomID        int         `json:"room_id" db:"room_id"`
	Room          Room        `json:"room"`
	ReservationID int         `json:"reservation_id" db:"reservation_id"`
	Reservation   Reservation `json:"reservation"`
	RestrictionID int         `json:"restriction_id" db:"restriction_id"`
	Restriction   Restriction `json:"restriction"`
	StartDate     time.Time `json:"start_date" db:"start_date"`
	EndDate       time.Time `json:"end_date" db:"end_date"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}
