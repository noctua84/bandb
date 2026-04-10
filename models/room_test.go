package models

import (
	"testing"
	"time"
)

func TestRoom(t *testing.T) {
	now := time.Now()

	r := Room{
		ID:        1,
		RoomName:  "Major's Suite",
		CreatedAt: now,
		UpdatedAt: now,
	}

	if r.ID != 1 {
		t.Errorf("ID = %d, want %d", r.ID, 1)
	}

	if r.RoomName != "Major's Suite" {
		t.Errorf("RoomName = %q, want %q", r.RoomName, "Major's Suite")
	}
}
