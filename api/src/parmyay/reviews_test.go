package parmyay

import (
	"testing"
)

func TestGetReviews(t *testing.T) {
	PurgeDB()

	GetFunc(t, "/api/v1/reviews/", 200, []Review{})

	// Create test data
	var r1 = Review{ID: 1, UserID: 1, VenueID: 1, Rating: 5, Notes: "Good Parmy", Created: getNow(), Updated: getNow()}
	var r2 = Review{ID: 2, UserID: 1, VenueID: 1, Rating: 10, Notes: "Great Parmy", Created: getNow(), Updated: getNow()}
	DB.Create(r1)
	DB.Create(r2)

	// Test cases
	cases := []Case{
		{"/api/v1/reviews/", []byte(``), 200, []Review{r1, r2}},
		{"/api/v1/reviews/0", []byte(``), 404, ErrorResult{"Review #0 not found"}},
		{"/api/v1/reviews/9", []byte(``), 404, ErrorResult{"Review #9 not found"}},
		{"/api/v1/reviews/1", []byte(``), 200, r1},
	}

	for _, c := range cases {
		GetFunc(t, c.query, c.expectedCode, c.expected)
	}
}

func TestPostReview(t *testing.T) {
	PurgeDB()

	var u1 = User{ID: 1, FirstName: "Daniel", LastName: "Mitchell", UserName: "DMitch", Email: "d@gmail.com", Password: "Daniel", Created: getNow(), Updated: getNow()}
	var v1 = Venue{ID: 1, Address: "30 Willy Wonka Way", VenueName: "Dans House", Longitude: 22.5, Latitude: 50.44, Created: getNow(), Updated: getNow()}
	DB.Create(u1)
	DB.Create(v1)

	// Test cases
	cases := []Case{
		{"/api/v1/reviews/", []byte(`{"VenueID": 1, "Rating": 5, "Notes": "Good Parmy"}`), 422, ErrorResult{"Fields are empty or UserID/VenueID are 0"}},
		{"/api/v1/reviews/", []byte(`{"UserID": 12, "VenueID": 1, "Rating": 5, "Notes": "Good Parmy"}`), 404, ErrorResult{"User #12 not found"}},
		{"/api/v1/reviews/", []byte(`{"UserID": 1, "VenueID": 12, "Rating": 5, "Notes": "Good Parmy"}`), 404, ErrorResult{"Venue #12 not found"}},
		{"/api/v1/reviews/", []byte(`{"UserID": 1, "VenueID": 1, "Rating": 5, "Notes": "Good Parmy"}`), 201, SuccessResult{
			"success": Review{ID: 1, UserID: 1, VenueID: 1, Rating: 5, Notes: "Good Parmy", Created: getNow(), Updated: getNow()},
		}},
		{"/api/v1/reviews/", []byte(`{"UserID": 1, "VenueID": 1, "Rating": 10, "Notes": "Best Parmy"}`), 201, SuccessResult{
			"success": Review{ID: 2, UserID: 1, VenueID: 1, Rating: 10, Notes: "Best Parmy", Created: getNow(), Updated: getNow()},
		}},
	}

	for _, c := range cases {
		PostFunc(t, c.query, c.json, c.expectedCode, c.expected)
	}
}

func TestUpdateReview(t *testing.T) {
	PurgeDB()
	// Create test data
	var r1 = Review{ID: 1, UserID: 1, VenueID: 1, Rating: 5, Notes: "Good Parmy", Created: getNow(), Updated: getNow()}
	var r2 = Review{ID: 2, UserID: 1, VenueID: 1, Rating: 10, Notes: "Great Parmy", Created: getNow(), Updated: getNow()}
	DB.Create(r1)
	DB.Create(r2)

	cases := []Case{
		{"/api/v1/reviews/1", []byte(`{"Rating": 5}`), 422, ErrorResult{"Notes are empty"}},
		{"/api/v1/reviews/0", []byte(`{"Rating": 5, "Notes": "Good Parmy"}`), 404, ErrorResult{"Review #0 not found"}},
		{"/api/v1/reviews/22", []byte(`{"Rating": 5, "Notes": "Good Parmy"}`), 404, ErrorResult{"Review #22 not found"}},
		{"/api/v1/reviews/1", []byte(`{"Rating": 4, "Notes": "Alright Parmy"}`), 200, SuccessResult{
			"success": Review{ID: 1, UserID: 1, VenueID: 1, Rating: 4, Notes: "Alright Parmy", Created: getNow(), Updated: getNow()},
		}},
		{"/api/v1/reviews/2", []byte(`{"Notes": "Very Bad Parmy"}`), 200, SuccessResult{
			"success": Review{ID: 2, UserID: 1, VenueID: 1, Rating: 0, Notes: "Very Bad Parmy", Created: getNow(), Updated: getNow()},
		}},
	}

	for _, c := range cases {
		PutFunc(t, c.query, c.json, c.expectedCode, c.expected)
	}
}

func TestDeleteReview(t *testing.T) {
	PurgeDB()
	// Create test data
	var r1 = Review{ID: 1, UserID: 1, VenueID: 1, Rating: 5, Notes: "Good Parmy", Created: getNow(), Updated: getNow()}
	var r2 = Review{ID: 2, UserID: 1, VenueID: 1, Rating: 10, Notes: "Great Parmy", Created: getNow(), Updated: getNow()}
	DB.Create(r1)
	DB.Create(r2)

	cases := []Case{
		{"/api/v1/reviews/0", []byte(``), 404, ErrorResult{"Review #0 not found"}},
		{"/api/v1/reviews/22", []byte(``), 404, ErrorResult{"Review #22 not found"}},
		{"/api/v1/reviews/1", []byte(``), 200, SuccessResult{
			"success": Review{ID: 1, UserID: 1, VenueID: 1, Rating: 5, Notes: "Good Parmy", Created: getNow(), Updated: getNow(), ValidTo: NullTime{Time: getNow(), Valid: true}},
		}},
		{"/api/v1/reviews/2", []byte(``), 200, SuccessResult{
			"success": Review{ID: 2, UserID: 1, VenueID: 1, Rating: 10, Notes: "Great Parmy", Created: getNow(), Updated: getNow(), ValidTo: NullTime{Time: getNow(), Valid: true}},
		}},
	}

	for _, c := range cases {
		SoftDeleteFunc(t, c.query, c.json, c.expectedCode, c.expected)
	}
}
