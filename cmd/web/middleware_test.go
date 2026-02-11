package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var nh myHandler
	wrappedHandler := NoSurf(&nh)

	if wrappedHandler == nil {
		t.Error("NoSurf returned nil handler")
	}

	switch h := wrappedHandler.(type) {
	case http.Handler:
		// Test passed
	default:
		t.Errorf("NoSurf returned type %T, expected http.Handler", h)
	}
}

func TestSessionLoad(t *testing.T) {
	var nh myHandler
	wrappedHandler := SessionLoad(&nh)

	if wrappedHandler == nil {
		t.Error("SessionLoad returned nil handler")
	}

	switch h := wrappedHandler.(type) {
	case http.Handler:
		// Test passed
	default:
		t.Errorf("SessionLoad returned type %T, expected http.Handler", h)
	}
}
