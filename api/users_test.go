package main

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestGetUsers(t *testing.T) {
	GetFunc(t, "/api/v1/users/", 200)
}
