package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// User is used by pop to map your users database table to your go code.
type User struct {
	ID          int       `json:"id" db:"id"`
	ExternalID  uuid.UUID `json:"external_id" db:"external_id"`
	FirstName   string    `json:"first_name" db:"first_name" validate:"required"`
	LastName    string    `json:"last_name" db:"last_name" validate:"required"`
	Email       string    `json:"email" db:"email" validate:"required,email"`
	Password    string    `json:"password" db:"password" validate:"required"`
	Phone       string    `json:"phone" db:"phone"`
	AccessLevel int       `json:"access_level" db:"access_level"`
	JoinedAt    time.Time `json:"joined_at" db:"joined_at"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
