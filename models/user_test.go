package models

import (
	"testing"
	"time"

	"github.com/gofrs/uuid"
)

func TestUser(t *testing.T) {
	now := time.Now()
	uid, _ := uuid.NewV4()

	u := User{
		ID:          1,
		ExternalID:  uid,
		FirstName:   "Jane",
		LastName:    "Doe",
		Email:       "jane@example.com",
		Password:    "hashed-password",
		Phone:       "555-5678",
		AccessLevel: 1,
		JoinedAt:    now,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if u.ExternalID != uid {
		t.Errorf("ExternalID = %v, want %v", u.ExternalID, uid)
	}

	if u.AccessLevel != 1 {
		t.Errorf("AccessLevel = %d, want %d", u.AccessLevel, 1)
	}

	if u.Email != "jane@example.com" {
		t.Errorf("Email = %q, want %q", u.Email, "jane@example.com")
	}
}
