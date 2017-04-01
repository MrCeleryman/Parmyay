package main

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestGetReviews(t *testing.T) {
	GetFunc(t, "/api/v1/reviews/", 200)
}

func TestGetReviewIncorrectId(t *testing.T) {
	GetFunc(t, "/api/v1/reviews/0", 404)
}

func TestGetReviewCorrectId(t *testing.T) {
	GetFunc(t, "/api/v1/reviews/1", 200)
}

func TestDeleteReviewCorrectId(t *testing.T) {
	SoftDeleteFunc(t, "/api/v1/reviews/2", 200)
}

func TestDeleteReviewIncorrectId(t *testing.T) {
	SoftDeleteFunc(t, "/api/v1/reviews/0", 404)
}
