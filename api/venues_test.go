package main

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestGetVenues(t *testing.T) {
	GetFunc(t, "/api/v1/venues/", 200)
}

func TestGetVenueIncorrectId(t *testing.T) {
	GetFunc(t, "/api/v1/venues/0", 404)
}

func TestGetVenueCorrectId(t *testing.T) {
	GetFunc(t, "/api/v1/venues/1", 200)
}

func TestDeleteVenueCorrectId(t *testing.T) {
	SoftDeleteFunc(t, "/api/v1/venues/2", 200)
}

func TestDeleteVenueIncorrectId(t *testing.T) {
	SoftDeleteFunc(t, "/api/v1/venues/0", 404)
}
