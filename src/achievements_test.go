package main

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestGetAchievements(t *testing.T) {
	GetFunc(t, "/api/v1/achievements/", 200)
}

func TestGetAchievementWithIdZero(t *testing.T) {
	GetFunc(t, "/api/v1/achievements/0", 404)
}

func TestGetAchievementWithCorrectId(t *testing.T) {
	GetFunc(t, "/api/v1/achievements/1", 200)
}

func TestPostAchievementCorrectModel(t *testing.T) {
	PostFunc(t, `{"achievement":"Reviewed second parmy"}`, "/api/v1/achievements/", 201)
}

func TestPostAchievementIncorrectModel(t *testing.T) {
	PostFunc(t, `{"description":"Reviewed second parmy"}`, "/api/v1/achievements/", 422)
}

func TestPostAchievementBadRequest(t *testing.T) {
	PostFunc(t, `{"achievement":1}`, "/api/v1/achievements/", 400)
}

func TestPostAchievementNoNumbers(t *testing.T) {
	PostFunc(t, `{"achievement":"1"}`, "/api/v1/achievements/", 400)
}

func TestUpdateAchievementCorrectId(t *testing.T) {
	PutFunc(t, `{"achievement":"Updated Achievement"}`, "/api/v1/achievements/2", 200)
}

func TestUpdateAchievementCorrectIdIncorrectModel(t *testing.T) {
	PutFunc(t, `{"description":"Updated Achievement"}`, "/api/v1/achievements/2", 422)
}

func TestUpdateAchievementIncorrectId(t *testing.T) {
	PutFunc(t, `{"achievement":"Updated Achievement"}`, "/api/v1/achievements/0", 404)
}

func TestDeleteAchievementCorrectId(t *testing.T) {
	DeleteFunc(t, "/api/v1/achievements/2", 204)
}

func TestDeleteAchievementIncorrectId(t *testing.T) {
	DeleteFunc(t, "/api/v1/achievements/0", 404)
}
