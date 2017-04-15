package parmyay

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("DB_NAME", ":memory:")
	InitDb()
	DB.LogMode(false)

	DB.Create(Venue{ID: 1, Address: "30 Willy Wonka Way", VenueName: "Dans House"})
	DB.Create(Venue{ID: 2, Address: "42 Wallaby Way, Sydney", VenueName: "Dans Old House"})

	DB.Create(Achievement{ID: 1, Achievement: "Reviewed first Parmy!"})
	DB.Create(Achievement{ID: 2, Achievement: "Ate first Parmy!"})

	DB.Create(Review{ID: 1, Notes: "It was prety good", UserID: 1, VenueID: 1, Rating: 10})
	DB.Create(Review{ID: 2, Notes: "Too much water", UserID: 1, VenueID: 1, Rating: 7.8})

	DB.Create(User{ID: 1, FirstName: "Daniel", LastName: "Mitchell", UserName: "DMitch", Email: "d@gmail.com", Password: []byte("Daniel")})
	DB.Create(User{ID: 2, FirstName: "Jerry", LastName: "Seinfeld", UserName: "DarkLordDD", Email: "bee@movie.com", Password: []byte("dfdfgdfgdfg")})

	os.Exit(m.Run())
}

// Test helper methods
// PostFunc is a handler function which sends a POST request to the local API
func PostFunc(t *testing.T, json string, url string, expectedCode int) {
	testRouter := SetupRouter(true, false)

	var jsonByte = []byte(json)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonByte))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
	}
	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)
	if response.Code != expectedCode {
		t.Errorf("Expected %d", expectedCode)
	}
}

// GetFunc is a handler function which sends a GET request to the local API
func GetFunc(t *testing.T, url string, expectedCode int) *httptest.ResponseRecorder {
	testRouter := SetupRouter(true, false)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)
	return response
}

// DeleteFunc is a handler function which sends a DELETE request to the local API
func DeleteFunc(t *testing.T, url string, expectedCode int) {
	testRouter := SetupRouter(true, false)

	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)
	if response.Code != expectedCode {
		t.Errorf("Expected %d", expectedCode)
	}
}

// SoftDeleteFunc is a handler function which sends a PATCH request to the local API
func SoftDeleteFunc(t *testing.T, url string, expectedCode int) {
	testRouter := SetupRouter(true, false)

	request, err := http.NewRequest("PATCH", url, bytes.NewBuffer([]byte("")))
	if err != nil {
		fmt.Println(err)
	}
	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)
	if response.Code != expectedCode {
		t.Errorf("Expected %d", expectedCode)
	}
}

// PutFunc is a handler function which sends a PUT request to the local API
func PutFunc(t *testing.T, json string, url string, expectedCode int) {
	testRouter := SetupRouter(true, false)

	var jsonByte = []byte(json)
	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonByte))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
	}
	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)
	if response.Code != expectedCode {
		t.Errorf("Expected %d", expectedCode)
	}
}
