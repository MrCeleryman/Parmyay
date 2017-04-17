package parmyay

import (
	"testing"
)

func TestGetAchievements(t *testing.T) {
	PurgeDB()

	// Test get all when the db is empty
	GetFunc(t, "/api/v1/achievements/", 200, []Achievement{})

	var ac1 = Achievement{ID: 1, Achievement: "Reviewed first Parmy!"}
	var ac2 = Achievement{ID: 2, Achievement: "Ate first Parmy!"}
	var ac3 = Achievement{ID: 3, Achievement: "Ate first Parmy!"}
	DB.Create(ac1)
	DB.Create(ac2)
	DB.Create(ac3)

	// Test get all when there are achievements
	GetFunc(t, "/api/v1/achievements/", 200, []Achievement{ac1, ac2, ac3})

	// Test valid query cases
	cases := []struct {
		query        string
		expectedCode int
		expected     Achievement
	}{
		{"/api/v1/achievements/1", 200, ac1},
		{"/api/v1/achievements/2", 200, ac2},
		{"/api/v1/achievements/3", 200, ac3},
	}
	for _, c := range cases {
		GetFunc(t, c.query, c.expectedCode, c.expected)
	}

	// Test not found conditions
	GetFunc(t, "/api/v1/achievements/0", 404, ErrorResult{"Achievement #0 not found"})
	GetFunc(t, "/api/v1/achievements/20", 404, ErrorResult{"Achievement #20 not found"})
}

func TestPostAchievementCorrectModel(t *testing.T) {
	PurgeDB()

	// Test invalid query cases
	failCases := []struct {
		query        string
		json         []byte
		expectedCode int
		expected     ErrorResult
	}{
		{"/api/v1/achievements/", []byte(`{"id": -1, "achievement":"Reviewed first Parmy!"}`), 400, ErrorResult{"ID must not be set"}},
		{"/api/v1/achievements/", []byte(`{"id": 1, "achievement":"Reviewed first Parmy!"}`), 400, ErrorResult{"ID must not be set"}},
		{"/api/v1/achievements/", []byte(`{"achievement":"5"}`), 400, ErrorResult{"Achievement cannot be a number"}},
		{"/api/v1/achievements/", []byte(`{"description":"Mad dog"}`), 422, ErrorResult{"Fields are empty"}},
		{"/api/v1/achievements/", []byte(`{}`), 422, ErrorResult{"Fields are empty"}},
	}
	for _, c := range failCases {
		PostFunc(t, c.query, c.json, c.expectedCode, c.expected)
	}

	// Test valid query cases
	validCases := []struct {
		query        string
		json         []byte
		expectedCode int
		expected     SuccessResult
	}{
		{"/api/v1/achievements/", []byte(`{"id": 0, "achievement":"Reviewed first Parmy!"}`), 201, SuccessResult{
			"success": Achievement{1, "Reviewed first Parmy!"},
		}},
		{"/api/v1/achievements/", []byte(`{"achievement":"Reviewed second Parmy!"}`), 201, SuccessResult{
			"success": Achievement{2, "Reviewed second Parmy!"},
		}},
		{"/api/v1/achievements/", []byte(`{"achievement":"Reviewed first Parmy!"}`), 201, SuccessResult{
			"success": Achievement{3, "Reviewed first Parmy!"},
		}},
	}
	for _, c := range validCases {
		PostFunc(t, c.query, c.json, c.expectedCode, c.expected)
	}
}

/*
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
*/
