package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// PostFunc is a handler function which sends a POST request to the local API
func PostFunc(t *testing.T, json string, url string, expectedCode int) {
	os.Setenv("TEST", "1")
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
func GetFunc(t *testing.T, url string, expectedCode int) {
	os.Setenv("TEST", "1")
	testRouter := SetupRouter(true, false)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)
	if response.Code != expectedCode {
		t.Errorf("Expected %d", expectedCode)
	}
}

// DeleteFunc is a handler function which sends a DELETE request to the local API
func DeleteFunc(t *testing.T, url string, expectedCode int) {
	os.Setenv("TEST", "1")
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
	os.Setenv("TEST", "1")
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
	os.Setenv("TEST", "1")
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
