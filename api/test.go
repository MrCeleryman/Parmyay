package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

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
