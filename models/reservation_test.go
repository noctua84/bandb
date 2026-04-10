package models

import (
	"testing"
	"time"
)

func TestReservation(t *testing.T) {
	now := time.Now()
	room := Room{ID: 1, RoomName: "General's Quarters"}

	r := Reservation{
		ID:        1,
		RoomID:    1,
		Room:      room,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Phone:     "555-1234",
		StartDate: now,
		EndDate:   now.Add(24 * time.Hour),
		CreatedAt: now,
		UpdatedAt: now,
	}

	if r.RoomID != room.ID {
		t.Errorf("RoomID = %d, want %d", r.RoomID, room.ID)
	}

	if r.Room.RoomName != "General's Quarters" {
		t.Errorf("Room.RoomName = %q, want %q", r.Room.RoomName, "General's Quarters")
	}

	if r.FirstName != "John" {
		t.Errorf("FirstName = %q, want %q", r.FirstName, "John")
	}

	if r.Email != "john@example.com" {
		t.Errorf("Email = %q, want %q", r.Email, "john@example.com")
	}
}
