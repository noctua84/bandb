package models

import (
	"time"
)

// Room is used by pop to map your rooms database table to your go code.
type Room struct {
	ID        int       `json:"id" db:"id"`
	RoomName  string    `json:"room_name" db:"room_name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
