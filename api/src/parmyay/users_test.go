package parmyay

import (
	"testing"
)

func TestGetUsers(t *testing.T) {
	//GetFunc(t, "/api/v1/users/", 200)
}

func TestGetUsersCorrectId(t *testing.T) {
	//GetFunc(t, "/api/v1/users/1", 200)
}

func TestGetUsersIncorrectId(t *testing.T) {
	//GetFunc(t, "/api/v1/users/0", 404)
}

func TestDeleteUsersCorrectId(t *testing.T) {
	SoftDeleteFunc(t, "/api/v1/users/2", 200)
}

func TestDeleteUsersIncorrectId(t *testing.T) {
	SoftDeleteFunc(t, "/api/v1/users/0", 404)
}
