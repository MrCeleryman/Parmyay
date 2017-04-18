package parmyay

/*
func TestGetReviews(t *testing.T) {
	//GetFunc(t, "/api/v1/reviews/", 200)
	PurgeDB()

	// Create test data
	var ac1 = Review{ID: 1, UserID: 1, VenueID: 1, Rating: 5, Notes: "", Created: time.Time{sec: 1492412725}, Updated: time.Time{sec: 1492412725}, ValidTo: NullTime{}}
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
}

func TestUpdateReview(t *testing.T) {

}

func TestDeleteReview(t *testing.T) {
	SoftDeleteFunc(t, "/api/v1/reviews/2", 200)
}
*/
