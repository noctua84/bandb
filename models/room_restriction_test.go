package models

import (
	"testing"
	"time"
)

func TestRoomRestriction(t *testing.T) {
	now := time.Now()
	room := Room{ID: 1, RoomName: "General's Quarters"}
	reservation := Reservation{ID: 10, FirstName: "John"}
	restriction := Restriction{ID: 2, RestrictionName: "Owner Block"}

	rr := RoomRestriction{
		ID:            1,
		RoomID:        room.ID,
		Room:          room,
		ReservationID: reservation.ID,
		Reservation:   reservation,
		RestrictionID: restriction.ID,
		Restriction:   restriction,
		StartDate:     now,
		EndDate:       now.Add(48 * time.Hour),
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	if rr.Room.RoomName != "General's Quarters" {
		t.Errorf("Room.RoomName = %q, want %q", rr.Room.RoomName, "General's Quarters")
	}

	if rr.Reservation.FirstName != "John" {
		t.Errorf("Reservation.FirstName = %q, want %q", rr.Reservation.FirstName, "John")
	}

	if rr.Restriction.RestrictionName != "Owner Block" {
		t.Errorf("Restriction.RestrictionName = %q, want %q", rr.Restriction.RestrictionName, "Owner Block")
	}

	if rr.RoomID != room.ID {
		t.Errorf("RoomID = %d, want %d", rr.RoomID, room.ID)
	}

	if rr.ReservationID != reservation.ID {
		t.Errorf("ReservationID = %d, want %d", rr.ReservationID, reservation.ID)
	}

	if rr.RestrictionID != restriction.ID {
		t.Errorf("RestrictionID = %d, want %d", rr.RestrictionID, restriction.ID)
	}
}
