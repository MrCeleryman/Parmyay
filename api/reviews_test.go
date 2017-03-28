package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestGetReviews(t *testing.T) {
	testRouter := SetupRouter(false, false)
	req, err := http.NewRequest("GET", "/api/v1/reviews/", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	if resp.Code != 200 {
		t.Errorf("Expected 200")
	}
}
