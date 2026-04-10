package models

import (
	"testing"
	"time"
)

func TestRestriction(t *testing.T) {
	now := time.Now()

	r := Restriction{
		ID:              1,
		RestrictionName: "Owner Block",
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	if r.ID != 1 {
		t.Errorf("ID = %d, want %d", r.ID, 1)
	}

	if r.RestrictionName != "Owner Block" {
		t.Errorf("RestrictionName = %q, want %q", r.RestrictionName, "Owner Block")
	}
}
