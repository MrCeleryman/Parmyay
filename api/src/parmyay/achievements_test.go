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

	// Test cases
	cases := []Case{
		{"/api/v1/achievements/1", []byte(``), 200, ac1},
		{"/api/v1/achievements/2", []byte(``), 200, ac2},
		{"/api/v1/achievements/3", []byte(``), 200, ac3},

		{"/api/v1/achievements/0", []byte(``), 404, ErrorResult{"Achievement #0 not found"}},
		{"/api/v1/achievements/20", []byte(``), 404, ErrorResult{"Achievement #20 not found"}},
	}
	for _, c := range cases {
		GetFunc(t, c.query, c.expectedCode, c.expected)
	}
}

func TestPostAchievement(t *testing.T) {
	PurgeDB()

	// Test cases
	cases := []Case{
		{"/api/v1/achievements/", []byte(`{"id": -1, "achievement":"Reviewed first Parmy!"}`), 400, ErrorResult{"ID must not be set"}},
		{"/api/v1/achievements/", []byte(`{"id": 1, "achievement":"Reviewed first Parmy!"}`), 400, ErrorResult{"ID must not be set"}},
		{"/api/v1/achievements/", []byte(`{"achievement":"5"}`), 400, ErrorResult{"Achievement cannot be a number"}},
		{"/api/v1/achievements/", []byte(`{"description":"Mad dog"}`), 422, ErrorResult{"Fields are empty"}},
		{"/api/v1/achievements/", []byte(`{}`), 422, ErrorResult{"Fields are empty"}},

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
	for _, c := range cases {
		PostFunc(t, c.query, c.json, c.expectedCode, c.expected)
	}
}

func TestPutAchievement(t *testing.T) {
	PurgeDB()

	// Create test data
	var ac1 = Achievement{ID: 1, Achievement: "Reviewed first Parmy!"}
	var ac2 = Achievement{ID: 2, Achievement: "Ate first Parmy!"}
	DB.Create(ac1)
	DB.Create(ac2)

	// Test cases
	cases := []Case{
		{"/api/v1/achievements/0", []byte(`{"achievement":"Updated Achievement"}`), 404, ErrorResult{"Achievement not found"}},
		{"/api/v1/achievements/0", []byte(``), 404, ErrorResult{"Achievement not found"}},
		{"/api/v1/achievements/10", []byte(`{"achievement":"Updated Achievement"}`), 404, ErrorResult{"Achievement not found"}},
		{"/api/v1/achievements/1", []byte(`{"message":"Updated Achievement"}`), 422, ErrorResult{"The Achievement field is empty"}},
		{"/api/v1/achievements/1", []byte(``), 400, ErrorResult{"The Achievement field is empty"}},

		{"/api/v1/achievements/1", []byte(`{"achievement":"Nice Job"}`), 200, SuccessResult{
			"success": Achievement{1, "Nice Job"},
		}},
		{"/api/v1/achievements/2", []byte(`{"achievement":"second ach"}`), 200, SuccessResult{
			"success": Achievement{2, "second ach"},
		}},
	}

	for _, c := range cases {
		PutFunc(t, c.query, c.json, c.expectedCode, c.expected)
	}
}

func TestDeleteAchievement(t *testing.T) {
	PurgeDB()

	// Create test data
	var ac1 = Achievement{ID: 1, Achievement: "Reviewed first Parmy!"}
	var ac2 = Achievement{ID: 2, Achievement: "Ate first Parmy!"}
	var ac3 = Achievement{ID: 3, Achievement: "Ate first Parmy!"}
	DB.Create(ac1)
	DB.Create(ac2)
	DB.Create(ac3)

	// Test cases
	cases := []Case{
		{"/api/v1/achievements/0", []byte(``), 404, ErrorResult{"Achievement not found"}},
		{"/api/v1/achievements/9", []byte(``), 404, ErrorResult{"Achievement not found"}},

		{"/api/v1/achievements/1", []byte(``), 204, ""},
		{"/api/v1/achievements/2", []byte(``), 204, ""},
	}

	for _, c := range cases {
		DeleteFunc(t, c.query, c.expectedCode, c.expected)
	}
}
