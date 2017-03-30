package main

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestGetVenues(t *testing.T) {
	GetFunc(t, "/api/v1/venues/", 200)
}
