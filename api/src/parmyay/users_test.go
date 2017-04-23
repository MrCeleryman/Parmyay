package parmyay

import (
	"testing"
)

func TestGetUser(t *testing.T) {
	PurgeDB()

	GetFunc(t, "/api/v1/users/", 200, []User{})

	// Create test data
	var u1 = User{ID: 1, FirstName: "Daniel", LastName: "Mitchell", UserName: "DMitch", Email: "d@gmail.com", Password: "Daniel", Created: getNow(), Updated: getNow()}
	var u2 = User{ID: 2, FirstName: "Jerry", LastName: "Seinfeld", UserName: "DarkLordDD", Email: "bee@movie.com", Password: "dfdfgdfgdfg", Created: getNow(), Updated: getNow()}
	DB.Create(u1)
	DB.Create(u2)

	// Test cases
	cases := []Case{
		{"/api/v1/users/", []byte(``), 200, []User{u1, u2}},
		{"/api/v1/users/0", []byte(``), 404, ErrorResult{"User #0 not found"}},
		{"/api/v1/users/9", []byte(``), 404, ErrorResult{"User #9 not found"}},
		{"/api/v1/users/1", []byte(``), 200, u1},
	}

	for _, c := range cases {
		GetFunc(t, c.query, c.expectedCode, c.expected)
	}
}

func TestCreateUser(t *testing.T) {
	PurgeDB()

	// Test cases
	cases := []Case{
		{"/api/v1/users/", []byte(`{"lastName": "Mitchell", "userName": "DMitch", "email": "d@gmail.com", "password": "Daniel"}`), 422, ErrorResult{"Fields are empty"}},
		{"/api/v1/users/", []byte(`{"firstName": "Daniel", "lastName": "Mitchell", "userName": "DMitch"}`), 422, ErrorResult{"Fields are empty"}},
		{"/api/v1/users/", []byte(`{"firstName": "Daniel", "lastName": "Mitchell", "userName": "DMitch", "email": "d@gmail.com", "password": "Daniel"}`), 201, SuccessResult{
			"success": User{ID: 1, FirstName: "Daniel", LastName: "Mitchell", UserName: "DMitch", Email: "d@gmail.com", Password: "Daniel", Created: getNow(), Updated: getNow()},
		}},
		{"/api/v1/users/", []byte(`{"firstName": "Daniel", "lastName": "Mitchell", "userName": "DMitch", "email": "d@gmail.com", "password": "Daniel"}`), 422, ErrorResult{"User already exists"}},
		{"/api/v1/users/", []byte(`{"firstName": "Nich", "lastName": "Guy", "userName": "nn", "email": "nnn@gmail.com", "password": "nichnich"}`), 201, SuccessResult{
			"success": User{ID: 2, FirstName: "Nich", LastName: "Guy", UserName: "nn", Email: "nnn@gmail.com", Password: "nichnich", Created: getNow(), Updated: getNow()},
		}},
	}

	for _, c := range cases {
		PostFunc(t, c.query, c.json, c.expectedCode, c.expected)
	}
}

func TestUpdateUser(t *testing.T) {
	PurgeDB()
	// Create test data
	var u1 = User{ID: 1, FirstName: "Daniel", LastName: "Mitchell", UserName: "DMitch", Email: "d@gmail.com", Password: "Daniel", Created: getNow(), Updated: getNow()}
	var u2 = User{ID: 2, FirstName: "Jerry", LastName: "Seinfeld", UserName: "DarkLordDD", Email: "bee@movie.com", Password: "dfdfgdfgdfg", Created: getNow(), Updated: getNow()}
	DB.Create(u1)
	DB.Create(u2)

	cases := []Case{
		{"/api/v1/users/1", []byte(`{"lastName": "Mitchell", "userName": "DMitch", "email": "d@gmail.com", "password": "Daniel"}`), 422, ErrorResult{"Fields are empty"}},
		{"/api/v1/users/0", []byte(`{"firstName": "Daniel", "lastName": "Mitchell", "userName": "DMitch", "email": "d@gmail.com", "password": "Daniel"}`), 404, ErrorResult{"User #0 not found"}},
		{"/api/v1/users/9", []byte(`{"firstName": "Daniel", "lastName": "Mitchell", "userName": "DMitch", "email": "d@gmail.com", "password": "Daniel"}`), 404, ErrorResult{"User #9 not found"}},
		{"/api/v1/users/1", []byte(`{"firstName": "Nich", "lastName": "Mitchell", "userName": "DarkLordDD", "email": "d@gmail.com", "password": "Daniel"}`), 422, ErrorResult{"User already exists"}},
		{"/api/v1/users/1", []byte(`{"firstName": "Nich", "lastName": "Guy", "userName": "nn", "email": "nnn@gmail.com", "password": "nichnich"}`), 200, SuccessResult{
			"success": User{ID: 1, FirstName: "Nich", LastName: "Guy", UserName: "nn", Email: "nnn@gmail.com", Password: "nichnich", Created: getNow(), Updated: getNow()},
		}},
		{"/api/v1/users/2", []byte(`{"firstName": "Nich", "lastName": "Guy", "userName": "nelly", "email": "nnn@gmail.com", "password": "nichnich"}`), 200, SuccessResult{
			"success": User{ID: 2, FirstName: "Nich", LastName: "Guy", UserName: "nelly", Email: "nnn@gmail.com", Password: "nichnich", Created: getNow(), Updated: getNow()},
		}},
	}

	for _, c := range cases {
		PutFunc(t, c.query, c.json, c.expectedCode, c.expected)
	}
}

func TestDeleteUser(t *testing.T) {
	PurgeDB()
	// Create test data
	var u1 = User{ID: 1, FirstName: "Daniel", LastName: "Mitchell", UserName: "DMitch", Email: "d@gmail.com", Password: "Daniel", Created: getNow(), Updated: getNow()}
	var u2 = User{ID: 2, FirstName: "Jerry", LastName: "Seinfeld", UserName: "DarkLordDD", Email: "bee@movie.com", Password: "dfdfgdfgdfg", Created: getNow(), Updated: getNow()}
	DB.Create(u1)
	DB.Create(u2)

	cases := []Case{
		{"/api/v1/users/0", []byte(``), 404, ErrorResult{"User #0 not found"}},
		{"/api/v1/users/9", []byte(``), 404, ErrorResult{"User #9 not found"}},

		{"/api/v1/users/1", []byte(``), 200, SuccessResult{
			"success": User{ID: 1, FirstName: "Daniel", LastName: "Mitchell", UserName: "DMitch", Email: "d@gmail.com", Password: "Daniel", Created: getNow(), Updated: getNow(), Deleted: NullTime{Time: getNow(), Valid: true}},
		}},
		{"/api/v1/users/2", []byte(``), 200, SuccessResult{
			"success": User{ID: 2, FirstName: "Jerry", LastName: "Seinfeld", UserName: "DarkLordDD", Email: "bee@movie.com", Password: "dfdfgdfgdfg", Created: getNow(), Updated: getNow(), Deleted: NullTime{Time: getNow(), Valid: true}},
		}},
	}

	for _, c := range cases {
		SoftDeleteFunc(t, c.query, c.json, c.expectedCode, c.expected)
	}
}
