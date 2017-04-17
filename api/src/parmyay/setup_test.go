package parmyay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("DB_NAME", ":memory:")
	InitDb()
	DB.LogMode(false)

	os.Exit(m.Run())
}

func DoAsserts(t *testing.T, expectedCode int, expected interface{}, serverResponse *httptest.ResponseRecorder) {
	if serverResponse.Code != expectedCode {
		t.Errorf("Expected %d, Got %d", expectedCode, serverResponse.Code)
	}
	// Pull out expected type
	var expectedType = reflect.TypeOf(expected)

	// Initialise a new empty struct of the expected type
	var responsePayload = reflect.Indirect((reflect.New(expectedType))).Interface()
	var expectedValue = reflect.Indirect((reflect.New(expectedType))).Interface()

	// Unmarshalling puts arrays into maps, so also marhsal the expected
	// to turn its arrays into maps, and unify the key names (eg ID -> id)
	encodedExpected, _ := json.Marshal(expected)
	json.Unmarshal(encodedExpected, &expectedValue)
	json.Unmarshal(serverResponse.Body.Bytes(), &responsePayload)

	if reflect.DeepEqual(expectedValue, responsePayload) == false {
		t.Errorf("Expected %+v, Got %+v\n", expectedValue, responsePayload)
	}
}

// Test helper methods
// PostFunc is a handler function which sends a POST request to the local API
func PostFunc(t *testing.T, url string, sendPayload []byte, expectedCode int, expected interface{}) {
	testRouter := SetupRouter(true, false)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(sendPayload))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
	}
	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)

	DoAsserts(t, expectedCode, expected, response)
}

// GetFunc is a handler function which sends a GET request to the local API
func GetFunc(t *testing.T, url string, expectedCode int, expected interface{}) {
	testRouter := SetupRouter(true, false)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)
	DoAsserts(t, expectedCode, expected, response)
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

func PurgeDB() {
	DB.Close()
	InitDb()
	DB.Create(Venue{ID: 1, Address: "30 Willy Wonka Way", VenueName: "Dans House"})
	DB.Create(Venue{ID: 2, Address: "42 Wallaby Way, Sydney", VenueName: "Dans Old House"})

	DB.Create(Review{ID: 1, Notes: "It was prety good", UserID: 1, VenueID: 1, Rating: 10})
	DB.Create(Review{ID: 2, Notes: "Too much water", UserID: 1, VenueID: 1, Rating: 7.8})

	DB.Create(User{ID: 1, FirstName: "Daniel", LastName: "Mitchell", UserName: "DMitch", Email: "d@gmail.com", Password: []byte("Daniel")})
	DB.Create(User{ID: 2, FirstName: "Jerry", LastName: "Seinfeld", UserName: "DarkLordDD", Email: "bee@movie.com", Password: []byte("dfdfgdfgdfg")})

}
