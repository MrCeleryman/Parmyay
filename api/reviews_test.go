package main

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestGetReviews(t *testing.T) {
	GetFunc(t, "/api/v1/reviews/", 200)
}
