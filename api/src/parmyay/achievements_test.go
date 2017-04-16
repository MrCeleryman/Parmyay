package parmyay

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGetAchievements(t *testing.T) {
	PurgeDB()

	var assertMany = func(expected []Achievement) {
		got := GetFunc(t, "/api/v1/achievements/", 200)
		if got.Code != 200 {
			t.Errorf("Expected %d, Got %d", 200, got.Code)
		}
		var gotAchievements []Achievement
		json.Unmarshal(got.Body.Bytes(), &gotAchievements)
		if reflect.DeepEqual(gotAchievements, expected) == false {
			t.Errorf("Expected %+v, Got %+v\n", expected, gotAchievements)
		}
	}

	var assertSingle = func(query string, expectedCode int, expected Achievement) {
		got := GetFunc(t, query, expectedCode)
		if got.Code != expectedCode {
			t.Errorf("Expected %d, Got %d", expectedCode, got.Code)
		}
		var gotAchievements Achievement
		json.Unmarshal(got.Body.Bytes(), &gotAchievements)
		if reflect.DeepEqual(gotAchievements, expected) == false {
			t.Errorf("Expected %+v, Got %+v\n", expected, gotAchievements)
		}
	}

	assertMany([]Achievement{})

	var ac1 = Achievement{ID: 1, Achievement: "Reviewed first Parmy!"}
	var ac2 = Achievement{ID: 2, Achievement: "Ate first Parmy!"}
	var ac3 = Achievement{ID: 3, Achievement: "Ate first Parmy!"}
	DB.Create(ac1)
	DB.Create(ac2)
	DB.Create(ac3)
	assertMany([]Achievement{ac1, ac2, ac3})

	cases := []struct {
		query        string
		expectedCode int
		expected     Achievement
	}{
		{"/api/v1/achievements/0", 404, Achievement{}},
		{"/api/v1/achievements/1", 200, ac1},
		{"/api/v1/achievements/2", 200, ac2},
		{"/api/v1/achievements/3", 200, ac3},
	}

	for _, c := range cases {
		assertSingle(c.query, c.expectedCode, c.expected)
	}
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
