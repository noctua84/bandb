package models

import (
	"time"
)

// Reservation is used by pop to map your reservations database table to your go code.
type Reservation struct {
	ID        int       `json:"id" db:"id"`
	RoomID    int       `json:"room_id" db:"room_id"`
	Room      Room      `json:"room"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	Phone     string    `json:"phone" db:"phone"`
	StartDate time.Time `json:"start_date" db:"start_date"`
	EndDate   time.Time `json:"end_date" db:"end_date"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
