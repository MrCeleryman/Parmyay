package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestGetVenues(t *testing.T) {
	os.Setenv("TEST", "1")
	db := InitDb()
	defer db.Close()
	testRouter := SetupRouter(true, false)
	request, err := http.NewRequest("GET", "/api/v1/venues/", nil)
	if err != nil {
		fmt.Println(err)
	}
	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)
	if response.Code != 200 {
		t.Errorf("Expected 200")
	}
}
