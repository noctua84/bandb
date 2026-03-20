package models

import (
	"time"
)

// Restriction is used by pop to map your restrictions database table to your go code.
type Restriction struct {
	ID              int       `json:"id" db:"id"`
	RestrictionName string    `json:"restriction_name" db:"restriction_name"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}
